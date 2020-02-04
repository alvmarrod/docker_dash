package main

import (
    "fmt"
    "time"
    "strconv"
    "net/http"
    "math/rand"
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

func writeLog(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s - %s\n", time.Now().Format(time.RFC850), msg)
    }
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
    Size int64 `json:"size"`    // In Bytes
}

type DockerContainer struct {
    Name string `json:"name"`
    ID string `json:"id"`
    Image DockerImage `json:"image"`
    CMD string `json:"cmd"`
    Created string `json:"created"`
    Status string `json:"status"`
    Ports [] int `json:"ports"`
}

var images []DockerImage
var containers []DockerContainer

/* Get complete list of items */

func getImages(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(images)
}

func getContainers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(containers)
}

/* Get specific item */

func getImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
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

func getContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
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
func createDockerImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var image DockerImage
    err := json.NewDecoder(r.Body).Decode(&image)

    if (err == nil) {
        images = append(images, image)
        // json.NewEncoder(w).Encode(&image)
        lddOperationResponse := lddOperationResponse{ Status: "OK", Message: "Image added" }
        json.NewEncoder(w).Encode(&lddOperationResponse)
    } else {
        fmt.Printf("Error! %s\n", err) // has to be changed for callback to log!
        lddOperationResponse := lddOperationResponse{ Status: "ERROR", Message: "Image not added due to an error!"}
        json.NewEncoder(w).Encode(&lddOperationResponse)
    }
    
}

func createDockerContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var container DockerContainer
    _ = json.NewDecoder(r.Body).Decode(container)
    container.ID = strconv.Itoa(rand.Intn(1000000))
    containers = append(containers, container)
    json.NewEncoder(w).Encode(&container)
}

/* Update Item - Execute action over it */
func updateImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range images {
      if item.ID == params["id"] {
        images = append(images[:index], images[index+1:]...)
        var image DockerImage
        _ = json.NewDecoder(r.Body).Decode(image)
        image.ID = params["id"]
        images = append(images, image)
        json.NewEncoder(w).Encode(&image)
        return
      }
    }
    json.NewEncoder(w).Encode(images)
}

func updateContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range containers {
      if item.ID == params["id"] {
        containers = append(containers[:index], containers[index+1:]...)
        var container DockerContainer
        _ = json.NewDecoder(r.Body).Decode(container)
        container.ID = params["id"]
        containers = append(containers, container)
        json.NewEncoder(w).Encode(&container)
        return
      }
    }
    json.NewEncoder(w).Encode(containers)
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range images {
      if item.ID == params["id"] {
        images = append(images[:index], images[index+1:]...)
        break
      }
    }
    json.NewEncoder(w).Encode(images)
}

func deleteContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range containers {
      if item.ID == params["id"] {
        containers = append(containers[:index], containers[index+1:]...)
        break
      }
    }
    json.NewEncoder(w).Encode(containers)
}

func main() {

    images = append(images, DockerImage{Repository: "prueba/imagen1", TAG: "23984723894723984", ID: "1", Created: "25/01/2020", Size: 255})
    images = append(images, DockerImage{Repository: "prueba/imagen2", TAG: "23984723894723984", ID: "2", Created: "25/01/2020", Size: 255})
    images = append(images, DockerImage{Repository: "prueba/imagen3", TAG: "23984723894723984", ID: "3", Created: "25/01/2020", Size: 255})
    images = append(images, DockerImage{Repository: "prueba/imagen4", TAG: "23984723894723984", ID: "4", Created: "25/01/2020", Size: 255})

    router := mux.NewRouter()

    /* Get complete list of items */
    router.HandleFunc("/images", getImages).Methods("GET")
    router.HandleFunc("/containers", getContainers).Methods("GET")

    /* Get specific item */
    router.HandleFunc("/images/{id}", getImage).Methods("GET")
    router.HandleFunc("/containers/{id}", getContainer).Methods("GET")

    /* Create New Item */
    router.HandleFunc("/images", createDockerImage).Methods("POST")
    router.HandleFunc("/containers", createDockerContainer).Methods("POST")

    /* Update Item - Execute action over it */
    router.HandleFunc("/images/{id}", updateImage).Methods("PUT")
    router.HandleFunc("/containers/{id}", updateContainer).Methods("PUT")

    router.HandleFunc("/images/{id}", deleteImage).Methods("DELETE")
    router.HandleFunc("/containers/{id}", deleteContainer).Methods("DELETE")
    
    writeLog("Listening", Info)
    http.ListenAndServe(":8000", router)


}