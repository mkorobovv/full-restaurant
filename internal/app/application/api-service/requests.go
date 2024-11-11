package api_service

import "time"

type GetEmployeesOrdersCountRequest struct {
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}
