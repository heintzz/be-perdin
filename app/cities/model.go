package cities

import "time"

type City struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Province  string  `json:"province"`
	Island    string  `json:"island"`
	IsForeign bool    `json:"isForeign"`
	CreatedAt string  `json:"createdAt,omitempty"`
	UpdatedAt string  `json:"updatedAt,omitempty"`
}

func NewCity(
	name string,
	latitude float64,
	longitude float64,
	province string,
	island string,
	is_foreign bool,
) City {
	currentTime := time.Now()
	return City{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		Province:  province,
		Island:    island,
		IsForeign: is_foreign || false,
		CreatedAt: currentTime.Format(time.RFC3339),
		UpdatedAt: currentTime.Format(time.RFC3339),
	}
}
