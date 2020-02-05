package main

import (
    "log"
    "fmt"
    "time"
    "strings"
    "strconv"
    "os/exec"
    "net/http"
    // "math/rand"
    "encoding/json"
    "github.com/gorilla/mux"
)

/* Log */

type debugLevel int

const (
    Debug  debugLevel = iota
    Info
    Warning
    Critical
)

var DEBUG_LEVEL debugLevel = Info

func logEvent(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s - %s\n", time.Now().Format(time.RFC850), msg)
    }
}

func logDetail(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s\n%s\n", time.Now().Format(time.RFC850), msg)
    }
}

/* External Invocation Functions */
func runOnHost(cmd string) string {

    parts := strings.Fields(cmd)
    app := parts[0]
    args := parts[1:len(parts)]

    // Add format settings after Fields command
    if args[0] == "image" && args[1] == "ls" {
        args = append(args, "--format")
        args = append(args, `{{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.CreatedAt}}\t{{.Size}}`)
    } else if (args[0] == "container" && args[1] == "ls") {
        // docker container ls --format '{{json .}}'
        args = append(args, "--format")
        args = append(args, `{{.Names}}\t{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}`)
    }

    out, err := exec.Command(app, args...).Output()
    // logDetail(string(out), Critical)

    if (err != nil) {
        logEvent("Error", Critical)
        log.Fatal(err)
        out = []byte("Error")
    }

    return string(out)

}

func parseDockerImages(data string) {

    lines := strings.Split(data, "\n")

    // Remove previous data
    images = nil

    // Line Format (split by \t)
    // Repository Tag Id Created Size
    for i:=0; i < len(lines) && len(lines[i]) > 0; i++ {
        
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
    for i:=0; i < len(lines) && len(lines[i]) > 0; i++ {
        
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

/* General functions */
func strListToIntList(list []string) []int {

    var result []int

    for _, element := range list {
        i, _ := strconv.Atoi(element)
        result = append(result, i)
    }

    return result

}

/* Enable Cors */
// Should disable on production
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

/* Application side */

type lddOperationResponse struct {
    Status string `json:"status"`
    Message string `json:"msg"`
}

type DockerImage struct {
    Repository string `json:"repository"`
    TAG string `json:"tag"`
    ID string `json:"id"`
    Created string `json:"created"`
    Size string `json:"size"`
}

type DockerContainer struct {
    Name string `json:"name"`
    ID string `json:"id"`
    Image string `json:"image"`
    CMD string `json:"cmd"`
    Created string `json:"created"`
    Status string `json:"status"`
    Ports [] int `json:"ports"`
}

var images []DockerImage
var containers []DockerContainer

/* Get complete list of items */

func reqGetDockerImages(w http.ResponseWriter, r *http.Request) {
    // Update images availables
    output := runOnHost("docker image ls")
    parseDockerImages(output)

    // Send response
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    json.NewEncoder(w).Encode(images)
}

func reqGetDockerContainers(w http.ResponseWriter, r *http.Request) {
    // Update images availables
    output := runOnHost("docker container ls")
    parseDockerContainers(output)

    // Send response
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    json.NewEncoder(w).Encode(containers)
}

/* Get specific item */

func reqGetImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    for _, item := range images {
      if item.ID == params["id"] {
        json.NewEncoder(w).Encode(item)
        break
      }
      return
    }
    // json.NewEncoder(w).Encode(&DockerImage{})
}

func reqGetDockerContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    for _, item := range containers {
      if item.ID == params["id"] {
        json.NewEncoder(w).Encode(item)
        break
      }
      return
    }
}

/* Create New Item */

func reqCreateDockerImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    var image DockerImage
    err := json.NewDecoder(r.Body).Decode(&image)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Image not added due to an error!"}
    
    if (err == nil) {
        images = append(images, image)
        answer = lddOperationResponse{ Status: "OK", Message: "Image added" }
        // json.NewEncoder(w).Encode(&image)
    } else {
        fmt.Printf("Error! %s\n", err) // has to be changed for callback to log!
    }

    json.NewEncoder(w).Encode(&answer)
    
}

func reqCreateDockerContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    var container DockerContainer
    err := json.NewDecoder(r.Body).Decode(&container)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Container not created due to an error!"}

    if (err == nil) {
        containers = append(containers, container)
        answer = lddOperationResponse{ Status: "OK", Message: "Container created" }
        // json.NewEncoder(w).Encode(&container)
    } else {
        fmt.Printf("Error! %s\n", err) // has to be changed for callback to log!
    }

    json.NewEncoder(w).Encode(&answer)

}

/* Update Item - Execute action over it */
func reqUpdateImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Image not updated!" }

    for index, item := range images {
      if item.ID == params["id"] {

        // images = append(images[:index], images[index+1:]...)
        var image DockerImage
        err := json.NewDecoder(r.Body).Decode(&image)

        if (err == nil){
            images[index] = image
            answer = lddOperationResponse{ Status: "OK", Message: "Image updated!" }
        }

        break

      }
    }
    
    json.NewEncoder(w).Encode(answer)

}

func reqUpdateContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Container not updated!" }

    for index, item := range containers {
      if item.ID == params["id"] {

        // containers = append(containers[:index], containers[index+1:]...)
        var container DockerContainer
        err := json.NewDecoder(r.Body).Decode(&container)
        
        if (err == nil){
            containers[index] = container
            answer = lddOperationResponse{ Status: "OK", Message: "Container updated!" }
        }

        break

      }
    }

    json.NewEncoder(w).Encode(answer)

}

func reqDeleteImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Image was not removed!"}

    for index, item := range images {
      if item.ID == params["id"] {
        images = append(images[:index], images[index+1:]...)
        answer = lddOperationResponse{ Status: "OK", Message: "Image was removed!"}    
        break
      }
    }

    // json.NewEncoder(w).Encode(images)
    json.NewEncoder(w).Encode(answer)

}

func reqDeleteContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    params := mux.Vars(r)
    answer := lddOperationResponse{ Status: "ERROR", Message: "Container was not removed!"}

    for index, item := range containers {
      if item.ID == params["id"] {
        containers = append(containers[:index], containers[index+1:]...)
        answer = lddOperationResponse{ Status: "OK", Message: "Container was removed!"}    
        break
      }
    }

    // json.NewEncoder(w).Encode(containers)
    json.NewEncoder(w).Encode(answer)

}

func main() {

    images = append(images, DockerImage{Repository: "prueba/imagen1", TAG: "23984723894723984", ID: "1", Created: "25/01/2020", Size: "255KB"})
    images = append(images, DockerImage{Repository: "prueba/imagen2", TAG: "23984723894723984", ID: "2", Created: "25/01/2020", Size: "255KB"})
    images = append(images, DockerImage{Repository: "prueba/imagen3", TAG: "23984723894723984", ID: "3", Created: "25/01/2020", Size: "255KB"})
    images = append(images, DockerImage{Repository: "prueba/imagen4", TAG: "23984723894723984", ID: "4", Created: "25/01/2020", Size: "255KB"})

    router := mux.NewRouter()

    /* Get complete list of items */
    router.HandleFunc("/images", reqGetDockerImages).Methods("GET")
    router.HandleFunc("/containers", reqGetDockerContainers).Methods("GET")

    /* Get specific item */
    router.HandleFunc("/images/{id}", reqGetImage).Methods("GET")
    router.HandleFunc("/containers/{id}", reqGetDockerContainer).Methods("GET")

    /* Create New Item */
    router.HandleFunc("/images", reqCreateDockerImage).Methods("POST")
    router.HandleFunc("/containers", reqCreateDockerContainer).Methods("POST")

    /* Update Item - Execute action over it */
    router.HandleFunc("/images/{id}", reqUpdateImage).Methods("PUT")
    router.HandleFunc("/containers/{id}", reqUpdateContainer).Methods("PUT")

    router.HandleFunc("/images/{id}", reqDeleteImage).Methods("DELETE")
    router.HandleFunc("/containers/{id}", reqDeleteContainer).Methods("DELETE")

    logEvent("Listening", Info)
    http.ListenAndServe(":8000", router)


}