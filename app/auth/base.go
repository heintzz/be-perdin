package auth

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Run(app *fiber.App, db *sql.DB, jwtSecret string) {
	repo := newRepository(db)
	svc := newService(repo, jwtSecret)
	h := newHandler(svc)

	api := app.Group("/api/v1/auth")
	api.Post("/register", h.register)
	api.Post("/login", h.login)
}
