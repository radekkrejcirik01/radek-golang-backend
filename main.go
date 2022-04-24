package main

import (
	"log"

	"github.com/radekkrejcirik01/radek-golang-backend/pkg/database"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/model/users"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/rest"
)

func main() {
	database.Connect()
	if err := database.DB.AutoMigrate(&users.USERS{}); err != nil {
		log.Fatal(err)
	}

	if err := rest.Create().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
