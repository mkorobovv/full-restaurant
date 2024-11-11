package dish

import "time"

type Dish struct {
	DishID      int64        `db:"dish_id"`
	DishName    string       `db:"dish_name"`
	Price       float64      `db:"price"`
	Ingredients []Ingredient `db:"ingredients"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Product struct {
	ProductID        int64     `db:"product_id"`
	Name             string    `db:"name"`
	StockQuantity    int       `db:"stock_quantity"`
	OkeiID           string    `db:"okei_id"`
	PricePerUnit     float64   `db:"price_per_unit"`
	DateOfProduction time.Time `db:"date_of_production"`
	DateOfExpiry     time.Time `db:"date_of_expiry"`
}
