package cities

import (
	"context"
	"database/sql"
	"fmt"
)

type repository interface {
	createCity(city City) (createCityResponse, error)
	listCitiesByName(nameQuery string, limit int, offset int) ([]City, error)
	getCityByID(id int64) (City, error)
	updateCity(id int64, city City) (City, error)
	deleteCity(id int64) error
}

type service struct {
	repo repository
}

func newService(repo repository) service {
	return service{repo: repo}
}

func (s service) getCityByID(ctx context.Context, id int64) (City, error) {
	return s.repo.getCityByID(id)
}

func (s service) listCities(query listQuery) ([]City, error) {
	return s.repo.listCitiesByName(query.Q, query.Limit, query.Offset)
}

func (s service) createCity(ctx context.Context, req createCityRequest) (createCityResponse, error) {
	if err := req.validate(); err != nil {
		return createCityResponse{}, err
	}

	newCity := NewCity(req.Name, req.Latitude, req.Longitude, req.Province, req.Island, req.IsForeign)
	created, err := s.repo.createCity(newCity)
	if err != nil {
		return createCityResponse{}, err
	}

	fmt.Println(created)

	return created, err
}

func (s service) updateCity(ctx context.Context, id int64, req updateCityRequest) (City, error) {
	if err := req.validate(); err != nil {
		return City{}, err
	}

	// Load current city
	current, err := s.repo.getCityByID(id)
	if err != nil {
		return City{}, err
	}

	if req.Name != nil {
		current.Name = *req.Name
	}
	if req.Latitude != nil {
		current.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		current.Longitude = *req.Longitude
	}
	if req.Province != nil {
		current.Province = *req.Province
	}
	if req.Island != nil {
		current.Island = *req.Island
	}
	if req.IsForeign != nil {
		current.IsForeign = *req.IsForeign
	}

	// Persist update
	updated, err := s.repo.updateCity(id, current)
	if err != nil {
		return City{}, err
	}
	return updated, nil
}

func (s service) deleteCity(id int64) error {
	err := s.repo.deleteCity(id)
	if err == sql.ErrNoRows {
		return fmt.Errorf("city with id %d not found", id)
	}
	return err
}
