package projects

import (
	"errors"
	"gitlaboratory/utils"
	"log"
)

func ExecuteInDockerContainer(command string, container *Project) (string, error) {
	// Replace "container_id" with your actual container ID or name
	containerID := container.ContainerID
	if containerID == "" {
		return "", errors.New("no containerID was provided for the given project struct")
	}

	// Capture the output
	output, err := utils.RunCommandWithSudo("docker", "exec", containerID, "sh", "-c", command)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	// Print the output
	return string(output), nil
}
