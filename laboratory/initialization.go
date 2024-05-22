package laboratory

import (
	"errors"
	"fmt"
	"os/exec"
)

func Initialization() (string, error) {
	cmd := exec.Command("git", "init")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("failed")
	}

	return string(output), nil
}
