package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)


var db *gorm.DB

func Get() *gorm.DB {
	return db
}

func Connect(url string, maxIdleConnections, maxOpenConnections int) error {
	var err error
	db, err = gorm.Open("postgres", url)
	if err != nil {
		return err
	}
	db.DB()
	err = db.DB().Ping()
	if err != nil {
		return err
	}
	db.LogMode(false)
	db.DB().SetMaxIdleConns(maxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)
	db.SingularTable(false)
	return nil
}

func Close() {
	db.Close()
}

func init() {
    url := "host=10.10.10.10 user=postgres dbname=sample_go_api password= port=5432 sslmode=disable"
    err := Connect(url,10,10)
	if err != nil {
		panic(err)
	}
	log.Printf("DB connection : Done")
}

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