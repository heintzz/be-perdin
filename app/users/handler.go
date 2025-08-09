package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}


func (h handler) updateUserRole(c *fiber.Ctx) error {
	userID := c.Params("id")
	var req UpdateUserRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	req.UserID = userID

	resp, err := h.svc.UpdateUserRole(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h handler) getUserProfile(c *fiber.Ctx) error {
	userID := c.Params("id")
	req := GetUserProfileRequest{UserID: userID}
	resp, err := h.svc.GetUserProfile(req)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
