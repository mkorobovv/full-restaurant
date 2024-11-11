
-- 1. **Поиск сотрудников по определенным позициям с зарплатой выше заданного значения.** ✅
SELECT e.first_name, e.last_name, p.position_name, p.salary, p.phone_number
FROM restaurant.employees e
         JOIN restaurant.positions p ON e.position_id = p.position_id
WHERE p.salary > $1;  -- Замените $1 на заданное значение

-- 2. **Поиск блюд, содержащих определенные ингредиенты.** ✅
SELECT d.name AS dish_name, p.name AS product_name
FROM restaurant.dishes d
         JOIN restaurant.recipes r ON d.dish_id = r.dish_id
         JOIN restaurant.products p ON r.product_id = p.product_id
WHERE p.name LIKE 'Свекла';  -- Замените $1 на ингредиент

-- 3. **Поиск активных поставщиков с определенным количеством на складе.**✅
SELECT s.company_name, p.name, p.stock_quantity
FROM restaurant.suppliers s
         JOIN restaurant.supplies sp ON s.supplier_id = sp.supplier_id
         JOIN restaurant.supply_product spd ON sp.supply_id = spd.supply_id
         JOIN restaurant.products p ON spd.product_id = p.product_id
WHERE p.stock_quantity > $1;  -- Замените $1 на заданное количество

-- 4. **Поиск продуктов с истекающим сроком годности.**✅
SELECT p.product_id, p.name, p.date_of_expiry
FROM restaurant.products p
WHERE p.date_of_expiry BETWEEN NOW() AND NOW() + INTERVAL '7 days';  -- Продукты с истекающим сроком годности в ближайшую неделю

-- 5. **Сумма затрат на поставки за месяц.**✅
SELECT SUM(price) AS total_cost
FROM restaurant.supplies
WHERE date_from >= DATE_TRUNC('month', NOW()) AND date_from < DATE_TRUNC('month', NOW()) + INTERVAL '1 month';

-- 6. **Список блюд и их ингредиентов с указанием количества.**✅
SELECT d.name AS dish_name, p.name AS product_name, r.product_quantity
FROM restaurant.dishes d
         JOIN restaurant.recipes r ON d.dish_id = r.dish_id
         JOIN restaurant.products p ON r.product_id = p.product_id;

-- 7. **Сколько заказов выполнил каждый сотрудник за определенное время.**✅
SELECT e.first_name, e.last_name, COUNT(o.order_id) AS order_count
FROM restaurant.employees e
         JOIN restaurant.order_employee oe ON e.employee_id = oe.employee_id
         JOIN restaurant.orders o ON oe.order_id = o.order_id
WHERE o.created_at BETWEEN '2023-10-01' AND '2023-10-31'  -- Укажите даты в формате 'YYYY-MM-DD'
GROUP BY e.employee_id;

-- 8. **Поиск чистой прибыли.**✅
WITH revenue AS (
    SELECT SUM(o.price) AS total_revenue
    FROM restaurant.orders o
             JOIN restaurant.transactions t ON o.transaction_id = t.transaction_id
),
     cost AS (
         SELECT SUM(s.price) AS total_cost
         FROM restaurant.supplies s
     )
SELECT (r.total_revenue - c.total_cost) AS net_profit
FROM revenue r, cost c;

-- 9 **Количество поставок определенного поставщика за определенное время.**✅
SELECT s.company_name, COUNT(sp.supply_id) AS supply_count
FROM restaurant.suppliers s
         JOIN restaurant.supplies sp ON s.supplier_id = sp.supplier_id
WHERE sp.date_from BETWEEN '2023-11-15' AND '2023-10-15'  -- Замените $1 и $2 на временной диапазон
GROUP BY s.supplier_id;

-- 10. **Средняя стоимость поставки за определенное время.**✅
SELECT AVG(price) AS avg_supply_cost
FROM restaurant.supplies
WHERE date_from BETWEEN '2023-10-15' AND '2023-11-15';  -- Замените $1 и $2 на временной диапазон

-- 11. **Средний чек за определенное время.**✅
SELECT AVG(price) AS avg_order_value
FROM restaurant.orders
WHERE created_at BETWEEN '2023-10-15' AND '2023-11-15';  -- Замените $1 и $2 на временной диапазон

-- 12. **Список самых популярных блюд.**✅
SELECT d.name AS dish_name, COUNT(od.order_dish_id) AS dish_count
FROM restaurant.dishes d
         JOIN restaurant.order_dish od ON d.dish_id = od.dish_id
GROUP BY d.dish_id
ORDER BY dish_count DESC
LIMIT 10;  -- Показывает топ-10 самых популярных блюд

-- 13. **Поиск недополученной выручки из-за истечения срока годности продуктов**✅
SELECT SUM(p.price_per_unit * p.stock_quantity) AS lost_revenue
FROM restaurant.products p
WHERE p.date_of_expiry < '2023-11-15';

