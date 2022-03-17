package config

import (
	"assignment_2/structs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	port     = "5432"
	dbName   = "assignment_2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	db.Debug().AutoMigrate(structs.Orders{}, structs.Items{})
}

func GetDB() *gorm.DB {
	return db
}
