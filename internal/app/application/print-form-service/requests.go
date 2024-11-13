package print_form_service

import "time"

type CreateReportRequest struct {
	DateFrom time.Time
	DateTo   time.Time
}

type Report struct {
}
