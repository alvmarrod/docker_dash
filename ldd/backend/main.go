package main

import (
    "fmt"
    "time"
    // "strconv"
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

func writeLog(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s - %s\n", time.Now().Format(time.RFC850), msg)
    }
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
    enableCors(&w)
    json.NewEncoder(w).Encode(images)
}

func getContainers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    json.NewEncoder(w).Encode(containers)
}

/* Get specific item */

func getImage(w http.ResponseWriter, r *http.Request) {
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

func getContainer(w http.ResponseWriter, r *http.Request) {
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

func createDockerImage(w http.ResponseWriter, r *http.Request) {
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

func createDockerContainer(w http.ResponseWriter, r *http.Request) {
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
func updateImage(w http.ResponseWriter, r *http.Request) {
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

func updateContainer(w http.ResponseWriter, r *http.Request) {
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

func deleteImage(w http.ResponseWriter, r *http.Request) {
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

func deleteContainer(w http.ResponseWriter, r *http.Request) {
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