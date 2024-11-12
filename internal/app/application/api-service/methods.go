package api_service

import (
	"context"

	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/supplier"
)

func (svc *APIService) GetExpiringProducts(ctx context.Context) (products []GetExpiringSoonProductsResponse, err error) {
	products, err = svc.restaurantRepository.GetExpiringProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (svc *APIService) GetEmployeesOrdersCount(ctx context.Context, request GetEmployeesOrdersCountRequest) (employees []GetEmployeesOrdersCountResponse, err error) {
	employees, err = svc.restaurantRepository.GetEmployeesOrdersCount(ctx, request)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (svc *APIService) GetMostPopularDishes(ctx context.Context) (dishes []GetMostPopularDishesResponse, err error) {
	dishes, err = svc.restaurantRepository.GetMostPopularDishes(ctx)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (svc *APIService) GetDishesWithIngredients(ctx context.Context) (dishes []dish.Dish, err error) {
	dishes, err = svc.restaurantRepository.GetDishesWithIngredients(ctx)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (svc *APIService) GetDishesByIngredients(ctx context.Context, ingredient string) (dishes []dish.Dish, err error) {
	dishes, err = svc.restaurantRepository.GetDishesByIngredient(ctx, ingredient)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (svc *APIService) GetCustomerOrderHistory(ctx context.Context, customerID int64) (history GetCustomerOrderHistoryResponse, err error) {
	history, err = svc.restaurantRepository.GetCustomerOrderHistory(ctx, customerID)
	if err != nil {
		return GetCustomerOrderHistoryResponse{}, err
	}

	return history, nil
}

func (svc *APIService) GetSuppliersByProduct(ctx context.Context, productName string) (suppliers []supplier.Supplier, err error) {
	suppliers, err = svc.restaurantRepository.GetSuppliersByProduct(ctx, productName)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}
