package employee

import "time"

type Employee struct {
	EmployeeID   int64   `db:"employee_id"`
	PositionID   int     `db:"position_id"`
	FirstName    string  `db:"first_name"`
	LastName     string  `db:"last_name"`
	PositionName string  `db:"position_name"`
	Salary       float64 `db:"salary"`
	PhoneNumber  string  `db:"phone_number"`
	Payments     []Payment
}

type Payment struct {
	PaymentID     int64     `db:"payment_id"`
	EmployeeID    int64     `db:"employee_id"`
	TransactionID int64     `db:"transaction_id"`
	DateFrom      time.Time `db:"date_from"`
	Bonus         int       `db:"bonus"`
}
