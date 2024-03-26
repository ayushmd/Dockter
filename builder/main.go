package builder

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"github.com/go-git/go-git/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Builder struct {
	Port   int
	Master url.URL
}

var Builder_ *Builder = &Builder{}

func (b *Builder) GetBaseEnvironment(RuntimeEnv string) string {
	switch strings.ToLower(RuntimeEnv) {
	case "python":
		return "FROM python:3.9"
	case "node":
		return "FROM node:10-alpine"
	case "go":
		return "FROM golang:1.16"
	}
	return ""
}

func (b *Builder) GetWorkdir() string {
	return "WORKDIR /app"
}

func (b *Builder) GetEnvVariables(EnvVars map[string]string) string {
	variables := ""
	for key, value := range EnvVars {
		variables += fmt.Sprintf("ENV %s=%s", key, value) + "\n"
	}
	return variables
}

func (b *Builder) Copyfiles(Name string) string {
	return "COPY . ."
}

func (b *Builder) GetRunCommand(BuildCmd string) string {
	return fmt.Sprintf("RUN %s", BuildCmd)
}

func (b *Builder) GetStartCommand(StartCmd string) string {
	splitted := strings.Split(StartCmd, " ")
	cmdStr := "["
	for i, key := range splitted {
		splitted[i] = fmt.Sprintf("\"%s\"", key)
	}
	cmdStr += strings.Join(splitted, ",")
	cmdStr += "]"
	return fmt.Sprintf("CMD %s", cmdStr)
}

func (b *Builder) BuildDockerLayers(Name, BuildCmd, StartCmd, RuntimeEnv string, EnvVars map[string]string) string {
	var DockerFileContent string
	DockerFileContent += b.GetBaseEnvironment(RuntimeEnv) + "\n"
	DockerFileContent += b.GetWorkdir() + "\n"
	DockerFileContent += b.GetEnvVariables(EnvVars) + "\n"
	DockerFileContent += b.Copyfiles(Name) + "\n"
	DockerFileContent += b.GetRunCommand(BuildCmd) + "\n"
	DockerFileContent += b.GetStartCommand(StartCmd) + "\n"
	return DockerFileContent
}

const localCloneRepo string = "repos"

func (b *Builder) BuildRaw(
	Name, GitLink, Branch, BuildCmd, StartCmd, RuntimeEnv string, runningPort string,
	EnvVars map[string]string,
) (string, error) {
	var (
		// buildCtx io.ReadCloser
		err error
	)

	filpth := filepath.Join(localCloneRepo, Name)
	relDockerFile := filepath.Join(filpth, "Dockerfile")

	_, err = git.PlainClone(filpth, false, &git.CloneOptions{
		URL:      GitLink,
		Progress: os.Stdout,
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		fmt.Printf("Failed to clone repository: %v\n", err)
		return "", err
	}

	dockerfileContent := []byte(b.BuildDockerLayers(Name, BuildCmd, StartCmd, RuntimeEnv, EnvVars))
	err = os.WriteFile(
		relDockerFile,
		dockerfileContent,
		0644,
	)
	if err != nil {
		fmt.Printf("Failed to write Dockerfile: %v\n", err)
		return "", err
	}
	CreateImage(Name, filpth)
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	hostport, err := internal.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	// hostConfig := &container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		nat.Port(fmt.Sprintf("%s/tcp", runningPort)): []nat.PortBinding{
	// 			{
	// 				HostIP:   "0.0.0.0",
	// 				HostPort: fmt.Sprintf("%d/tcp", hostport),
	// 			},
	// 		},
	// 	},
	// 	NetworkMode: "host",
	// }
	containerID, err := doc.RunContainer(Name, Name, []string{fmt.Sprintf("%d:%s/tcp", hostport, runningPort)})
	if err != nil {
		log.Fatal(err)
	}
	defer doc.TrashContainer(containerID)
	tag := doc.PushToRegistry(
		Name,
		auth.GetKey("DOCKER_HUB_REPO_NAME"),
	)
	return tag, nil
}

func (w *Builder) JoinMaster(masterurl string) {
	w.Master = url.URL{
		Host: masterurl,
	}
	conn, err := grpc.Dial(
		masterurl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	fmt.Println("here is executed")
	if err != nil {
		panic("Connection failed")
	}
	master := masterrpc.NewMasterServiceClient(conn)
	myurl := fmt.Sprintf(":%d", w.Port)
	fmt.Println("Till here ", myurl)
	CpuUsage, MemUsage, DiskUsage, err := internal.HealthMetrics()
	if err != nil {
		panic("Health")
	}
	_, err = master.Join(
		context.Background(),
		&masterrpc.JoinServer{
			Url:       myurl,
			State:     "BUILDER",
			CpuUsage:  float32(CpuUsage),
			MemUsage:  float32(MemUsage),
			DiskUsage: float32(DiskUsage),
		},
	)
	if err != nil {
		panic("Join error")
	}
	// do get response and hydrate Serverpool
}
