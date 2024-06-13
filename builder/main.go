package builder

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	cloud_aws "github.com/ayush18023/Load_balancer_Fyp/internal/aws"
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

const tempFolder string = "tempKeyConf"

func (b *Builder) GetBaseEnvironment(RuntimeEnv string) string {
	switch strings.ToLower(RuntimeEnv) {
	case "python":
		return "FROM python:3.9"
	case "node":
		return "FROM node:10-alpine"
	case "go":
		return "FROM golang:1.22"
	}
	return ""
}

func (b *Builder) GetWorkdir() string {
	return "WORKDIR /app"
}

func (b *Builder) GetEnvVariables(EnvVars map[string]string) string {
	variables := ""
	for key, value := range EnvVars {
		var envval string = value
		if strings.Contains(value, " ") && string(value[0]) != "'" && string(value[0]) != `"` {
			envval = fmt.Sprintf("'%s'", value)
		}
		variables += fmt.Sprintf("ENV %s=%s", key, envval) + "\n"
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

func (b *Builder) BuildDockerLayers(Name, BuildCmd, StartCmd, RuntimeEnv, KeyGroup string, EnvVars map[string]string) string {
	var DockerFileContent string
	// if KeyGroup != "" {
	// 	DockerFileContent += b.GetBaseEnvironment(RuntimeEnv) + "\n"
	// }
	DockerFileContent += b.GetBaseEnvironment(RuntimeEnv) + "\n"
	DockerFileContent += b.GetWorkdir() + "\n"
	DockerFileContent += b.GetEnvVariables(EnvVars) + "\n"
	DockerFileContent += b.Copyfiles(Name) + "\n"
	DockerFileContent += b.GetRunCommand(BuildCmd) + "\n"
	DockerFileContent += b.GetStartCommand(StartCmd) + "\n"
	return DockerFileContent
}

func (b *Builder) BuildDockerByLang(Name, BuildCmd, StartCmd, RuntimeEnv string, EnvVars map[string]string) string {
	switch strings.ToLower(RuntimeEnv) {
	case "python":
		dockerImage := fmt.Sprintf(`
FROM python:3.9
%s
%s
%s
%s
%s
`, b.GetWorkdir(), b.GetEnvVariables(EnvVars), b.Copyfiles(Name), b.GetRunCommand(BuildCmd), b.GetStartCommand(StartCmd))
		return dockerImage
	case "node":
		dockerImage := fmt.Sprintf(`
FROM node:10-alpine
%s
%s
%s
%s
%s
`, b.GetWorkdir(), b.GetEnvVariables(EnvVars), b.Copyfiles(Name), b.GetRunCommand(BuildCmd), b.GetStartCommand(StartCmd))
		return dockerImage
	case "go":
		dockerImage := fmt.Sprintf(`
FROM golang:1.22
%s
COPY go.mod go.sum ./
RUN go mod download && go mod verify
%s
%s
%s
%s
`, b.GetWorkdir(), b.GetEnvVariables(EnvVars), b.Copyfiles(Name), b.GetRunCommand(BuildCmd), b.GetStartCommand(StartCmd))
		return dockerImage
	}
	return ""
}

func (b *Builder) BuildDockerBySSH(Name, BuildCmd, StartCmd, RuntimeEnv string, EnvVars map[string]string, pubKey string, supervisorConf string) string {
	// pubKey := fmt.Sprintf("../tempKeys/%s.pub.pem",KeyGroup)
	switch strings.ToLower(RuntimeEnv) {
	case "python":
		dockerImage := fmt.Sprintf(`
FROM python:3.9
%s
%s
%s
%s
%s
`, b.GetWorkdir(), b.GetEnvVariables(EnvVars), b.Copyfiles(Name), b.GetRunCommand(BuildCmd), b.GetStartCommand(StartCmd))
		return dockerImage
	case "node":
		dockerImage := fmt.Sprintf(`
FROM ubuntu:latest

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y vim net-tools nodejs npm openssh-server supervisor sudo && \
    groupadd sshgroup && useradd -ms /bin/bash -g sshgroup user && \
    mkdir -p /home/user/.ssh


ADD ../tempKeyConf/%s /home/user/.ssh/authorized_keys
ADD ../tempKeyConf/%s /etc/supervisord.conf

RUN chown user:sshgroup /home/user/.ssh/authorized_keys && \
    chmod 600 /home/user/.ssh/authorized_keys && \
    service ssh start && \
    mkdir -p /var/log/supervisor


WORKDIR /home/user
COPY . .
%s 
RUN %s

CMD ["/usr/bin/supervisord", "--configuration=/etc/supervisord.conf"]

EXPOSE 22
`, pubKey, supervisorConf, b.GetEnvVariables(EnvVars), b.GetRunCommand(BuildCmd))
		return dockerImage
	case "go":
		dockerImage := fmt.Sprintf(`
FROM golang:1.22
%s
COPY go.mod go.sum ./
RUN go mod download && go mod verify
%s
%s
%s
%s
`, b.GetWorkdir(), b.GetEnvVariables(EnvVars), b.Copyfiles(Name), b.GetRunCommand(BuildCmd), b.GetStartCommand(StartCmd))
		return dockerImage
	}
	return ""
}

func (b *Builder) FetchSSHKeys(KeyGroup string) error { //keygroup is name of key
	svc := cloud_aws.NewS3()
	var err error
	keyName := fmt.Sprintf("%s.pem.pub", KeyGroup)
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(keyName),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// keyPth := fmt.Sprintf("../tempKeyConf/%s", keyName)
	keyPth := filepath.Join(tempFolder, keyName)
	return os.WriteFile(keyPth, bt, 0644)
}

func (b *Builder) CreateSupervisorConfig(command string, KeyGroup string) error {
	conf := fmt.Sprintf(`
[supervisord]
logfile=/var/log/supervisor/supervisord.log  ; (main log file;default $CWD/supervisord.log)
logfile_maxbytes=50MB       ; (max main logfile bytes b4 rotation;default 50MB)
logfile_backups=10          ; (num of main logfile rotation backups;default 10)
loglevel=info               ; (log level;default info; others: debug,warn,trace)
nodaemon=true               ; (start in foreground if true;default false)
pidfile=/var/run/supervisord.pid

[program:mainapp]
command=%s
priority=999
autostart=true                ; start at supervisord start (default: true)
autorestart=true
startretries=3

[program:sshd]
command = /usr/sbin/sshd -D 
priority = 10
autorestart = true
startretries = 3	
`, command)
	// keyPth := fmt.Sprintf("../tempKeyConf/%s.conf", KeyGroup)
	keyPth := filepath.Join(tempFolder, KeyGroup+".conf")
	return os.WriteFile(keyPth, []byte(conf), 0600)
}

func (b *Builder) RemoveSSHEnv(keypth, confpth string) {
	os.Remove(keypth)
	os.Remove(confpth)
}

const localCloneRepo string = "repos"

func (b *Builder) BuildRaw(
	Name, GitLink, Branch, BuildCmd, StartCmd, RuntimeEnv, runningPort, KeyGroup string,
	EnvVars map[string]string,
) (string, *internal.ContainerBasedMetric, error) {
	start := time.Now()
	var (
		// buildCtx io.ReadCloser
		err error
	)

	filpth := filepath.Join(localCloneRepo, Name)
	relDockerFile := filepath.Join(filpth, "Dockerfile")
	defer os.RemoveAll(filpth)
	_, err = git.PlainClone(filpth, false, &git.CloneOptions{
		URL:      GitLink,
		Progress: os.Stdout,
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		fmt.Printf("Failed to clone repository: %v\n", err)
		return "", nil, err
	}
	var dockfile string
	if KeyGroup != "" {
		err := b.FetchSSHKeys(KeyGroup)
		if err != nil {
			return "", nil, err
		}
		err = b.CreateSupervisorConfig(StartCmd, KeyGroup)
		if err != nil {
			return "", nil, err
		}
		keyName := fmt.Sprintf("%s.pem.pub", KeyGroup)
		// keyPth := filepath.Join(tempFolder, keyName)
		// conf := filepath.Join(tempFolder, KeyGroup+".conf")
		dockfile = b.BuildDockerBySSH(Name, BuildCmd, StartCmd, RuntimeEnv, EnvVars, keyName, KeyGroup+".conf")
		// defer b.RemoveSSHEnv(keyPth, conf)
	} else {
		dockfile = b.BuildDockerByLang(Name, BuildCmd, StartCmd, RuntimeEnv, EnvVars)
	}
	dockerfileContent := []byte(dockfile)
	err = os.WriteFile(
		relDockerFile,
		dockerfileContent,
		0644,
	)
	if err != nil {
		fmt.Printf("Failed to write Dockerfile: %v\n", err)
		return "", nil, err
	}
	// CreateImage(Name, filpth)
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	hostport, err := internal.GetFreePort()
	if err != nil {
		return "", nil, err
	}
	doc.BuildNewImage(Name, filpth)
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
	basedMetrics, err := internal.GetBasedMetrics(containerID)
	runTime := time.Now()
	if err != nil {
		return "", nil, err
	}
	tag := doc.PushToRegistry(
		Name,
		auth.GetKey("DOCKER_HUB_REPO_NAME"),
	)
	pushTime := time.Now()
	doc.TrashContainer(containerID)
	doc.ClearImages(tag)
	trashTime := time.Now()
	log.Printf("%s ran:%s push:%s trash:%s\n", Name, runTime.Sub(start), pushTime.Sub(runTime), trashTime.Sub(pushTime))
	return tag, basedMetrics, nil
}

func (w *Builder) JoinMaster(masterurl string) {
	w.Master = url.URL{
		Host: masterurl,
	}
	conn, err := grpc.Dial(
		masterurl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	// fmt.Println("here is executed")
	if err != nil {
		panic("Connection failed")
	}
	master := masterrpc.NewMasterServiceClient(conn)
	myurl := fmt.Sprintf(":%d", w.Port)
	// fmt.Println("Till here ", myurl)
	basedMetrics, err := internal.HealthMetricsBased()
	if err != nil {
		panic(err)
	}
	HealthStats := &masterrpc.BasedMetrics{
		CpuPercent:       basedMetrics.CpuPercent,
		MemUsage:         basedMetrics.MemUsage,
		TotalMem:         basedMetrics.TotalMem,
		MemUsedPercent:   float32(basedMetrics.MemUsedPercent),
		DiskUsage:        basedMetrics.DiskUsage,
		TotalDisk:        basedMetrics.TotalDisk,
		DiskUsagePercent: float32(basedMetrics.DiskUsagePercent),
	}
	_, err = master.Join(
		context.Background(),
		&masterrpc.JoinServer{
			Url:   myurl,
			State: "BUILDER",
			Stats: HealthStats,
		},
	)
	if err != nil {
		panic("Join error")
	}
	fmt.Printf("Joined Master Server(%s)\n", masterurl)
	// do get response and hydrate Serverpool
}
