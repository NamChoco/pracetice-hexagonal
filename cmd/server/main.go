package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/routes"
	"github.com/NamChoco/pracetice-hexagonal/internal/adapter/sqlite"
)

func main() {
	// setup DB
	db, err := sqlite.InitDB("qa.db")
	if err != nil {
		log.Fatal(err)
	}

	// Fiber app
	app := fiber.New()

	// Register routes
	routes.RegisterQARoutes(app, db)

	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
