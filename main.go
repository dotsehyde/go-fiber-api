package main

import (
	"api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	// Use the logger middleware
	app.Use(logger.New())
	// engine := django.New("./views", ".django")
	//POST APIs
	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":8500"))
}
