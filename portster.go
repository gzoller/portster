package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strings"
    "strconv"

    "github.com/fsouza/go-dockerclient"
    "github.com/gorilla/mux"
)

var allPorts []docker.APIPort

func main() {
    client, _ := docker.NewClientFromEnv()

    cmd := exec.Command("containerId.sh")
    output, _ := cmd.CombinedOutput()
    cid := strings.Trim(string(output),"\n")
    fmt.Println(cid)

    container, _ := client.InspectContainer(cid)
    allPorts = container.NetworkSettings.PortMappingAPI()
    //fmt.Println( FindPort(9091,ports) )

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/port/{intPort}", GetPort)
    router.HandleFunc("/ping", GetPing)
    log.Fatal(http.ListenAndServe(":1411",router))
}

func FindPort(port int64) int64 {
    for i := 0; i < len(allPorts); i++ {
        if allPorts[i].PrivatePort == port { return allPorts[i].PublicPort }
    }
    return -1
}

func GetPing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w,"pong")
}

func GetPort(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    foundPort, err := strconv.ParseInt(vars["intPort"], 10, 64)
    if( err != nil ) {
        w.WriteHeader(http.StatusBadRequest) 
    } else {
        found := FindPort(foundPort)
        if( found < 0 ) { 
            w.WriteHeader(http.StatusNotFound) 
        } else { 
            fmt.Fprintln(w,found) 
        }
    }

}
