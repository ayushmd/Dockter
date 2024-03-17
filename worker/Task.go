package worker

import (
	"fmt"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func AddTask(id, imageName, runningPort string) (string, error) {
	// repoimageName := internal.GetKey("DOCKER_HUB_REPO_NAME") + imageName
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	err := doc.PullFromRegistery(imageName)
	if err != nil {
		return "", err
	}
	port, err := internal.GetFreePort()
	if err != nil {
		panic("port not found")
	}
	portBindings := nat.PortMap{
		nat.Port(runningPort): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: fmt.Sprintf("%d", port),
			},
		},
	}
	containerID, err := doc.RunContainer(
		imageName,
		&container.HostConfig{
			NetworkMode:  "host",
			PortBindings: portBindings,
		},
	)
	if err != nil {
		return "", nil
	}
	return containerID, nil
}

func TerminateTask(id string) error {
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	return doc.TrashContainer(id)
}
