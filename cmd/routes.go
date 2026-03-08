package main

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView, "new")

	app.Post("/fact", handlers.CreateFact)

	// JSON API
	app.Get("/api/facts", handlers.ListFactsAPI)
	app.Post("/api/facts", handlers.CreateFactAPI)
	app.Get("/api/facts/:id", handlers.GetFactAPI)
	app.Delete("/api/facts/:id", handlers.DeleteFactAPI)
}
