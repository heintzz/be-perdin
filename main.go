package main

import (
	"fmt"
	"log"

	"heintzz/be-perdin/internal/config"
	"heintzz/be-perdin/internal/db"
)

func main() {
	cfg := config.Load()
	database, err := db.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	fmt.Print("connnected to database")
}
