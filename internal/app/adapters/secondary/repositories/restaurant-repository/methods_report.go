package restaurant_repository

import (
	"context"
	"time"

	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
)

type amountSupplyCostsDTO struct {
	TotalCost float64 `db:"total_cost"`
}

// GetAmountSupplyCosts - получить сумму затрат на поставки за определенный срок
func (repo *RestaurantRepository) GetAmountSupplyCosts(ctx context.Context, request api_service.CreateReportRequest) (amountSupplyCosts float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	const query = `
		SELECT SUM(price) AS total_cost
		FROM restaurant.supplies
		WHERE date_from >= $1 AND date_from <= $2;
	`

	var dto amountSupplyCostsDTO

	err = repo.DB.GetContext(ctx, &dto, query, request.DateFrom, request.DateTo)
	if err != nil {
		return 0.0, err
	}

	return dto.TotalCost, nil
}

type netProfitDTO struct {
	NetProfit float64 `db:"net_profit"`
}

// GetNetProfit - поиск чистой прибыли
func (repo *RestaurantRepository) GetNetProfit(ctx context.Context, request api_service.CreateReportRequest) (netProfit float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	const query = `
		WITH revenue AS (
			SELECT
			    SUM(o.price) AS total_revenue
			FROM restaurant.orders o
			JOIN restaurant.transactions t ON o.transaction_id = t.transaction_id
			WHERE o.created_at BETWEEN $1 AND $2
		),
		cost AS (
			SELECT SUM(s.price) AS total_cost
			FROM restaurant.supplies s
			WHERE s.date_from BETWEEN $1 AND $2
		)
		SELECT (r.total_revenue - c.total_cost) AS net_profit
		FROM revenue r, cost c;
	`

	var dto netProfitDTO

	err = repo.DB.GetContext(ctx, &dto, query, request.DateFrom, request.DateTo)
	if err != nil {
		return 0.0, err
	}

	return dto.NetProfit, nil
}

type avgOrderCheckDTO struct {
	AvgOrderValue float64 `db:"avg_order_value"`
}

// GetAverageOrderCheck - средний чек заказа
func (repo *RestaurantRepository) GetAverageOrderCheck(ctx context.Context, request api_service.CreateReportRequest) (avgOrderValue float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	const query = `
		SELECT AVG(price) AS avg_order_value
		FROM restaurant.orders
		WHERE created_at BETWEEN $1 AND $2;
	`

	var dto avgOrderCheckDTO

	err = repo.DB.GetContext(ctx, &dto, query, request.DateFrom, request.DateTo)
	if err != nil {
		return 0.0, err
	}

	return dto.AvgOrderValue, nil
}

type avgSupplyCheckDTO struct {
	AvgSupplyValue float64 `db:"avg_supply_value"`
}

// GetAverageSupplyCheck - средний чек поставки
func (repo *RestaurantRepository) GetAverageSupplyCheck(ctx context.Context, request api_service.CreateReportRequest) (avgSupplyValue float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	const query = `
		SELECT AVG(price) AS avg_supply_value
		FROM restaurant.supplies
		WHERE date_from BETWEEN $1 AND $2;
	`

	var dto avgSupplyCheckDTO

	err = repo.DB.GetContext(ctx, &dto, query, request.DateFrom, request.DateTo)
	if err != nil {
		return 0.0, err
	}

	return dto.AvgSupplyValue, nil
}

type lostRevenueDTO struct {
	LostRevenue float64 `db:"lost_revenue"`
}

// GetLostRevenue - поиск недополученной выручки из-за истечения срока годности продуктов
func (repo *RestaurantRepository) GetLostRevenue(ctx context.Context, request api_service.CreateReportRequest) (lostRevenue float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	const query = `
		SELECT SUM(p.price_per_unit * p.stock_quantity) AS lost_revenue
		FROM restaurant.products p
		WHERE p.date_of_expiry < $1;
	`

	var dto lostRevenueDTO

	err = repo.DB.GetContext(ctx, &dto, query, request.DateTo)
	if err != nil {
		return 0.0, err
	}

	return dto.LostRevenue, nil
}
