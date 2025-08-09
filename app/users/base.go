package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Run(app *fiber.App, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := handler{svc: svc}

	api := app.Group("/api/v1/users")

	api.Post("/", h.createUser)
	api.Get("/:id", h.getUserProfile)
	api.Patch("/:id/role", h.updateUserRole)
}
