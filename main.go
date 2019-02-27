package main

import (
	"log"
	"net/http"

	"github.com/GetSimpl/sample-go-api/app/routes"
	"github.com/GetSimpl/sample-go-api/config/sampledb"
	"github.com/gorilla/mux"
)

var err error

func main() {
	//Intialize Router
	router := mux.NewRouter()

	defer sampledb.Close()

	log.Printf("Connecting to database")
	sampledb.Connect()
	routes.Init(router)

	//Starting Server
	log.Printf("Starting web server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
