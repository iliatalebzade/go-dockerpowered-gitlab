package projects

import (
	"errors"
	"fmt"
	"gitlaboratory/utils"
	"regexp"
	"strings"
)

func isValidID(input string) bool {
	// Define the regular expression pattern
	pattern := "^[a-zA-Z0-9]+$"

	// Compile the regular expression
	regExp := regexp.MustCompile(pattern)

	// Check if the input matches the pattern
	return regExp.MatchString(input)
}

func SetupContainer(project *Project) error {
	// Build the Docker image
	output, err := utils.RunCommandWithSudo("docker", "run", "-d", project.ImageID)
	if err != nil {
		fmt.Println("Error:", err)
		return errors.New(output)
	}

	outputLines := strings.Split(output, "\n")
	possibleOutput := outputLines[0]
	isValid := isValidID(possibleOutput)
	if !isValid {
		return errors.New("had a problem retreiving the container ID from output")
	}

	project.ContainerID = possibleOutput

	return nil
}
