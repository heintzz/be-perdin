package cities

import (
	"database/sql"
	"time"
)

type cityRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return &cityRepository{db: db}
}

func (r *cityRepository) getCityByID(id int64) (City, error) {
	const query = `
		SELECT id, name, latitude, longitude, province, island, is_foreign
		FROM cities
		WHERE id = $1
	`
	var city City
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&city.ID,
		&city.Name,
		&city.Latitude,
		&city.Longitude,
		&city.Province,
		&city.Island,
		&city.IsForeign,
	); err != nil {
		return City{}, err
	}
	return city, nil
}

func (r *cityRepository) createCity(city City) (newCity createCityResponse, err error) {
	const query = `
		INSERT INTO cities (name, latitude, longitude, province, island, is_foreign, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING name, province, island
	`

	err = r.db.QueryRow(
		query,
		city.Name,
		city.Latitude,
		city.Longitude,
		city.Province,
		city.Island,
		city.IsForeign,
		city.CreatedAt,
		city.UpdatedAt,
	).Scan(
		&newCity.Name,
		&newCity.Province,
		&newCity.Island,
	)
	if err != nil {
		return createCityResponse{}, err
	}

	return newCity, nil
}

func (r *cityRepository) listCitiesByName(nameQuery string, limit int, offset int) ([]City, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	const query = `
		SELECT id, name, latitude, longitude, province, island, is_foreign, created_at, updated_at
		FROM cities
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
		ORDER BY name ASC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, nameQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []City
	for rows.Next() {
		var city City
		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.Latitude,
			&city.Longitude,
			&city.Province,
			&city.Island,
			&city.IsForeign,
			&city.CreatedAt,
			&city.UpdatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, city)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *cityRepository) updateCity(id int64, city City) (City, error) {
	const query = `
		UPDATE cities
		SET name = $1,
				latitude = $2,
				longitude = $3,
				province = $4,
				island = $5,
				is_foreign = $6,
				updated_at = $7
		WHERE id = $8
		RETURNING id, name, latitude, longitude, province, island, is_foreign, updated_at
	`
	currentTime := time.Now()
	err := r.db.QueryRow(
		query,
		city.Name,
		city.Latitude,
		city.Longitude,
		city.Province,
		city.Island,
		city.IsForeign,
		currentTime.Format(time.RFC3339),
		id,
	).Scan(
		&city.ID,
		&city.Name,
		&city.Latitude,
		&city.Longitude,
		&city.Province,
		&city.Island,
		&city.IsForeign,
		&city.UpdatedAt,
	)
	if err != nil {
		return City{}, err
	}
	return city, nil
}

func (r *cityRepository) deleteCity(id int64) error {
	const query = `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
