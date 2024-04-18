package master

import (
	"context"
	"fmt"

	"github.com/ayush18023/Load_balancer_Fyp/rpc/workerrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type policy string

const (
	CLEAR_INACTIVE policy = "CLEAR_INACTIVE"
	START_INACTIVE policy = "START_INACTIVE"
)

func hasContainer(imagename string, containers []*workerrpc.RunningTask) int {
	for i, container := range containers {
		if container.ImageName == imagename {
			return i
		}
	}
	return -1
}

func RemoveFromDNS(peerurl, imageName string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE HostIp='%s' AND ImageName='%s'", "dns", peerurl, imageName)
	_, err := Master_.dbDns.Exec(query)
	return err
}

func StateManage(state string) (bool, bool) {
	//returns wether to clear dns,wether Trash container on worker
	switch state {
	case "created":
		return false, false
	case "running":
		return false, false
	case "paused":
		return false, false
	case "restarting":
		return false, false
	case "stopped":
		return true, true
	case "exited":
		return true, true
	case "dead":
		return true, true
	case "removing":
		return false, true
	case "deleted":
		return false, true
	}
	return false, false
}

func ClearDNS(peerurl string) error {
	conn, err := grpc.Dial(
		peerurl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	// fmt.Println("here is executed")
	if err != nil {
		return err
	}
	worker := workerrpc.NewWorkerServiceClient(conn)
	containers, err := worker.GetTasks(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE HostIp='%s'", "dns", peerurl)
	rows, err := Master_.dbDns.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var Subdomain string
		var HostIp string
		var HostPort string
		var RunningPort string
		var ImageName string
		var ContainerID string
		err = rows.Scan(&Subdomain, &HostIp, &HostPort, &RunningPort, &ImageName, &ContainerID)
		if err != nil {
			return err
		}
		if ind := hasContainer(ImageName, containers.Containers); ind != -1 {
			toEvcit, _ := StateManage(containers.Containers[ind].State)
			if toEvcit {
				err := RemoveFromDNS(peerurl, ImageName)
				if err != nil {
					return err
				}
			}
		} else {
			err := RemoveFromDNS(peerurl, ImageName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func HandshakePolicy(policyName policy, peerurl string) error {
	if policyName == CLEAR_INACTIVE {
		return ClearDNS(peerurl)
	}
	return nil
}
