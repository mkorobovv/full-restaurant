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
	GetDishesByIngredient(ctx context.Context, ingredient string) (dishes []dish.RecieveDish, err error)
	GetMostPopularDishes(ctx context.Context) (responses []GetMostPopularDishesResponse, err error)
	GetUnorderedDishes(ctx context.Context) (dishes []dish.RecieveDish, err error)

	// Customers

	GetCustomerOrderHistory(ctx context.Context, request GetCustomerOrdersHistoryRequest) (response GetCustomerOrderHistoryResponse, err error)

	// Suppliers

	GetSuppliersByProduct(ctx context.Context, productName string) (suppliers []supplier.Supplier, err error)

	// Report

	GetAverageSupplyCheck(ctx context.Context, request CreateReportRequest) (avgSupplyValue float64, err error)
	GetAverageOrderCheck(ctx context.Context, request CreateReportRequest) (avgOrderValue float64, err error)
	GetLostRevenue(ctx context.Context, request CreateReportRequest) (lostRevenue float64, err error)
	GetNetProfit(ctx context.Context, request CreateReportRequest) (netProfit float64, err error)
	GetAmountSupplyCosts(ctx context.Context, request CreateReportRequest) (amountSupplyCosts float64, err error)
}

func New(repo restaurantRepository) *APIService {
	return &APIService{
		restaurantRepository: repo,
	}
}
