package main

import (
	"fmt"
	"log"

	"heintzz/be-perdin/app/auth"
	"heintzz/be-perdin/app/cities"
	"heintzz/be-perdin/app/trips"
	"heintzz/be-perdin/app/users"
	"heintzz/be-perdin/internal/config"
	"heintzz/be-perdin/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.Load()
	jwtSecret := cfg.JWTSecret
	database, err := db.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	fmt.Print("connnected to database")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
	}))

	// mount modules
	auth.Run(app, database, jwtSecret)
	users.Run(app, database, jwtSecret)
	cities.Run(app, database, jwtSecret)
	trips.Run(app, database, jwtSecret)

	port := cfg.Port
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
