package builder

import (
	"os/exec"
)

var allowedCmdNames []string = []string{"git", "node", "npm", "go", "docker"}

func isAllowed(name string) bool {
	for _, cmdNames := range allowedCmdNames {
		if cmdNames == name {
			return true
		}
	}
	return false
}

func isClean(name string) bool {
	return isAllowed(name)
}

func Cmd(name string, args ...string) *exec.Cmd {
	if isClean(name) {
		cmd := exec.Command(name, args...)
		// cmd.Stdout = os.Stdout
		return cmd
	}
	return nil
}
