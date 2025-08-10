package trips

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type repository interface {
	createTrip(t Trip) (Trip, error)
	listTrips(q string, limit int, offset int, employeeIDFilter string) ([]Trip, error)
	getTripByID(id int64) (Trip, error)
	updateTrip(id int64, t Trip) (Trip, error)
	deleteTrip(id int64) error
}

type service struct {
	repo repository
}

func newService(repo repository) service {
	return service{repo: repo}
}

func (s service) getTripByID(_ context.Context, id int64) (Trip, error) {
	return s.repo.getTripByID(id)
}

func (s service) listTripsForRole(query listQuery, role string, userID string) ([]Trip, error) {
	// If role is PEGAWAI, constrain to own trips; if SDM, no filter
	employeeFilter := ""
	if role == "PEGAWAI" {
		employeeFilter = userID
	}
	return s.repo.listTrips(query.Q, query.Limit, query.Offset, employeeFilter)
}

func (s service) createTrip(_ context.Context, req createTripRequest) (Trip, error) {
	if err := req.validate(); err != nil {
		return Trip{}, err
	}
	t := Trip{
		EmployeeID:        req.EmployeeID,
		Purpose:           req.Purpose,
		DepartDate:        req.DepartDate,
		ReturnDate:        req.ReturnDate,
		OriginCityID:      req.OriginCityID,
		DestinationCityID: req.DestinationCityID,
		DurationDays:      req.DurationDays,
		DistanceKm:        req.DistanceKm,
		Allowance:         req.Allowance,
		Status:            "PENDING",
	}
	created, err := s.repo.createTrip(t)
	if err != nil {
		return Trip{}, err
	}
	return created, nil
}

func (s service) updateTrip(_ context.Context, id int64, req updateTripRequest) (Trip, error) {
	if err := req.validate(); err != nil {
		return Trip{}, err
	}

	current, err := s.repo.getTripByID(id)
	if err != nil {
		return Trip{}, err
	}

	if req.EmployeeID != nil {
		current.EmployeeID = *req.EmployeeID
	}
	if req.Purpose != nil {
		current.Purpose = *req.Purpose
	}
	if req.DepartDate != nil {
		current.DepartDate = *req.DepartDate
	}
	if req.ReturnDate != nil {
		current.ReturnDate = *req.ReturnDate
	}
	if req.OriginCityID != nil {
		current.OriginCityID = *req.OriginCityID
	}
	if req.DestinationCityID != nil {
		current.DestinationCityID = *req.DestinationCityID
	}
	if req.DurationDays != nil {
		current.DurationDays = *req.DurationDays
	}
	if req.DistanceKm != nil {
		current.DistanceKm = *req.DistanceKm
	}
	if req.Allowance != nil {
		// allow explicit nulling by sending negative sentinel? keep simple: set value
		v := *req.Allowance
		current.Allowance = &v
	}
	if req.Status != nil {
		current.Status = *req.Status
		if current.Status == "APPROVED" && current.ApprovedAt == nil {
			now := time.Now().Format(time.RFC3339)
			current.ApprovedAt = &now
		}
		if current.Status != "APPROVED" {
			// clear approval metadata for non-approved statuses
			current.ApprovedBy = nil
			current.ApprovedAt = nil
		}
	}
	if req.ApprovedBy != nil {
		v := *req.ApprovedBy
		current.ApprovedBy = &v
	}
	if req.ApprovedAt != nil {
		v := *req.ApprovedAt
		current.ApprovedAt = &v
	}

	updated, err := s.repo.updateTrip(id, current)
	if err != nil {
		if err == sql.ErrNoRows {
			return Trip{}, fmt.Errorf("trip with id %d not found", id)
		}
		return Trip{}, err
	}
	return updated, nil
}

func (s service) deleteTrip(id int64) error {
	err := s.repo.deleteTrip(id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("trip with id %d not found", id)
	}
	return err
}

func (s service) approveTrip(id int64, approverUserID string) (Trip, error) {
	if approverUserID == "" {
		return Trip{}, fmt.Errorf("missing approver user id")
	}
	current, err := s.repo.getTripByID(id)
	if err != nil {
		return Trip{}, err
	}
	if current.Status == "APPROVED" {
		return current, nil
	}
	if current.Status == "REJECTED" {
		return Trip{}, fmt.Errorf("cannot approve a rejected trip")
	}
	now := time.Now().Format(time.RFC3339)
	current.Status = "APPROVED"
	current.ApprovedBy = &approverUserID
	current.ApprovedAt = &now
	updated, err := s.repo.updateTrip(id, current)
	if err != nil {
		return Trip{}, err
	}
	return updated, nil
}

func (s service) rejectTrip(id int64, rejectorUserID string) (Trip, error) {
	if rejectorUserID == "" {
		return Trip{}, fmt.Errorf("missing rejector user id")
	}
	current, err := s.repo.getTripByID(id)
	if err != nil {
		return Trip{}, err
	}
	if current.Status == "REJECTED" {
		return current, nil
	}
	if current.Status == "APPROVED" {
		return Trip{}, fmt.Errorf("cannot reject an approved trip")
	}
	current.Status = "REJECTED"
	current.ApprovedBy = nil
	current.ApprovedAt = nil
	updated, err := s.repo.updateTrip(id, current)
	if err != nil {
		return Trip{}, err
	}
	return updated, nil
}
