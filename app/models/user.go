package models

import (
	"errors"

	"github.com/GetSimpl/sample-go-api/config/sampledb"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func FindAllUsers() ([]User, error) {
	var users []User
	var err error
	db := sampledb.Get()
	db.Find(&users)
	if len(users) == 0 {
		err = errors.New("Message: No User Found")
	}
	return users, err
}

func FindUserbyID(id string) (User, error) {
	var user User
	db := sampledb.Get()
	err := db.First(&user, id).Error
	return user, err
}

func CreateUser(name string, email string) (User, error) {
	db := sampledb.Get()
	user := User{Name: name, Email: email}
	err := db.Create(&user).Error
	return user, err
}

func UpdateUser(id string, name string, email string) (User, error) {
	var user User
	db := sampledb.Get()
	err := db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	user.Name = name
	user.Email = email
	err = db.Save(&user).Error
	return user, err
}

func DeleteUserbyID(id string) (User, error) {
	db := sampledb.Get()
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	err = db.Delete(&user).Error
	return user, err
}
