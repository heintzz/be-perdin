package trips

type Trip struct {
	ID                int64    `json:"id"`
	EmployeeID        string   `json:"employeeId"`
	Purpose           string   `json:"purpose"`
	DepartDate        string   `json:"departDate"`
	ReturnDate        string   `json:"returnDate"`
	OriginCityID      int64    `json:"originCityId"`
	DestinationCityID int64    `json:"destinationCityId"`
	DurationDays      int      `json:"durationDays"`
	DistanceKm        float64  `json:"distanceKm"`
	Allowance         *float64 `json:"allowance,omitempty"`
	Status            string   `json:"status"`
	ApprovedBy        *string  `json:"approvedBy,omitempty"`
	ApprovedAt        *string  `json:"approvedAt,omitempty"`
	CreatedAt         string   `json:"createdAt,omitempty"`
}
