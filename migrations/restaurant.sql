CREATE SCHEMA restaurant;

CREATE TABLE restaurant.positions (
    position_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    position_name VARCHAR(50),
    salary NUMERIC(10, 2)
);

CREATE TABLE restaurant.employees (
    employee_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    position_id BIGINT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(12),

    FOREIGN KEY (position_id) REFERENCES restaurant.positions(position_id) MATCH FULL
);

CREATE TABLE restaurant.customers (
    customer_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(12),
    email VARCHAR(80),
    discount INT
);

CREATE TABLE restaurant.transactions (
    transaction_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    transaction_type VARCHAR(70),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    price NUMERIC(10, 2)
);

CREATE TABLE restaurant.orders (
    order_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id BIGINT,
    transaction_id BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    price NUMERIC(10, 2),

    FOREIGN KEY (customer_id) REFERENCES restaurant.customers(customer_id) MATCH FULL,
    FOREIGN KEY (transaction_id) REFERENCES restaurant.transactions(transaction_id) MATCH FULL
);

CREATE TABLE restaurant.dishes (
    dish_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(80),
    price NUMERIC(10, 2)
);

CREATE TABLE restaurant.products (
    product_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50),
    stock_quantity INT,
    okei_id VARCHAR(20),
    price_per_unit NUMERIC,
    date_of_production DATE,
    date_of_expiry DATE
);

CREATE TABLE restaurant.recipes (
    recipe_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id BIGINT,
    product_quantity INT,
    dish_id BIGINT,

    FOREIGN KEY (product_id) REFERENCES restaurant.products(product_id) MATCH FULL,
    FOREIGN KEY (dish_id) REFERENCES restaurant.dishes(dish_id) MATCH FULL
);

CREATE TABLE restaurant.suppliers (
    supplier_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_name VARCHAR(50),
    chief_name VARCHAR(50),
    email VARCHAR(50),
    address VARCHAR(50)
);

CREATE TABLE restaurant.supplies (
    supply_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    supplier_id BIGINT,
    transaction_id BIGINT,
    date_from TIMESTAMP,
    price NUMERIC,

    FOREIGN KEY (supplier_id) REFERENCES restaurant.suppliers(supplier_id) MATCH FULL,
    FOREIGN KEY (transaction_id) REFERENCES restaurant.transactions(transaction_id) MATCH FULL
);

CREATE TABLE restaurant.supply_product (
    supply_product_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    quantity INT,
    product_id BIGINT,
    supply_id BIGINT,

    FOREIGN KEY (product_id) REFERENCES restaurant.products(product_id) MATCH FULL,
    FOREIGN KEY (supply_id) REFERENCES restaurant.supplies(supply_id) MATCH FULL
);

CREATE TABLE restaurant.order_dish (
    order_dish_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    dish_id BIGINT,
    order_id BIGINT,
    dish_quantity INT,
    dish_price NUMERIC(10, 2),

    FOREIGN KEY (dish_id) REFERENCES restaurant.dishes(dish_id) MATCH FULL,
    FOREIGN KEY (order_id) REFERENCES restaurant.orders(order_id) MATCH FULL
);

CREATE TABLE restaurant.payments (
    payment_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    employee_id BIGINT NOT NULL,
    transaction_id BIGINT NOT NULL,
    date_from TIMESTAMP WITHOUT TIME ZONE,
    bonus INT,

    FOREIGN KEY (employee_id) REFERENCES restaurant.employees(employee_id) MATCH FULL,
    FOREIGN KEY (transaction_id) REFERENCES restaurant.transactions(transaction_id) MATCH FULL
);

CREATE TABLE restaurant.order_employee (
    order_employee_id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    employee_id BIGINT,
    order_id BIGINT,

    FOREIGN KEY (employee_id) REFERENCES restaurant.employees(employee_id) MATCH FULL,
    FOREIGN KEY (order_id) REFERENCES restaurant.orders(order_id) MATCH FULL
);