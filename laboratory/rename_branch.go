package laboratory

import (
	"fmt"
	"os/exec"
)

func RenameBranch(newBranchName string) (string, error) {
	// Command to rename the current branch
	cmd := exec.Command("git", "branch", "-m", newBranchName)

	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to rename branch: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}
