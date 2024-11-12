package api_service

import "time"

type GetEmployeesOrdersCountResponse struct {
	EmployeeID   int64   `json:"employee_id"`
	PositionID   int     `json:"position_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	PositionName string  `json:"position_name"`
	Salary       float64 `json:"salary"`
	PhoneNumber  string  `json:"phone_number"`
	OrdersCount  int     `json:"orders_count"`
}

type GetMostPopularDishesResponse struct {
	DishID       int64   `json:"dish_id"`
	DishName     string  `json:"dish_name"`
	Price        float64 `json:"price"`
	OrderedCount int     `json:"ordered_count"`
}

type GetCustomerOrderHistoryResponse struct {
	CustomerID  int64   `json:"customer_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Discount    int     `json:"discount"`
	Orders      []Order `json:"orders"`
}

type Order struct {
	OrderID       int64     `json:"order_id"`
	CustomerID    int64     `json:"customer_id"`
	TransactionID int64     `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	Price         float64   `json:"price"`
}

type GetExpiringSoonProductsResponse struct {
	ProductID        int64  `db:"product_id"`
	Name             string `db:"name"`
	DateOfProduction string `db:"date_of_production"`
	DateOfExpiry     string `db:"date_of_expiry"`
	Status           string `db:"status"`
}
