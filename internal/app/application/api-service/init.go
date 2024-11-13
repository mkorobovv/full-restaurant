package api_service

import (
	"context"

	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/supplier"
)

type APIService struct {
	restaurantRepository restaurantRepository
}

type restaurantRepository interface {
	// Supplies

	GetExpiringProducts(ctx context.Context) (products []GetExpiringSoonProductsResponse, err error)

	// Employees

	GetEmployeesOrdersCount(ctx context.Context, request GetEmployeesOrdersCountRequest) (responses []GetEmployeesOrdersCountResponse, err error)

	// Dishes

	GetDishesWithIngredients(ctx context.Context) (dishes []dish.Dish, err error)
	GetDishesByIngredient(ctx context.Context, ingredient string) (dishes []dish.Dish, err error)
	GetMostPopularDishes(ctx context.Context) (responses []GetMostPopularDishesResponse, err error)

	// Customers

	GetCustomerOrderHistory(ctx context.Context, customerID int64) (response GetCustomerOrderHistoryResponse, err error)

	// Suppliers

	GetSuppliersByProduct(ctx context.Context, productName string) (suppliers []supplier.Supplier, err error)
}

func New(repo restaurantRepository) *APIService {
	return &APIService{
		restaurantRepository: repo,
	}
}
