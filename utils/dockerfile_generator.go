package utils

import (
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// If there is some other error, it might be a permission issue, etc.
	return false
}

func GenerateDockerFile() (string, error) {
	// Define the content of the Dockerfile
	dockerfileContent := `# Use the official Alpine Linux image from the Docker Hub
FROM alpine:latest

# Update the package list and install basic tools, including git
RUN apk update && apk add --no-cache \
    bash \
    curl \
    git

# Set the working directory
WORKDIR /workspace

# Add a simple command to keep the container running
CMD ["tail", "-f", "/dev/null"]
`

	// Create the Dockerfile
	fileName := "Dockerfile"
	tempDir := "temp/"
	fullAddress := tempDir + fileName

	if fileExists(fullAddress) {
		return fullAddress, nil
	}

	// Write the content to the Dockerfile
	err := os.WriteFile(fullAddress, []byte(dockerfileContent), 0644)
	if err != nil {
		return "", err
	}

	return fullAddress, nil
}
