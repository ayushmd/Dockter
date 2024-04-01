package builder

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const reponame = "tyranthex/fyp_deps"

func CloneRepo(link string) {
	cmd := Cmd("git", "clone", link)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error in cloning")
	}
}

func CreateImage(imageName string, path string) {
	cmd := Cmd(
		"docker",
		"build",
		"-t",
		imageName+":latest",
		path,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error in creating image ", err)
	}
}

func BuildContainer(imageName string) {
	cmd := Cmd(
		"docker",
		"run",
		"-d",
		"-i",
		"--name",
		imageName,
		imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal("Couldnt Run Docker container")
	}
}

func RunContainerWithPort(imageName string, hostPorst string, runningPort string) {
	cmd := Cmd(
		"docker",
		"run",
		"-p",
		fmt.Sprintf("%s:%s", hostPorst, runningPort),
		"-d",
		"--name",
		imageName,
		imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal("Couldnt Run Docker container")
	}
}

func TrashContainer(imageName string) {
	var cmd *exec.Cmd
	cmd = Cmd(
		"docker",
		"kill",
		imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
	cmd = Cmd(
		"docker",
		"rm",
		imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
}

func ClearImages(imageName string) {
	var cmd *exec.Cmd
	cmd = Cmd(
		"docker",
		"rmi",
		imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
	cmd = Cmd(
		"docker",
		"rmi",
		reponame+":"+imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
}

func FindPort(imageName string) string {
	cmd := Cmd(
		"docker",
		"exec",
		imageName,
		"netstat",
		"-tuln",
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing command:", err)
	}
	outputLines := strings.Split(out.String(), "\n")
	fmt.Println(outputLines)
	var runningPorts []string
	for _, line := range outputLines {
		fmt.Println(line)
		if strings.Contains(line, "LISTEN") {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				splits := strings.Split(fields[3], ":")
				runningPorts = append(
					runningPorts,
					splits[len(splits)-1],
				)
			}
		}
	}
	return runningPorts[0]
}

func PushToRegistery(imageName string) {
	var cmd *exec.Cmd
	cmd = Cmd(
		"docker",
		"tag",
		imageName,
		reponame+":"+imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
	cmd = Cmd(
		"docker",
		"push",
		reponame+":"+imageName,
	)
	if cmd == nil {
		log.Fatal("Command not allowed")
	}
	cmd.Run()
}
