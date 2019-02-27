package main

import (
	"encoding/json"
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

	db.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	if len(users) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No User Found"})
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	err = db.First(&user, params["id"]).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No User Found"})
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	name := r.FormValue("name")
	email := r.FormValue("email")
	user = User{Name: name, Email: email}
	err = db.Create(&user).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Created User"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	err = db.First(&user, params["id"]).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No User Found"})
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	user.Name = name
	user.Email = email
	err = db.Save(&user).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Updated User"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	err = db.Delete(&user).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Deleted User"})
}

func main() {
	//Intialize Router
	router := mux.NewRouter()

	defer db.Close()
	
	log.Printf("Connecting to database")
	initDB()

	//Routes
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	//Starting Server
	log.Printf("Starting web server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
