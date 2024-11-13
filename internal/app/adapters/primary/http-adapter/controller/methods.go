package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	controller_gen "github.com/mkorobovv/full-restaurant/internal/app/adapters/primary/http-adapter/controller-gen"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
)

func (ctr *Controller) GetCustomerOrderHistory(w http.ResponseWriter, r *http.Request, customerId int) {
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

	req := api_service.GetCustomerOrdersHistoryRequest{
		CustomerID: customerId,
		DateFrom:   dateFrom,
		DateTo:     dateTo,
	}

	history, err := ctr.apiService.GetCustomerOrderHistory(r.Context(), req)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if history.CustomerID == 0 {
		err = errors.New("customer not found")

		writeErr(w, err.Error(), http.StatusNotFound)

		return
	}

	dto := GetCustomerOrderHistoryDTO{
		CustomerID:  history.CustomerID,
		FirstName:   history.FirstName,
		LastName:    history.LastName,
		PhoneNumber: history.PhoneNumber,
		Email:       history.Email,
		Discount:    history.Discount,
		Orders:      history.Orders,
	}

	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) GetDishesWithIngredients(w http.ResponseWriter, r *http.Request) {
	dishes, err := ctr.apiService.GetDishesWithIngredients(r.Context())
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	dtos := make([]GetDishesWithIngredientsDTO, len(dishes))

	for i, dish := range dishes {
		dtos[i] = GetDishesWithIngredientsDTO{
			DishID:      dish.DishID,
			DishName:    dish.DishName,
			Price:       dish.Price,
			Ingredients: dish.Ingredients,
		}
	}

	err = json.NewEncoder(w).Encode(dtos)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) GetDishesByIngredients(w http.ResponseWriter, r *http.Request) {
	ingredient := r.URL.Query().Get("ingredient")
	if len(ingredient) == 0 {
		writeErr(w, errors.New("ingredient mustn't be empty").Error(), http.StatusBadRequest)

		return
	}

	dishes, err := ctr.apiService.GetDishesByIngredients(r.Context(), ingredient)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(dishes) == 0 {
		writeErr(w, errors.New("dish not found").Error(), http.StatusNotFound)

		return
	}

	err = json.NewEncoder(w).Encode(dishes)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}
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

func (ctr *Controller) GetSuppliersByProduct(w http.ResponseWriter, r *http.Request, productName string) {
	suppliers, err := ctr.apiService.GetSuppliersByProduct(r.Context(), productName)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(suppliers)
	if err != nil {
		writeErr(w, err.Error(), http.StatusInternalServerError)

		return
	}

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
