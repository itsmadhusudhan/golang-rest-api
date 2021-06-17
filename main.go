package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang-rest-api/src/database"
	"golang-rest-api/src/router"
	"log"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.SetUp(app)

	fmt.Println("Application is running...")

	log.Fatal(app.Listen(":3000"))
}
