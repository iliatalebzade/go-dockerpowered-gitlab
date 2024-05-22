package laboratory

import (
	"errors"
	"fmt"
	"os/exec"
)

func NewBranch(doSwtich bool, branchName string) (string, error) {
	var cmd *exec.Cmd
	if doSwtich {
		cmd = exec.Command("git", "checkout", "-b", branchName)
	} else {
		cmd = exec.Command("git", "branch", branchName)
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("failed")
	}

	return string(output), nil
}
