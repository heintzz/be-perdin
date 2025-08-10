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
	Status            string   `json:"status"`
	CreatedAt         string   `json:"createdAt,omitempty"`
	Allowance         *float64 `json:"allowance,omitempty"`
	ApprovedBy        *string  `json:"approvedBy,omitempty"`
	ApprovedAt        *string  `json:"approvedAt,omitempty"`
}
