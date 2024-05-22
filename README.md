# Docker-Powered Project Workspace

A side project to create a GitLab-like environment powered by Docker, enabling each project or workspace to operate within the isolated space of a fresh container. This setup provides an isolated development environment for each project, ensuring consistent and reproducible builds.

## Features

- **Isolated Environment**: Each project runs inside its own Docker container, ensuring no conflicts between projects.
- **Automated Setup**: Automatically builds Docker images and sets up containers for new projects.
- **Command Execution**: Allows execution of commands inside the Docker container, providing a shell-like interface for managing the project.

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed and running on your machine.
- [Go](https://golang.org/doc/install) installed on your machine.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/docker-powered-workspace.git
    cd docker-powered-workspace
    ```

2. Build the Go application:

    ```sh
    go build -o docker-powered-workspace cmd/main.go
    ```

### Usage

1. Create a new project by editing the `main.go` file and setting up your project details:

    ```go
    me := projects.Project{
        Name:        "uncleilia",
        ImageID:     "",
        Tag:         "",
        Created:     "",
        Size:        "",
        ContainerID: "",
    }
    ```

2. Run the application:

    ```sh
    ./docker-powered-workspace
    ```

3. The application will:
   - Build a Docker image for the project.
   - Create and start a Docker container from the built image.
   - Execute the specified command inside the container and display the output.

### Project Structure

- **cmd/main.go**: The main entry point for the application.
- **projects/execution.go**: Contains functions to set up the project image and container, and to execute commands inside the container.
- **utils/sudocommands.go**: Utility functions for running commands with sudo.

### Example `main.go`

```go
package main

import (
    "fmt"
    "gitlaboratory/projects"
)

func main() {
    me := projects.Project{
        Name:        "uncleilia",
        ImageID:     "",
        Tag:         "",
        Created:     "",
        Size:        "",
        ContainerID: "",
    }

    err := projects.SetupImage(&me)
    if err != nil {
        panic("COULD NOT SET UP THE PROJECT IMAGE " + err.Error())
    }

    err = projects.SetupContainer(&me)
    if err != nil {
        panic("COULD NOT SET UP THE PROJECT CONTAINER" + err.Error())
    }

    fmt.Println("Name: ", me.Name)
    fmt.Println("ImageID: ", me.ImageID)
    fmt.Println("Tag: ", me.Tag)
    fmt.Println("Created: ", me.Created)
    fmt.Println("Size: ", me.Size)
    fmt.Println("ContainerID: ", me.ContainerID)

    // Adding a command to create files for testing
    _, err = projects.ExecuteInDockerContainer("touch /workspace/file1 /workspace/file2", &me)
    if err != nil {
        panic("COULD NOT CREATE TEST FILES INSIDE CONTAINER " + err.Error())
    }

    output, err := projects.ExecuteInDockerContainer("ls -ltrh /workspace", &me)
    if err != nil {
        panic("COULD NOT EXECUTE THE COMMAND INSIDE CONTAINER " + err.Error())
    }

    fmt.Println("=====================================")
    fmt.Println(output)
    fmt.Println("=====================================")
}
