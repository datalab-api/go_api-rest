package main

import (
	"api-rest-go/configs"
	"api-rest-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// run database
	configs.ConnectDB()

	// add routes
	routes.UserRoute(app)

	app.Listen(":8082")
}