-- 14. Сложный расчет рентабельности блюд с учетом затрат на продукты
-- Для каждого блюда рассчитывается себестоимость продуктов и рентабельность в процентах.

WITH dish_cost AS (
    SELECT d.dish_id, SUM(p.price_per_unit * r.product_quantity) AS total_cost
    FROM restaurant.dishes d
             JOIN restaurant.recipes r ON d.dish_id = r.dish_id
             JOIN restaurant.products p ON r.product_id = p.product_id
    GROUP BY d.dish_id
),
     dish_revenue AS (
         SELECT d.dish_id, SUM(od.dish_price * od.dish_quantity) AS total_revenue
         FROM restaurant.dishes d
                  JOIN restaurant.order_dish od ON d.dish_id = od.dish_id
         GROUP BY d.dish_id
     )
SELECT d.name AS dish_name,
       dc.total_cost,
       dr.total_revenue,
       ROUND((dr.total_revenue - dc.total_cost) * 100.0 / NULLIF(dc.total_cost, 0), 2) AS profit_margin
FROM dish_cost dc
         JOIN dish_revenue dr ON dc.dish_id = dr.dish_id
         JOIN restaurant.dishes d ON dc.dish_id = d.dish_id
ORDER BY profit_margin DESC;

--15 Чтобы получить список ингредиентов для каждого блюда в формате JSON (ARRAY_TO_JSON)
SELECT d.name AS dish_name,
       ARRAY_TO_JSON(ARRAY_AGG(json_build_object('name', p.name, 'price', p.price_per_unit, 'quantity', r.product_quantity))) AS ingredients
FROM restaurant.dishes d
         JOIN restaurant.recipes r ON d.dish_id = r.dish_id
         JOIN restaurant.products p ON r.product_id = p.product_id
GROUP BY d.dish_id;



-- 16 Поиск сотрудников, у которых есть заказы: (EXISTS)
SELECT e.first_name, e.last_name
FROM restaurant.employees e
WHERE EXISTS (
    SELECT 1 FROM restaurant.order_employee oe WHERE oe.employee_id = e.employee_id
);

-- 17 Поиск блюд, для которых средняя стоимость ингредиентов выше заданного значения (вложеность)
SELECT name
FROM restaurant.dishes
WHERE dish_id IN (
    SELECT r.dish_id
    FROM restaurant.recipes r
             JOIN restaurant.products p ON r.product_id = p.product_id
    GROUP BY r.dish_id
    HAVING AVG(p.price_per_unit * r.product_quantity) > 100
);

-- 18 Cколько составляют доходы от продаж и расходы на поставки за текущий месяц, а также общую чистую прибыль (UNION)
WITH monthly_data AS (
    -- Доходы от продаж
    SELECT 'Revenue' AS type, SUM(o.price) AS amount
    FROM restaurant.orders o
    WHERE o.created_at >= DATE_TRUNC('month', DATE '2023-10-27')
      AND o.created_at < DATE_TRUNC('month', DATE '2023-10-27') + INTERVAL '1 month'

    UNION ALL

    -- Расходы на поставки
    SELECT 'Cost' AS type, SUM(s.price) AS amount
    FROM restaurant.supplies s
    WHERE s.date_from >= DATE_TRUNC('month', DATE '2023-10-27')
      AND s.date_from < DATE_TRUNC('month', DATE '2023-10-27') + INTERVAL '1 month'
)

-- Суммируем по каждому типу и рассчитываем чистую прибыль
SELECT
    type,
    amount,
    CASE
        WHEN type = 'Revenue' THEN amount
        WHEN type = 'Cost' THEN -amount
        ELSE 0
        END AS signed_amount
FROM
    monthly_data

UNION ALL

-- Рассчитываем чистую прибыль
SELECT
    'Net Profit' AS type,
    SUM(CASE WHEN type = 'Revenue' THEN amount WHEN type = 'Cost' THEN -amount END) AS amount,
    NULL AS signed_amount
FROM
    monthly_data;


-- 19 Поиск сотрудников, чья зарплата выше любой зарплаты среди сотрудников определенной позиции (например, "Кассир")
SELECT e.first_name, e.last_name, p.position_name, p.salary
FROM restaurant.employees e
         JOIN restaurant.positions p ON e.position_id = p.position_id
WHERE e.salary > ANY (
    SELECT salary
    FROM restaurant.employees
             JOIN restaurant.positions ON employees.position_id = positions.position_id
    WHERE position_name = 'Кассир'
);

--20 Поиск блюд с ценой выше стоимости всех блюд: (Запросы на связанные подзапросы)
SELECT name, price
FROM restaurant.dishes
WHERE price > (SELECT AVG(price) FROM restaurant.dishes);
