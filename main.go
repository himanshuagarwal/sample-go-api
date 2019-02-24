package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    log.Printf("Getting User "+params["id"])
    json.NewEncoder(w).Encode("test")
}

func main() {
    router := mux.NewRouter()
    log.Printf("Starting web server at port 8000")
    router.HandleFunc("/user/{id}", GetUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}