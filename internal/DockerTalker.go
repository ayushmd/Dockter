package internal

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	registery "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

type Dockter struct {
	cli *client.Client
}

func (d *Dockter) Init() {
	var err error
	d.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}

func (d *Dockter) Close() {
	d.cli.Close()
}

func (d *Dockter) CreateImage(imageName string, reader io.Reader) {
	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{imageName + ":latest"},
		Context:    reader,
	}

	buildResponse, err := d.cli.ImageBuild(
		context.Background(),
		bytes.NewReader([]byte("")),
		buildOptions,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Dockter) RunContainer(imageName string, hostConfig *container.HostConfig) (string, error) {
	container, err := d.cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: imageName,
			//may require Cmd:
		}, hostConfig, nil, nil,
		imageName,
	)

	if err != nil {
		return "", err
	}

	if err := d.cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	// hostConfig := &container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		"8080/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "8080"}},
	// 	},
	// }

	return container.ID, nil
}

func (d *Dockter) TrashContainer(containerID string) error {
	if err := d.cli.ContainerKill(context.Background(), containerID, "SIGKILL"); err != nil {
		return err
	}
	removeOptions := container.RemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}
	if err := d.cli.ContainerRemove(context.Background(), containerID, removeOptions); err != nil {
		return err
	}
	return nil
}

func (d *Dockter) ClearImages(imageName string) {
	_, err := d.cli.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = d.cli.ImageRemove(context.Background(), imageName+":latest", types.ImageRemoveOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Dockter) ExecuteCommand(containerID string, cmd []string) []string {
	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Tty:          false,
	}

	execID, err := d.cli.ContainerExecCreate(context.Background(), containerID, execConfig)
	if err != nil {
		panic(err)
	}

	response, err := d.cli.ContainerExecAttach(context.Background(), execID.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}
	defer response.Close()

	scanner := bufio.NewScanner(response.Reader)
	var outputLines []string
	for scanner.Scan() {
		outputLines = append(outputLines, scanner.Text())
	}
	return outputLines
}

func encodeAuthToBase64(authConfig registery.AuthConfig) (string, error) {
	authJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(authJSON), nil
}

func GetKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func (d *Dockter) PushToRegistry(imageName, reponame string) string {
	authConfig := registery.AuthConfig{
		Username:      GetKey("DOCKER_USER"),
		Password:      GetKey("DOCKER_PAT"),
		ServerAddress: "https://index.docker.io/v1/",
	}

	encodedAuth, err := encodeAuthToBase64(authConfig)
	if err != nil {
		panic(err)
	}

	tag := reponame + ":" + imageName

	// Tag the image
	if err := d.cli.ImageTag(context.Background(), imageName, tag); err != nil {
		log.Fatal(err)
	}

	// Push the image to the registry
	pushResponse, err := d.cli.ImagePush(context.Background(), tag, types.ImagePushOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer pushResponse.Close()

	// Log push output
	_, err = io.Copy(os.Stdout, pushResponse)
	if err != nil {
		log.Fatal(err)
	}
	return tag
}

func (d *Dockter) PullFromRegistery(repoimageName string) error {
	out, err := d.cli.ImagePull(context.Background(), repoimageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func FindFirstPort(outputLines []string) string {
	var runningPorts []string
	for _, line := range outputLines {
		if strings.Contains(line, "LISTEN") {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				splits := strings.Split(fields[3], ":")
				runningPorts = append(runningPorts, splits[len(splits)-1])
			}
		}
	}
	if len(runningPorts) > 0 {
		return runningPorts[0]
	}

	return ""
}
