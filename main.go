package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB
var err error

func initDB() {
	url := "host=localhost user=postgres dbname=sample_go_api password= port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func main() {
	//Intialize Router
	router := mux.NewRouter()

	log.Printf("Connecting to database")
	initDB()

	//Routes
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	//Starting Server
	log.Printf("Starting web server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
