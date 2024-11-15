package api_service

import "time"

type GetEmployeesOrdersCountResponse struct {
	EmployeeID   int64   `json:"employee_id" db:"employee_id"`
	PositionID   int     `json:"position_id" db:"position_id"`
	FirstName    string  `json:"first_name" db:"first_name"`
	LastName     string  `json:"last_name" db:"last_name"`
	PositionName string  `json:"position_name" db:"position_name"`
	Salary       float64 `json:"salary" db:"salary"`
	PhoneNumber  string  `json:"phone_number" db:"phone_number"`
	OrdersCount  int     `json:"orders_count" db:"orders_count"`
}

type GetMostPopularDishesResponse struct {
	DishID       int64   `json:"dish_id" db:"dish_id"`
	DishName     string  `json:"dish_name" db:"dish_name"`
	Price        float64 `json:"price" db:"price"`
	OrderedCount int     `json:"ordered_count" db:"ordered_count"`
}

type GetCustomerOrderHistoryResponse struct {
	CustomerID  int64   `json:"customer_id" db:"customer_id"`
	FirstName   string  `json:"first_name"  db:"first_name"`
	LastName    string  `json:"last_name"   db:"last_name"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
	Email       string  `json:"email" db:"email"`
	Discount    int     `json:"discount" db:"discount"`
	OrdersBytes []byte  `db:"orders"`
	Orders      []Order `json:"orders"`
}

type Dish struct {
	DishName string  `json:"name"`
	Quantity int     `json:"dish_quantity"`
	Price    float64 `json:"dish_price"`
}

type Order struct {
	OrderID       int64   `json:"order_id"`
	CustomerID    int64   `json:"customer_id"`
	TransactionID int64   `json:"transaction_id"`
	CreatedAt     string  `json:"created_at"`
	Price         float64 `json:"price"`
	Dishes        []Dish  `json:"dishes"`
}

type GetExpiringSoonProductsResponse struct {
	ProductID        int64  `db:"product_id" json:"product_id"`
	Name             string `db:"name" json:"name"`
	DateOfProduction string `db:"date_of_production" json:"date_of_production"`
	DateOfExpiry     string `db:"date_of_expiry" json:"date_of_expiry"`
	Status           string `db:"status" json:"status"`
}

type Report struct {
	LostRevenue        float64   `json:"lost_revenue"`
	NetProfit          float64   `json:"net_profit"`
	AverageSupplyCheck float64   `json:"average_supply_check"`
	AverageOrderCheck  float64   `json:"average_order_check"`
	AmountSupplyCosts  float64   `json:"amount_supply_costs"`
	DateFrom           time.Time `json:"date_from"`
	DateTo             time.Time `json:"date_to"`
}
