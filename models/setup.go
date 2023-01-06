package models

import (
	"grom.io/driver/postgres"
	"grom.io/grom"
)

var DB *grom.DB

func ConnectDatabase() {
	dsn := "host:localhost user:postgres dbname=go_name post=5432 sslmode=disable"
	database, err := grom.Open(postgres.Open(dsn), &grom.config{})

	if err != nil {
		panic("Failed to connect to Database!")
	}

	database.AutoMigrate(&Post{})

	DB = database
}
