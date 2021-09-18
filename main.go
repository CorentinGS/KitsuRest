package main

import (
	"kitsurest/database"
	"kitsurest/models"
	"kitsurest/routes"
	"log"
)

func main() {

	// Try to connect to the database
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	// AutoMigrate models
	database.DBConn.AutoMigrate(&models.Message{})
	database.DBConn.AutoMigrate(&models.User{})
	database.DBConn.AutoMigrate(&models.Channel{})
	database.DBConn.AutoMigrate(&models.Guild{})

	// Create the app
	app := routes.New()
	// Listen to port 1812
	log.Fatal(app.Listen(":1812"))
}
