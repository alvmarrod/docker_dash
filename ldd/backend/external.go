package main

import (
	"log"
	"os/exec"
	"strings"
)

/* External Invocation Functions */

// runOnHost takes a command as input and executes it, returning its output
// Arguments are considered based on spaces.
// Returns "Error" string on execution failure
func runOnHost(cmd string) string {

	parts := strings.Fields(cmd)
	app := parts[0]
	args := parts[1:len(parts)]

	out, err := exec.Command(app, args...).Output()
	// logDetail(string(out), Critical)

	if err != nil {
		logEvent("Error", Critical)
		log.Fatal(err)
		out = []byte("Error")
	}

	return string(out)

}

// queryDockerOnHost takes a command as input and executes it, returning its output
// Arguments are keywords to specify what should be executed, NOT real commands
// Returns "Error" string on execution failure
func queryDockerOnHost(cmd []string) string {

	app := "docker"
	var args []string

	// Create command based on input
	// docker container ls --format '{{json .}}'
	if (cmd[0] == "image") && (cmd[1] == "list") {
		args = append(args, "image")
		args = append(args, "ls")
		args = append(args, "--format")
		args = append(args, `{{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.CreatedAt}}\t{{.Size}}`)

	} else if cmd[0] == "container" && cmd[1] == "list" {
		args = append(args, "container")
		args = append(args, "ls")
		args = append(args, "--format")
		args = append(args, `{{.Names}}\t{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}`)

		if cmd[2] == "active" {
			args = append(args, "--filter")
			args = append(args, `status=created`)
			args = append(args, "--filter")
			args = append(args, `status=restarting`)
			args = append(args, "--filter")
			args = append(args, `status=running`)
		} else {
			args = append(args, "--filter")
			args = append(args, `status=paused`)
			args = append(args, "--filter")
			args = append(args, `status=exited`)
			args = append(args, "--filter")
			args = append(args, `status=dead`)
			args = append(args, "--filter")
			args = append(args, `status=removing`)
		}

	} else if cmd[0] == "container" && cmd[1] == "Start" {
		args = append(args, cmd[0])
		args = append(args, "start")
		args = append(args, cmd[2])
	} else if cmd[0] == "container" && cmd[1] == "Stop" {
		args = append(args, cmd[0])
		args = append(args, "stop")
		args = append(args, cmd[2])
	}

	logEvent(strings.Join(args, " || "), Critical)

	out, err := exec.Command(app, args...).Output()
	// logDetail(string(out), Critical)

	if err != nil {
		logEvent("Error", Critical)
		log.Fatal(err)
		out = []byte("Error")
	}

	return string(out)

}
