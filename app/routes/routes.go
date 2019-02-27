package routes

import (
	"github.com/GetSimpl/sample-go-api/app/controllers"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	//Routes
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
}
