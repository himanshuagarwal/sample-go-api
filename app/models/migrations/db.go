package migrations

import (
	"github.com/GetSimpl/sample-go-api/config/sampledb"
	"github.com/GetSimpl/sample-go-api/app/models"
)

func Init() {
	db := sampledb.Get()
	AutoMigrate(&model.user{})
}