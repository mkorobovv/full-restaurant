package product

import "time"

type Product struct {
	ProductID        int64     `db:"product_id"`
	Name             string    `db:"name"`
	StockQuantity    int       `db:"stock_quantity"`
	OkeiID           string    `db:"okei_id"`
	PricePerUnit     float64   `db:"price_per_unit"`
	DateOfProduction time.Time `db:"date_of_production"`
	DateOfExpiry     time.Time `db:"date_of_expiry"`
}
