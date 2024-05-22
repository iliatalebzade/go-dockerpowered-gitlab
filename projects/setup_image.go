package projects

import (
	"errors"
	"fmt"
	"gitlaboratory/utils"
	"regexp"
	"strings"
)

func SetupImage(project *Project) error {
	// Generate the Dockerfile
	fullDockerfilePath, err := utils.GenerateDockerFile()
	if err != nil {
		panic(err)
	}

	// Build the Docker image
	output, err := utils.RunCommandWithSudo("docker", "build", "-f", fullDockerfilePath, "-t", project.Name, ".")
	if err != nil {
		fmt.Println("Error:", err)
		return errors.New(output)
	}

	// List Docker images
	output, err = utils.RunCommandWithSudo("docker", "images")
	if err != nil {
		fmt.Println("Error:", err)
		return errors.New(output)
	}

	// Compile the regular expression for splitting by whitespace
	re := regexp.MustCompile(`\s+`)

	// Process each line of the output
	for _, line := range strings.Split(output, "\n") {
		// Split the line into parts based on one or more whitespace characters
		parts := re.Split(line, -1)

		repository := parts[0]
		if repository == project.Name {
			tag := parts[1]
			imageID := parts[2]
			created := parts[3]
			size := parts[4]

			project.ImageID = imageID
			project.Tag = tag
			project.Created = created
			project.Size = size
		}
	}

	return nil
}
