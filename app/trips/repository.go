package trips

import (
	"database/sql"
	cities "heintzz/be-perdin/app/cities"
	users "heintzz/be-perdin/app/users"
)

type tripRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return &tripRepository{db: db}
}

func (r *tripRepository) getTripByID(id int64) (Trip, error) {
	const query = `
        SELECT
            id,
            employee_id,
            purpose,
            depart_date,
            return_date,
            origin_city_id,
            destination_city_id,
            duration_days,
            distance_km,
            allowance,
            status,
            approved_by,
            approved_at,
            created_at
        FROM trips
        WHERE id = $1
    `
	var t Trip
	var allowance sql.NullFloat64
	var approvedBy sql.NullString
	var approvedAt sql.NullString
	err := r.db.QueryRow(query, id).Scan(
		&t.ID,
		&t.EmployeeID,
		&t.Purpose,
		&t.DepartDate,
		&t.ReturnDate,
		&t.OriginCityID,
		&t.DestinationCityID,
		&t.DurationDays,
		&t.DistanceKm,
		&allowance,
		&t.Status,
		&approvedBy,
		&approvedAt,
		&t.CreatedAt,
	)
	if err != nil {
		return Trip{}, err
	}
	if allowance.Valid {
		t.Allowance = &allowance.Float64
	}
	if approvedBy.Valid {
		v := approvedBy.String
		t.ApprovedBy = &v
	}
	if approvedAt.Valid {
		v := approvedAt.String
		t.ApprovedAt = &v
	}
	return t, nil
}

func (r *tripRepository) createTrip(t Trip) (Trip, error) {
	const query = `
        INSERT INTO trips (
            employee_id, purpose, depart_date, return_date,
            origin_city_id, destination_city_id, duration_days,
            distance_km, allowance
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
        RETURNING id, employee_id, purpose, depart_date, return_date,
                  origin_city_id, destination_city_id, duration_days,
                  distance_km, allowance, status, approved_by, approved_at, created_at
    `
	var created Trip
	var allowance sql.NullFloat64
	var approvedBy sql.NullString
	var approvedAt sql.NullString
	var allowanceParam any
	if t.Allowance != nil {
		allowanceParam = *t.Allowance
	} else {
		allowanceParam = nil
	}

	err := r.db.QueryRow(
		query,
		t.EmployeeID,
		t.Purpose,
		t.DepartDate,
		t.ReturnDate,
		t.OriginCityID,
		t.DestinationCityID,
		t.DurationDays,
		t.DistanceKm,
		allowanceParam,
	).Scan(
		&created.ID,
		&created.EmployeeID,
		&created.Purpose,
		&created.DepartDate,
		&created.ReturnDate,
		&created.OriginCityID,
		&created.DestinationCityID,
		&created.DurationDays,
		&created.DistanceKm,
		&allowance,
		&created.Status,
		&approvedBy,
		&approvedAt,
		&created.CreatedAt,
	)
	if err != nil {
		return Trip{}, err
	}
	if allowance.Valid {
		created.Allowance = &allowance.Float64
	}
	if approvedBy.Valid {
		v := approvedBy.String
		created.ApprovedBy = &v
	}
	if approvedAt.Valid {
		v := approvedAt.String
		created.ApprovedAt = &v
	}
	return created, nil
}

func (r *tripRepository) listTrips(q string, limit int, offset int, employeeIDFilter string) ([]getTripResponse, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	const query = `
        SELECT
            t.id,
            t.purpose,
            t.depart_date,
            t.return_date,
            t.duration_days,
            t.distance_km,
            t.allowance,
            t.status,
            u.username,
            oc.id, oc.name, oc.is_foreign,
            dc.id, dc.name, dc.is_foreign
        FROM trips t
        JOIN cities oc ON oc.id = t.origin_city_id
        JOIN cities dc ON dc.id = t.destination_city_id
        JOIN users u ON u.id = t.employee_id
        WHERE ($1 = '' OR t.purpose ILIKE '%' || $1 || '%')
          AND ($4 = '' OR t.employee_id = $4)
        ORDER BY t.created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := r.db.Query(query, q, limit, offset, employeeIDFilter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []getTripResponse
	for rows.Next() {
		var t getTripResponse
		var allowance sql.NullFloat64
		var requester users.UserOnTripResponse
		var origin cities.CityOnTripResponse
		var destination cities.CityOnTripResponse
		if err := rows.Scan(
			&t.ID,
			&t.Purpose,
			&t.DepartDate,
			&t.ReturnDate,
			&t.DurationDays,
			&t.DistanceKm,
			&allowance,
			&t.Status,
			&requester.Username,
			&origin.ID, &origin.Name, &origin.IsForeign,
			&destination.ID, &destination.Name, &destination.IsForeign,
		); err != nil {
			return nil, err
		}
		if allowance.Valid {
			t.Allowance = &allowance.Float64
		}
		t.Employee = &requester
		t.OriginCity = &origin
		t.DestinationCity = &destination
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *tripRepository) updateTrip(id int64, t Trip) (Trip, error) {
	const query = `
        UPDATE trips
        SET employee_id = $1,
            purpose = $2,
            depart_date = $3,
            return_date = $4,
            origin_city_id = $5,
            destination_city_id = $6,
            duration_days = $7,
            distance_km = $8,
            allowance = $9,
            status = $10,
            approved_by = $11,
            approved_at = $12
        WHERE id = $13
        RETURNING id, employee_id, purpose, depart_date, return_date,
                  origin_city_id, destination_city_id, duration_days,
                  distance_km, allowance, status, approved_by, approved_at, created_at
    `

	var allowanceParam any
	if t.Allowance != nil {
		allowanceParam = *t.Allowance
	} else {
		allowanceParam = nil
	}
	var approvedByParam any
	if t.ApprovedBy != nil {
		approvedByParam = *t.ApprovedBy
	} else {
		approvedByParam = nil
	}
	var approvedAtParam any
	if t.ApprovedAt != nil {
		approvedAtParam = *t.ApprovedAt
	} else {
		approvedAtParam = nil
	}

	var updated Trip
	var allowance sql.NullFloat64
	var approvedBy sql.NullString
	var approvedAt sql.NullString

	err := r.db.QueryRow(
		query,
		t.EmployeeID,
		t.Purpose,
		t.DepartDate,
		t.ReturnDate,
		t.OriginCityID,
		t.DestinationCityID,
		t.DurationDays,
		t.DistanceKm,
		allowanceParam,
		t.Status,
		approvedByParam,
		approvedAtParam,
		id,
	).Scan(
		&updated.ID,
		&updated.EmployeeID,
		&updated.Purpose,
		&updated.DepartDate,
		&updated.ReturnDate,
		&updated.OriginCityID,
		&updated.DestinationCityID,
		&updated.DurationDays,
		&updated.DistanceKm,
		&allowance,
		&updated.Status,
		&approvedBy,
		&approvedAt,
		&updated.CreatedAt,
	)
	if err != nil {
		return Trip{}, err
	}
	if allowance.Valid {
		updated.Allowance = &allowance.Float64
	}
	if approvedBy.Valid {
		v := approvedBy.String
		updated.ApprovedBy = &v
	}
	if approvedAt.Valid {
		v := approvedAt.String
		updated.ApprovedAt = &v
	}
	return updated, nil
}

func (r *tripRepository) deleteTrip(id int64) error {
	_, err := r.db.Exec(`DELETE FROM trips WHERE id = $1`, id)
	return err
}
