package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	connectionString := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable "

	// Open a connection to the database using GORM
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
func GetDB() *gorm.DB {
	return DB
}
