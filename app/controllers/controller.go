package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GetSimpl/sample-go-api/app/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := models.FindAllUsers()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No User Found"})
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result, err := models.FindUserbyID(id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No User Found"})
		return
	}
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	result, err := models.CreateUser(name, email)
	log.Println(result)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Created User"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	result, err := models.UpdateUser(id, name, email)
	log.Println(result)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Updated User"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result, err := models.DeleteUserbyID(id)
	log.Println(result)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Some Error Occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Successfully Deleted User"})
}
