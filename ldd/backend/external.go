package main

import (
	"log"
	"strings"
	"os/exec"
)

/* External Invocation Functions */

func runOnHost(cmd string) string {

    parts := strings.Fields(cmd)
    app := parts[0]
    args := parts[1:len(parts)]

    out, err := exec.Command(app, args...).Output()
    // logDetail(string(out), Critical)

    if (err != nil) {
        logEvent("Error", Critical)
        log.Fatal(err)
        out = []byte("Error")
    }

    return string(out)

}

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

    } else if (cmd[0] == "container" && cmd[1] == "list") {
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
        
    }

    logEvent(strings.Join(args, " || "), Critical)

    out, err := exec.Command(app, args...).Output()
    // logDetail(string(out), Critical)

    if (err != nil) {
        logEvent("Error", Critical)
        log.Fatal(err)
        out = []byte("Error")
    }

    return string(out)

}