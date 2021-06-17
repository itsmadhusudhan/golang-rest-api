package router

import (
	"github.com/gofiber/fiber/v2"
	"golang-rest-api/src/controllers"
)

func SetUp(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/register", controllers.Register)

	api.Post("/login", controllers.Login)

	api.Get("/user", controllers.User)

	api.Post("/logout", controllers.Logout)

}
