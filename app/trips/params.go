package trips

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type createTripRequest struct {
	EmployeeID        string   `json:"employeeId"`
	Purpose           string   `json:"purpose"`
	DepartDate        string   `json:"departDate"`
	ReturnDate        string   `json:"returnDate"`
	OriginCityID      int64    `json:"originCityId"`
	DestinationCityID int64    `json:"destinationCityId"`
	DurationDays      int      `json:"durationDays"`
	DistanceKm        float64  `json:"distanceKm"`
	Allowance         *float64 `json:"allowance"`
}

func (req createTripRequest) validate() error {
	if req.EmployeeID == "" {
		return fmt.Errorf("employee id is required")
	}
	if req.Purpose == "" {
		return fmt.Errorf("purpose is required")
	}
	if req.DepartDate == "" || req.ReturnDate == "" {
		return fmt.Errorf("depart daate and return date are required")
	}
	if _, err := time.Parse("2006-01-02", req.DepartDate); err != nil {
		return fmt.Errorf("invalid depart date format (use YYYY-MM-DD)")
	}
	if _, err := time.Parse("2006-01-02", req.ReturnDate); err != nil {
		return fmt.Errorf("invalid return date format (use YYYY-MM-DD)")
	}
	if req.OriginCityID <= 0 || req.DestinationCityID <= 0 {
		return fmt.Errorf("origin city and destination city are required")
	}
	if req.DurationDays < 0 {
		return fmt.Errorf("durationDays must be >= 0")
	}
	if req.DistanceKm <= 0 {
		return fmt.Errorf("distanceKm must be > 0")
	}

	// ensure logical date order
	d1, _ := time.Parse("2006-01-02", req.DepartDate)
	d2, _ := time.Parse("2006-01-02", req.ReturnDate)
	if d2.Before(d1) {
		return fmt.Errorf("returnDate must be on or after departDate")
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

type updateTripRequest struct {
	EmployeeID        *string  `json:"employeeId"`
	Purpose           *string  `json:"purpose"`
	DepartDate        *string  `json:"departDate"`
	ReturnDate        *string  `json:"returnDate"`
	OriginCityID      *int64   `json:"originCityId"`
	DestinationCityID *int64   `json:"destinationCityId"`
	DurationDays      *int     `json:"durationDays"`
	DistanceKm        *float64 `json:"distanceKm"`
	Allowance         *float64 `json:"allowance"`
	Status            *string  `json:"status"`
	ApprovedBy        *string  `json:"approvedBy"`
	ApprovedAt        *string  `json:"approvedAt"`
}

func (req updateTripRequest) hasAnyField() bool {
	return req.EmployeeID != nil ||
		req.Purpose != nil ||
		req.DepartDate != nil ||
		req.ReturnDate != nil ||
		req.OriginCityID != nil ||
		req.DestinationCityID != nil ||
		req.DurationDays != nil ||
		req.DistanceKm != nil ||
		req.Allowance != nil ||
		req.Status != nil ||
		req.ApprovedBy != nil ||
		req.ApprovedAt != nil
}

func (req updateTripRequest) validate() error {
	if !req.hasAnyField() {
		return fiber.NewError(fiber.StatusBadRequest, "no fields to update")
	}
	if req.DepartDate != nil {
		if _, err := time.Parse("2006-01-02", *req.DepartDate); err != nil {
			return fmt.Errorf("invalid depart date format (use YYYY-MM-DD)")
		}
	}
	if req.ReturnDate != nil {
		if _, err := time.Parse("2006-01-02", *req.ReturnDate); err != nil {
			return fmt.Errorf("invalid return date format (use YYYY-MM-DD)")
		}

	}
	if req.DepartDate != nil && req.ReturnDate != nil {
		d1, _ := time.Parse("2006-01-02", *req.DepartDate)
		d2, _ := time.Parse("2006-01-02", *req.ReturnDate)
		if d2.Before(d1) {
			return fmt.Errorf("return date must be on or after depart date")
		}
	}

	if req.Status != nil {
		s := *req.Status
		if s != "PENDING" && s != "APPROVED" && s != "REJECTED" {
			return fmt.Errorf("invalid status")
		}
	}
	return nil
}
