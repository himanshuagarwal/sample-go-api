package sampledb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func Connect() {
	url := "host=localhost user=postgres dbname=sample_go_api password= port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
