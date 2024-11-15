package order

import "time"

type Order struct {
	OrderID       int64     `db:"order_id"`
	CustomerID    int64     `db:"customer_id"`
	TransactionID int64     `db:"transaction_id"`
	CreatedAt     time.Time `db:"created_at"`
	Price         float64   `db:"price"`
}
