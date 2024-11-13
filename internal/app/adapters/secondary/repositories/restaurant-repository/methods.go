package restaurant_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	api_service "github.com/mkorobovv/full-restaurant/internal/app/application/api-service"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/dish"
	"github.com/mkorobovv/full-restaurant/internal/app/domain/supplier"
)

func (repo *RestaurantRepository) GetExpiringProducts(ctx context.Context) (responses []api_service.GetExpiringSoonProductsResponse, err error) {
	query := `SELECT
    p.name,
	p.product_id,
	p.date_of_production,
    p.date_of_expiry,
    CASE
        WHEN p.date_of_expiry < NOW() THEN 'expired'
        WHEN p.date_of_expiry BETWEEN NOW() AND NOW() + INTERVAL '7 days' THEN 'expiring soon'
        END AS status
	FROM
    restaurant.products p
	WHERE
    p.date_of_expiry < NOW()
    OR p.date_of_expiry BETWEEN NOW() AND NOW() + INTERVAL '7 days';`

	err = repo.DB.SelectContext(ctx, &responses, query)

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (repo *RestaurantRepository) GetEmployeesOrdersCount(ctx context.Context, request api_service.GetEmployeesOrdersCountRequest) (responses []api_service.GetEmployeesOrdersCountResponse, err error) {
	query := `SELECT e.employee_id,
       e.position_id,
       e.first_name,
       e.last_name,
       p.position_name,
       p.salary,
       e.phone_number,
       COUNT(o.order_id) AS orders_count
	FROM restaurant.employees e
	JOIN restaurant.order_employee oe ON e.employee_id = oe.employee_id
	JOIN restaurant.orders o ON oe.order_id = o.order_id
	JOIN restaurant.positions p ON p.position_id = e.position_id
	WHERE o.created_at BETWEEN $1 AND $2
	GROUP BY e.employee_id, e.position_id, e.first_name, e.last_name, p.position_name, p.salary, e.phone_number;`

	err = repo.DB.SelectContext(ctx, &responses, query, request.DateFrom, request.DateTo)

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (repo *RestaurantRepository) GetDishesWithIngredients(ctx context.Context) (dishes []dish.Dish, err error) {
	query := `SELECT
    c.dish_id,
    c.price,
    c.name,
    (
        SELECT ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(o))) FROM (
            SELECT
                p.name,
                p.price_per_unit AS price,
                r.product_quantity AS quantity
            FROM restaurant.products p
            JOIN restaurant.recipes r ON p.product_id = r.product_id
            WHERE r.dish_id = c.dish_id
        ) o
    ) AS ingredients
FROM restaurant.dishes c;`

	err = repo.DB.SelectContext(ctx, &dishes, query)
	if err != nil {
		return nil, err
	}

	dishesResponse := make([]dish.Dish, len(dishes))

	for i, dish := range dishes {
		err = json.Unmarshal(dish.IngredientsBytes, &dish.Ingredients)
		if err != nil {
			return nil, err
		}

		dishesResponse[i] = dish
	}

	return dishesResponse, nil
}

func (repo *RestaurantRepository) GetDishesByIngredient(ctx context.Context, ingredient string) (dishes []dish.RecieveDish, err error) {
	query := `SELECT d.dish_id, d.name AS dish_name, d.price
 	FROM restaurant.dishes d
         JOIN restaurant.recipes r ON d.dish_id = r.dish_id
         JOIN restaurant.products p ON r.product_id = p.product_id
 	WHERE p.name LIKE $1;`

	err = repo.DB.SelectContext(ctx, &dishes, query, ingredient)

	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (repo *RestaurantRepository) GetMostPopularDishes(ctx context.Context) (responses []api_service.GetMostPopularDishesResponse, err error) {
	query := `SELECT d.name AS dish_name, d.price, d.dish_id, COUNT(od.order_dish_id) AS ordered_count
	FROM restaurant.dishes d
         JOIN restaurant.order_dish od ON d.dish_id = od.dish_id
	GROUP BY d.dish_id
	ORDER BY ordered_count DESC
	LIMIT 10;`

	err = repo.DB.SelectContext(ctx, &responses, query)

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (repo *RestaurantRepository) GetCustomerOrderHistory(ctx context.Context, request api_service.GetCustomerOrdersHistoryRequest) (response api_service.GetCustomerOrderHistoryResponse, err error) {
	query := `SELECT
    c.customer_id,
    c.first_name,
    c.last_name,
    c.phone_number,
    c.email,
    c.discount,
    (
        SELECT ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(o))) FROM (
            SELECT
                o.order_id,
                o.transaction_id,
                o.created_at,
                o.price,
                (
                    SELECT ARRAY_TO_JSON(ARRAY_AGG(ROW_TO_JSON(e))) FROM (
                        SELECT
                            d.name,
                            od.dish_quantity,
                            od.dish_price
                        FROM restaurant.dishes d
                        JOIN restaurant.order_dish od ON d.dish_id = od.dish_id
                        WHERE od.order_id = o.order_id
                    ) e
                ) AS dishes
            FROM restaurant.orders o
            WHERE o.customer_id = c.customer_id AND o.created_at BETWEEN $2 AND $3 
        ) o
    ) AS orders
	FROM restaurant.customers c
	WHERE c.customer_id = $1;`

	err = repo.DB.GetContext(ctx, &response, query, request.CustomerID, request.DateFrom, request.DateTo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api_service.GetCustomerOrderHistoryResponse{}, nil
		}

		return api_service.GetCustomerOrderHistoryResponse{}, err
	}

	response.Orders = make([]api_service.Order, 0)

	if len(response.OrdersBytes) == 0 {
		return response, nil
	}

	err = json.Unmarshal(response.OrdersBytes, &response.Orders)
	if err != nil {
		return api_service.GetCustomerOrderHistoryResponse{}, err
	}

	return response, nil
}

func (repo *RestaurantRepository) GetSuppliersByProduct(ctx context.Context, productName string) (suppliers []supplier.Supplier, err error) {
	query := `SELECT
    s.supplier_id,
    s.chief_name,
    s.company_name,
    s.email,
    s.address
	FROM restaurant.suppliers s
    JOIN restaurant.supplies sp ON sp.supplier_id = s.supplier_id
    JOIN restaurant.supply_product spd ON spd.supply_id = sp.supply_id
    JOIN restaurant.products p ON p.product_id = spd.product_id
	WHERE p.name LIKE $1`

	err = repo.DB.SelectContext(ctx, &suppliers, query, productName)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}
