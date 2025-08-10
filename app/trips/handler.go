package trips

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) createTrip(c *fiber.Ctx) error {
	var req createTripRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	// set employeeId from JWT subject so clients don't send it
	claimsVal := c.Locals("claims")
	claims, _ := claimsVal.(jwt.MapClaims)
	sub, _ := claims["sub"].(string)
	if sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
	}
	req.EmployeeID = sub
	created, err := h.svc.createTrip(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h handler) listTrips(c *fiber.Ctx) error {
	query := parseListQuery(c)
	// get role and subject from JWT
	claimsVal := c.Locals("claims")
	claims, _ := claimsVal.(jwt.MapClaims)
	role, _ := claims["role"].(string)
	sub, _ := claims["sub"].(string)
	trips, err := h.svc.listTripsForRole(query, role, sub)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(trips)
}

func (h handler) getTripByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id parameter"})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}
	trip, err := h.svc.getTripByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(trip)
}

func (h handler) updateTrip(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id parameter"})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}
	var req updateTripRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	updated, err := h.svc.updateTrip(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(updated)
}

func (h handler) deleteTrip(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id parameter"})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}
	if err := h.svc.deleteTrip(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (h handler) approveTrip(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id parameter"})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}
	// approver from JWT claims
	claimsVal := c.Locals("claims")
	claims, _ := claimsVal.(jwt.MapClaims)
	sub, _ := claims["sub"].(string)
	if sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
	}
	updated, err := h.svc.approveTrip(id, sub)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(updated)
}

func (h handler) rejectTrip(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id parameter"})
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}
	claimsVal := c.Locals("claims")
	claims, _ := claimsVal.(jwt.MapClaims)
	sub, _ := claims["sub"].(string)
	if sub == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
	}
	updated, err := h.svc.rejectTrip(id, sub)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(updated)
}
