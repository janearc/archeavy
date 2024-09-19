package main

import (
	fiber "github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Initialize the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber instance with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files (CSS, JS, images) from the public folder
	app.Static("/", "./public")

	// Render the home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Arc Heavy Industries",
			"Tagline": "Stuff and Things",
		})
	})

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
