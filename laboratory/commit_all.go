package laboratory

import (
	"errors"
	"fmt"
	"os/exec"
)

func CommitAll(commitMsg string) (string, error) {
	// First command: git add .
	addCmd := exec.Command("git", "add", ".")
	if err := addCmd.Run(); err != nil {
		fmt.Println("Error adding files:", err)
		return "", errors.New("failed to add files")
	}

	// Second command: git commit -m "<commitMsg>"
	commitCmd := exec.Command("git", "commit", "-m", commitMsg)
	output, err := commitCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return "", fmt.Errorf("failed to commit changes: %s", string(output))
	}

	return string(output), nil
}
