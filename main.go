package main

import (
	"fmt"
	"kitsurest/database"
	"kitsurest/models"
	"kitsurest/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Message{})
	database.DBConn.AutoMigrate(&models.User{})
	database.DBConn.AutoMigrate(&models.Channel{})
	database.DBConn.AutoMigrate(&models.Guild{})

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
