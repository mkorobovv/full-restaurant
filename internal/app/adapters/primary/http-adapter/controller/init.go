package controller

import (
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	print_form_service "github.com/mkorobovv/full-restaurant/internal/app/application/print-form-service"
)

type Controller struct {
	apiService       *api_service.APIService
	printFormService *print_form_service.PrintFormService
}

func New(
	apiSvc *api_service.APIService,
	printFormService *print_form_service.PrintFormService,
) *Controller {
	return &Controller{
		apiService:       apiSvc,
		printFormService: printFormService,
	}
}
