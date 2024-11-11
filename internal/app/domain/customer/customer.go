package customer

type Customer struct {
	CustomerID  int64  `db:"customer_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
	Discount    int    `db:"discount"`
}
