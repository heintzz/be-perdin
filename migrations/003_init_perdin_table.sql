DROP TABLE IF EXISTS trips;
DROP INDEX IF EXISTS idx_trips_employee;
DROP INDEX IF EXISTS idx_trips_status;

CREATE TABLE IF NOT EXISTS trips (
    id BIGSERIAL PRIMARY KEY,
    employee_id TEXT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    purpose TEXT NOT NULL,
    depart_date DATE NOT NULL,
    return_date DATE NOT NULL,
    origin_city_id BIGINT NOT NULL REFERENCES cities(id) ON DELETE RESTRICT,
    destination_city_id BIGINT NOT NULL REFERENCES cities(id) ON DELETE RESTRICT,
    duration_days INT NOT NULL,
    distance_km DOUBLE PRECISION NOT NULL,
    allowance NUMERIC(14,2),
    status TEXT NOT NULL CHECK (status IN ('PENDING','APPROVED','REJECTED')) DEFAULT 'PENDING',
    approved_by TEXT REFERENCES users(id) ON DELETE SET NULL,
    approved_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_trips_employee ON trips (employee_id);
CREATE INDEX IF NOT EXISTS idx_trips_status ON trips (status);


