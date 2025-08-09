package cities

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type createCityRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Province  string  `json:"province"`
	Island    string  `json:"island"`
	IsForeign bool    `json:"isForeign"`
}

type createCityResponse struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	Island   string `json:"island"`
}

func (req createCityRequest) validate() error {
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Island == "" {
		return fiber.NewError(fiber.StatusBadRequest, "island is required")
	}
	if req.Province == "" {
		return fiber.NewError(fiber.StatusBadRequest, "province is required")
	}
	if req.Longitude == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "longitude is required")
	}
	if req.Latitude == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "latitude is required")
	}
	if req.Latitude < -90 || req.Latitude > 90 {
		return fiber.NewError(fiber.StatusBadRequest, "latitude must be between -90 and 90")
	}
	if req.Longitude < -180 || req.Longitude > 180 {
		return fiber.NewError(fiber.StatusBadRequest, "longitude must be between -180 and 180")
	}
	return nil
}

type listQuery struct {
	Q      string
	Limit  int
	Offset int
}

func parseListQuery(c *fiber.Ctx) listQuery {
	q := c.Query("q", "")
	limitStr := c.Query("limit", "50")
	offsetStr := c.Query("offset", "0")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	return listQuery{Q: q, Limit: limit, Offset: offset}
}

type updateCityRequest struct {
	Name      *string  `json:"name"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Province  *string  `json:"province"`
	Island    *string  `json:"island"`
	IsForeign *bool    `json:"isForeign"`
}

func (req updateCityRequest) validate() error {
	if !req.hasAnyField() {
		return fiber.NewError(fiber.StatusBadRequest, "no fields to update")
	}
	if req.Latitude != nil {
		if *req.Latitude < -90 || *req.Latitude > 90 {
			return fiber.NewError(fiber.StatusBadRequest, "latitude must be between -90 and 90")
		}
	}
	if req.Longitude != nil {
		if *req.Longitude < -180 || *req.Longitude > 180 {
			return fiber.NewError(fiber.StatusBadRequest, "longitude must be between -180 and 180")
		}
	}
	return nil
}

func (req updateCityRequest) hasAnyField() bool {
	return req.Name != nil ||
		req.Latitude != nil ||
		req.Longitude != nil ||
		req.Province != nil ||
		req.Island != nil ||
		req.IsForeign != nil
}
