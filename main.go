package main

import (
	"fmt"
	"log"

	"heintzz/be-perdin/app/auth"
	"heintzz/be-perdin/app/cities"
	"heintzz/be-perdin/app/users"
	"heintzz/be-perdin/internal/config"
	"heintzz/be-perdin/internal/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	database, err := db.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	fmt.Print("connnected to database")

	app := fiber.New()

	// mount modules
	auth.Run(app, database, cfg.JWTSecret)
	users.Run(app, database, cfg.JWTSecret)
	cities.Run(app, database)

	port := cfg.Port
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
