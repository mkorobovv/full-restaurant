package api_service

import (
	"context"

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
