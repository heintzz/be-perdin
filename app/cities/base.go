package cities

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Run(app *fiber.App, db *sql.DB) {
	repo := NewRepository(db)
	svc := newService(repo)
	h := newHandler(svc)

	api := app.Group("/api/v1/cities")
	api.Post("/", h.createCity)
	api.Get("/", h.listCities)
	api.Get("/:id", h.getCityByID)
	api.Patch("/:id", h.updateCity)
	api.Delete("/:id", h.deleteCity)
}
