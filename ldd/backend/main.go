package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

/* Enable Cors */
// Should disable on production
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

/* Get complete list of items */

func reqGetDockerImages(w http.ResponseWriter, r *http.Request) {
    // Update images availables
    output := queryDockerOnHost([]string {"image", "list"})
    parseDockerImages(output)

    // Send response
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    json.NewEncoder(w).Encode(images)
}

func reqGetRunningDockerContainers(w http.ResponseWriter, r *http.Request) {
    // Update images availables
    output := queryDockerOnHost([]string{"container", "list", "active"})
    parseDockerContainers(output)

    // Send response
    w.Header().Set("Content-Type", "application/json")
    enableCors(&w)
    json.NewEncoder(w).Encode(containers)
}

func reqGetStoppedDockerContainers(w http.ResponseWriter, r *http.Request) {
    // Update images availables
    output := queryDockerOnHost([]string{"container", "list", "stopped"})
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

    router := mux.NewRouter()

    /* Get complete list of items */
    router.HandleFunc("/images", reqGetDockerImages).Methods("GET")
    router.HandleFunc("/stoppedcontainers", reqGetStoppedDockerContainers).Methods("GET")
    router.HandleFunc("/runningcontainers", reqGetRunningDockerContainers).Methods("GET")

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