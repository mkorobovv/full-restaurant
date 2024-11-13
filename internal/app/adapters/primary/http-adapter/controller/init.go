package controller

import (
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
)

type Controller struct {
	apiService *api_service.APIService
}

func New(
	apiSvc *api_service.APIService,
) *Controller {
	return &Controller{
		apiService: apiSvc,
	}
}
