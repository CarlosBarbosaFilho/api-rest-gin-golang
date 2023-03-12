package database

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {

	dsn := "host=localhost user=postgres password=postgres dbname=cbgomes-personas-gin port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.Persona{})

}
