package main

import (
	"strings"
)

/* Types definition */

type lddOperationResponse struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
}

type DockerImage struct {
	Repository string `json:"repository"`
	TAG        string `json:"tag"`
	ID         string `json:"id"`
	Created    string `json:"created"`
	Size       string `json:"size"`
}

type DockerContainer struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Image   string `json:"image"`
	CMD     string `json:"cmd"`
	Created string `json:"created"`
	Status  string `json:"status"`
	Ports   []int  `json:"ports"`
}

/* Global variables */
var images []DockerImage
var containers []DockerContainer

/* Functions */

func parseDockerImages(data string) {

	lines := strings.Split(data, "\n")

	// Remove previous data
	images = nil

	// Line Format (split by \t)
	// Repository Tag Id Created Size
	for i := 0; i < len(lines) && len(lines[i]) > 0; i++ {

		piece := strings.Split(lines[i], "\t")

		// logEvent("Line " + string(i) + " = " + lines[i], Critical)
		// for j:=0; j < len(piece); j++ {
		//     logEvent("\t" + piece[j], Critical)
		// }
		// logEvent(lines[i], Critical)

		// Add new item to array
		var image DockerImage
		image.Repository = piece[0]
		image.TAG = piece[1]
		image.ID = piece[2]
		image.Created = piece[3]
		image.Size = piece[4] // strconv.ParseInt(, 10, 64)

		images = append(images, image)

	}

}

func parseDockerContainers(data string) {

	lines := strings.Split(data, "\n")

	// Remove previous data
	containers = nil

	// Line Format (split by \t)
	// Repository Tag Id Created Size
	for i := 0; i < len(lines) && len(lines[i]) > 0; i++ {

		piece := strings.Split(lines[i], "\t")

		// logEvent("Line " + string(i) + " = " + lines[i], Critical)
		// for j:=0; j < len(piece); j++ {
		//     logEvent("\t" + piece[j], Critical)
		// }
		// logEvent(lines[i], Critical)

		// Add new item to array
		var container DockerContainer
		container.Name = piece[0]
		container.ID = piece[1]
		container.CMD = piece[3]
		container.Created = piece[4] // strconv.ParseInt(, 10, 64)
		container.Status = piece[5]

		// If it is NOT found by ID, means it came with a name
		img := getImageByID(piece[2])
		name := ""
		if img == nil {
			name = piece[2]
		} else {
			name = (*img).ID
		}
		container.Image = name

		// Ports need to be parsed as a list
		container.Ports = strListToIntList(strings.Split(piece[6], " "))

		containers = append(containers, container)

	}

}

func getImageByID(pID string) *DockerImage {

	var imagen *DockerImage = nil

	for _, element := range images {

		if element.ID == pID {
			imagen = &element
			break
		}

	}

	return imagen

}

func getContainerByID(pID string) *DockerContainer {

	var container *DockerContainer = nil

	for _, element := range containers {

		if element.ID == pID {
			container = &element
			break
		}

	}

	return container

}

func executeContainerAction(pContainer string, pAction string) string {

	// var container *DockerContainer = getContainerByID(pContainer)
	var cmd = []string{"container", pAction, pContainer}

	// if pAction == "Start" {
	// 	cmd = append(cmd, "")
	// } else if pAction == "Stop" {
	//     cmd = append(cmd, "")
	// }

	output := queryDockerOnHost(cmd)

	return output

}
