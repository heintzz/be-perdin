package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	middleware "heintzz/be-perdin/app"
)

func Run(app *fiber.App, db *sql.DB, jwtSecret string) {
	repo := NewRepository(db)
	svc := NewService(repo, jwtSecret)
	h := handler{svc: svc}

	api := app.Group("/api/v1/users", middleware.AuthenticateJWT(jwtSecret))
	api.Get("/:id", h.getUserProfile)
	api.Patch("/:id/role", middleware.RequireRole("SDM"), h.updateUserRole)
}
