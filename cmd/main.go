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

	fmt.Println("=====================================  DOCKER INFO  =====================================")
	fmt.Println("Name: ", me.Name)
	fmt.Println("ImageID: ", me.ImageID)
	fmt.Println("Tag: ", me.Tag)
	fmt.Println("Created: ", me.Created)
	fmt.Println("Size: ", me.Size)
	fmt.Println("ContainerID: ", me.ContainerID)
	fmt.Println("=========================================================================================")

	output, err := projects.ExecuteInDockerContainer("ls -ltrh", &me)
	if err != nil {
		panic("COULD NOT EXECUTE THE COMMAND INSIDE CONTAINER" + err.Error())
	}

	fmt.Println("=====================================  FINAL RESULT  =====================================")
	fmt.Println(output)
	fmt.Println("==========================================================================================")
}
