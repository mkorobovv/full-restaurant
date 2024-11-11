package api_service

import (
	"context"

	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/product"
)

func (svc *APIService) GetExpiringProducts(ctx context.Context) (products []product.Product, err error) {
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
