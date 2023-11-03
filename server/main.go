package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pattyarao/wechoose/database"
	"github.com/pattyarao/wechoose/router"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// makes sure the user_name is unique
	if err := database.CreateUniqueIndex(); err != nil {
		log.Fatal(err)
	}

	database.Disconnect()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
