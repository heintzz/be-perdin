package trips

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	middleware "heintzz/be-perdin/app"
)

func Run(app *fiber.App, db *sql.DB, jwtSecret string) {
	repo := NewRepository(db)
	svc := newService(repo)
	h := newHandler(svc)

	api := app.Group("/api/v1/trips", middleware.AuthenticateJWT(jwtSecret))
	api.Post("/", h.createTrip)
	api.Get("/", h.listTrips)
	api.Get("/:id", h.getTripByID)
	api.Patch("/:id", h.updateTrip)
	api.Delete("/:id", h.deleteTrip)

	api.Post("/:id/approve", middleware.RequireRole("SDM"), h.approveTrip)
	api.Post("/:id/reject", middleware.RequireRole("SDM"), h.rejectTrip)
}
