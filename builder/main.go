package builder

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/go-git/go-git/v5"
)

type Builder struct {
	Port   int
	Master url.URL
}

func (b *Builder) GetBaseEnvironment(RuntimeEnv string) string {
	switch RuntimeEnv {
	case "Python":
		return "FROM python:3.9"
	case "Node":
		return "FROM node:latest"
	case "Go":
		return "FROM golang:1.16"
	}
	return ""
}

func (b *Builder) GetWorkdir() string {
	return "WORKDIR /app"
}

func (b *Builder) Copyfiles(Name string) string {
	return fmt.Sprintf("COPY ./%s ./", Name)
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

func (b *Builder) BuildDockerLayers(Name, BuildCmd, StartCmd, RuntimeEnv string) string {
	var DockerFileContent string
	DockerFileContent += b.GetBaseEnvironment(RuntimeEnv) + "\n"
	DockerFileContent += b.GetWorkdir() + "\n"
	DockerFileContent += b.Copyfiles(Name) + "\n"
	DockerFileContent += b.GetRunCommand(BuildCmd) + "\n"
	DockerFileContent += b.GetStartCommand(StartCmd) + "\n"
	return DockerFileContent
}

func (b *Builder) BuildRaw(
	Name, GitLink, Branch, BuildCmd, StartCmd, RuntimeEnv string,
	EnvVars map[string]string,
) (string, error) {
	_, err := git.PlainClone(Name, false, &git.CloneOptions{
		URL:      GitLink,
		Progress: os.Stdout,
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		fmt.Printf("Failed to clone repository: %v\n", err)
		return "", err
	}
	dockerfileContent := []byte(b.BuildDockerLayers(Name, BuildCmd, StartCmd, RuntimeEnv))
	err = os.WriteFile(
		fmt.Sprintf("%s/Dockerfile", Name),
		dockerfileContent,
		0644,
	)
	if err != nil {
		fmt.Printf("Failed to write Dockerfile: %v\n", err)
		return "", err
	}

	buildContext, err := os.Open(Name)
	if err != nil {
		fmt.Printf("Failed to open build context: %v\n", err)
		return "", err
	}
	defer buildContext.Close()
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	doc.CreateImage(Name, buildContext)
	containerID, err := doc.RunContainer(Name, nil)
	if err != nil {
		return "", err
	}
	lines := doc.ExecuteCommand(containerID, []string{"netstat", "-tuln"})
	port := internal.FindFirstPort(lines)
	doc.TrashContainer(containerID)
	doc.PushToRegistry(
		Name,
		internal.GetKey("DOCKER_HUB_REPO_NAME"),
	)
	return port, nil
}
