package controller

import (
	"encoding/json"
	controller_gen "github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/controller-gen"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"log"
	"net/http"
	"time"
)

func (ctr *Controller) GetCustomerOrderHistory(w http.ResponseWriter, r *http.Request, customerId int) {

	return
}

func (ctr *Controller) GetDishesWithIngredients(w http.ResponseWriter, r *http.Request) {
	dishes, err := ctr.apiService.GetDishesWithIngredients(r.Context())
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(dishes)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) GetDishesByIngredients(w http.ResponseWriter, r *http.Request) {
	return
}

func (ctr *Controller) GetEmployeeOrderCount(w http.ResponseWriter, r *http.Request) {
	dateFromQueryParam := r.URL.Query().Get("date_from")
	dateToQueryParam := r.URL.Query().Get("date_to")

	dateFrom, err := time.Parse(time.DateOnly, dateFromQueryParam)
	if err != nil {
		writeErr(w, err.Error(), http.StatusBadRequest)

		return
	}

	dateTo, err := time.Parse(time.DateOnly, dateToQueryParam)
	if err != nil {
		writeErr(w, err.Error(), http.StatusBadRequest)

		return
	}

	req := api_service.GetEmployeesOrdersCountRequest{
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}

	employees, err := ctr.apiService.GetEmployeesOrdersCount(r.Context(), req)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) GetExpiringProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ctr.apiService.GetExpiringProducts(r.Context())
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) GetSuppliersForProduct(w http.ResponseWriter, r *http.Request, productName int) {
	return
}

func (ctr *Controller) DownloadReport(w http.ResponseWriter, r *http.Request) {
	return
}

func (ctr *Controller) GetMostPopularDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := ctr.apiService.GetMostPopularDishes(r.Context())
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(dishes)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func writeErr(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(http.StatusInternalServerError)

	errResponse := controller_gen.Error{
		Code:    statusCode,
		Message: errMsg,
	}

	err := json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return
}
