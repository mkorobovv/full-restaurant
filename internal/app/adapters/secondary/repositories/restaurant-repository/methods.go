package restaurant_repository

import (
	"context"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/supplier"
)

func (repo *RestaurantRepository) GetExpiringProducts(ctx context.Context) (responses []api_service.GetExpiringSoonProductsResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetEmployeesOrdersCount(ctx context.Context, request api_service.GetEmployeesOrdersCountRequest) (responses []api_service.GetEmployeesOrdersCountResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetDishesWithIngredients(ctx context.Context) (dishes []dish.Dish, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetDishesByIngredient(ctx context.Context, ingredient string) (dishes []dish.Dish, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetMostPopularDishes(ctx context.Context) (responses []api_service.GetMostPopularDishesResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetCustomerOrderHistory(ctx context.Context, customerID int64) (response api_service.GetCustomerOrderHistoryResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *RestaurantRepository) GetSuppliersByProduct(ctx context.Context, productName string) (suppliers []supplier.Supplier, err error) {
	//TODO implement me
	panic("implement me")
}
