package api_service

import "time"

type GetEmployeesOrdersCountRequest struct {
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

type GetCustomerOrdersHistoryRequest struct {
	CustomerID int       `json:"customer_id"`
	DateFrom   time.Time `json:"date_from"`
	DateTo     time.Time `json:"date_to"`
}

type CreateReportRequest struct {
	DateFrom time.Time
	DateTo   time.Time
}
