package main

import (
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
	app.Post("/create", createPost)
	app.Get("/get", getPost)
	app.Get("/getById", getById)
	app.Patch("/update/:id", updatePost)
	app.Delete("/del/:id", deletePost)

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.SendFile("./public/index.html", true)
		// return c.Render("index", fiber.Map{
		// 	"Title": "Hello, World!",
		// })
	})

	log.Fatal(app.Listen(":8000"))
}
