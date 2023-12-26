package routes

import (
	"api/api"

	"github.com/gofiber/fiber/v2"
)

var RegisterRoutes = func(app *fiber.App) {
	app.Post("/create", api.CreatePost)
	app.Get("/get", api.GetPost)
	app.Get("/getById", api.GetById)
	app.Patch("/update/:id", api.UpdatePost)
	app.Delete("/del/:id", api.DeletePost)

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.SendFile("./public/index.html", true)
		// return c.Render("index", fiber.Map{
		// 	"Title": "Hello, World!",
		// })
	})
}
