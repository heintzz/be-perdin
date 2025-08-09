package users

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func (h handler) createUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	user, err := h.svc.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(user)
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