package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/yogesh/shortenurl/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "API is working"})
	})

	app.Get("/:url", routes.ResolveURL)
	app.Post("/app/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
