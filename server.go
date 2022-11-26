package main

import (
	"Toolhub/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

// "github.com/gofiber/fiber/v2/middleware/cors"
func setUpRoutes(app *fiber.App) {
	app.Get("/", routes.IndexHandler)
	app.Get("/tools", routes.ToolHandler)
	// app.Get("/book/:id", routes.GetBook)
	// app.Post("/book", routes.AddBook)
	// app.Put("/book/:id", routes.Update)
	// app.Delete("/book/:id", routes.Delete)
}

func main() {
	app := fiber.New()

	// app.Get("/", routes.IndexHandler)
	// app.Get("/tools", routes.ToolHandler)
	setUpRoutes(app)
	app.Use(cors.New())
	app.Listen(":3000")

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
