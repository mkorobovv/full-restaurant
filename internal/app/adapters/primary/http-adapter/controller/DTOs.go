package controller

import (
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
)

type GetCustomerOrderHistoryDTO struct {
	CustomerID  int64               `json:"customer_id" db:"customer_id"`
	FirstName   string              `json:"first_name"  db:"first_name"`
	LastName    string              `json:"last_name"   db:"last_name"`
	PhoneNumber string              `json:"phone_number" db:"phone_number"`
	Email       string              `json:"email" db:"email"`
	Discount    int                 `json:"discount" db:"discount"`
	Orders      []api_service.Order `json:"orders"`
}

type GetDishesWithIngredientsDTO struct {
	DishID      int64             `db:"dish_id" json:"dish_id"`
	DishName    string            `db:"name" json:"name"`
	Price       float64           `db:"price" json:"price"`
	Ingredients []dish.Ingredient `json:"ingredients"`
}
