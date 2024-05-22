package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunCommandWithSudo(cmdName string, args ...string) (string, error) {
	// Prepend "sudo" to the command arguments
	cmdArgs := append([]string{cmdName}, args...)
	cmd := exec.Command("sudo", cmdArgs...)

	// Create a buffer to capture command output
	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	// Return the captured output as a string
	return outBuffer.String(), nil
}
