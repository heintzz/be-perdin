package db

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Open(dsn string) (*sql.DB, error) {
	_ = godotenv.Load()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return db, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
