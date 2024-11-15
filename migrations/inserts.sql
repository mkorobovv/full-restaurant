
-- Тестовые данные для таблицы positions
INSERT INTO restaurant.positions (position_name, salary) VALUES
	('Менеджер', 60000.00),
	('Официант', 35000.00),
	('Шеф-повар', 70000.00),
	('Кухонный работник', 30000.00),
	('Бармен', 40000.00),
	('Уборщик', 25000.00),
	('Кассир', 30000.00),
	('Сомелье', 50000.00),
	('Директор', 80000.00),
	('Помощник повара', 28000.00);

-- Тестовые данные для таблицы employees
INSERT INTO restaurant.employees (position_id, first_name, last_name, phone_number) VALUES
    (1, 'Александр', 'Смирнов', '79161234567'),
    (2, 'Елена', 'Иванова', '79261234568'),
    (3, 'Сергей', 'Кузнецов', '79361234569'),
    (4, 'Мария', 'Попова', '79461234570'),
    (5, 'Дмитрий', 'Семенов', '79561234571'),
    (6, 'Анна', 'Лебедева', '79661234572'),
    (7, 'Алексей', 'Ковалев', '79761234573'),
    (8, 'Наталья', 'Григорьева', '79861234574'),
    (9, 'Игорь', 'Зайцев', '79961234575'),
    (10, 'Ольга', 'Смирнова', '79061234576'),
    (1, 'Петр', 'Николаев', '79171234577'),
    (2, 'Светлана', 'Тихонова', '79271234578'),
    (3, 'Станислав', 'Федоров', '79371234579'),
    (4, 'Татьяна', 'Сидорова', '79471234580'),
    (5, 'Виктор', 'Котов', '79571234581'),
    (6, 'Вероника', 'Петрова', '79671234582'),
    (7, 'Роман', 'Денисов', '79771234583'),
    (8, 'Ирина', 'Сафонова', '79871234584'),
    (9, 'Денис', 'Калашников', '79971234585'),
    (10, 'Анастасия', 'Морозова', '79071234586'),
    (1, 'Кирилл', 'Егоров', '79181234587'),
    (2, 'Яна', 'Романовская', '79281234588'),
    (3, 'Артем', 'Дмитриев', '79381234589'),
    (4, 'Никита', 'Гаврилов', '79481234590'),
    (5, 'Ксения', 'Захарова', '79581234591'),
    (6, 'Максим', 'Фролов', '79681234592'),
    (7, 'Людмила', 'Шмидт', '79781234593'),
    (8, 'Олег', 'Трофимов', '79881234594'),
    (9, 'Евгения', 'Кузьмина', '79981234595'),
    (10, 'Станислав', 'Мартынов', '79081234596');

-- Тестовые данные для таблицы customers
INSERT INTO restaurant.customers (first_name, last_name, phone_number, email, discount) VALUES
    ('Александр', 'Смирнов', '79161234567', 'alexander.smirnov@example.com', 10),
    ('Елена', 'Иванова', '79261234568', 'elena.ivanova@example.com', 5),
    ('Сергей', 'Кузнецов', '79361234569', 'sergey.kuznetsov@example.com', 15),
    ('Мария', 'Попова', '79461234570', 'maria.popova@example.com', 20),
    ('Дмитрий', 'Семенов', '79561234571', 'dmitry.semenov@example.com', 0),
    ('Анна', 'Лебедева', '79661234572', 'anna.lebedeva@example.com', 10),
    ('Алексей', 'Ковалев', '79761234573', 'aleksey.kovalev@example.com', 5),
    ('Наталья', 'Григорьева', '79861234574', 'natalia.grigorieva@example.com', 15),
    ('Игорь', 'Зайцев', '79961234575', 'igor.zaytsev@example.com', 10),
    ('Ольга', 'Смирнова', '79061234576', 'olga.smirnova@example.com', 0),
    ('Петр', 'Николаев', '79171234577', 'petr.nikolaev@example.com', 5),
    ('Светлана', 'Тихонова', '79271234578', 'svetlana.tikhonova@example.com', 15),
    ('Станислав', 'Федоров', '79371234579', 'stanislav.fedorov@example.com', 10),
    ('Татьяна', 'Сидорова', '79471234580', 'tatiana.sidorova@example.com', 20),
    ('Виктор', 'Котов', '79571234581', 'viktor.kotov@example.com', 0),
    ('Вероника', 'Петрова', '79671234582', 'veronika.petrovа@example.com', 5),
    ('Роман', 'Денисов', '79771234583', 'roman.denisov@example.com', 10),
    ('Ирина', 'Сафонова', '79871234584', 'irina.safonova@example.com', 15),
    ('Денис', 'Калашников', '79971234585', 'denis.kalashnikov@example.com', 20),
    ('Анастасия', 'Морозова', '79071234586', 'anastasia.morozova@example.com', 10),
    ('Кирилл', 'Егоров', '79181234587', 'kirill.egarov@example.com', 5),
    ('Яна', 'Романовская', '79281234588', 'yana.romanovskaya@example.com', 15),
    ('Артем', 'Дмитриев', '79381234589', 'artem.dmitriev@example.com', 20),
    ('Никита', 'Гаврилов', '79481234590', 'nikita.gavrilov@example.com', 0),
    ('Ксения', 'Захарова', '79581234591', 'kseniya.zaharova@example.com', 5),
    ('Максим', 'Фролов', '79681234592', 'maxim.frolov@example.com', 10),
    ('Людмила', 'Шмидт', '79781234593', 'lyudmila.shmidt@example.com', 15),
    ('Олег', 'Трофимов', '79881234594', 'oleg.trofimov@example.com', 20),
    ('Евгения', 'Кузьмина', '79981234595', 'evgenia.kuzmina@example.com', 0),
    ('Станислав', 'Мартынов', '79081234596', 'stanislav.martynov@example.com', 5);

-- Тестовые данные для таблицы transactions
INSERT INTO restaurant.transactions (transaction_type, created_at, price) VALUES
    ('Оплата заказа', '2023-10-01 12:30:45', 1500.00),
    ('Оплата поставки', '2023-10-02 14:15:20', 2300.50),
    ('Оплата заказа', '2023-10-03 18:05:30', 870.00),
    ('Оплата бонусов сотрудникам', '2023-10-04 11:25:15', 5000.00),
    ('Оплата заказа', '2023-10-05 16:45:55', 1200.00),
    ('Оплата поставки', '2023-10-06 13:10:10', 3400.00),
    ('Оплата заказа', '2023-10-07 15:50:00', 950.00),
    ('Оплата заказа', '2023-10-08 17:35:22', 2100.75),
    ('Оплата бонусов сотрудникам', '2023-10-09 09:40:45', 4500.00),
    ('Оплата поставки', '2023-10-10 12:20:55', 1800.00),
    ('Оплата заказа', '2023-10-11 19:25:32', 1600.00),
    ('Оплата заказа', '2023-10-12 10:30:15', 750.00),
    ('Оплата поставки', '2023-10-13 11:50:50', 2700.00),
    ('Оплата заказа', '2023-10-14 14:05:25', 1850.00),
    ('Оплата бонусов сотрудникам', '2023-10-15 09:15:30', 5200.00),
    ('Оплата заказа', '2023-10-16 20:20:20', 930.00),
    ('Оплата заказа', '2023-10-17 18:35:45', 1450.00),
    ('Оплата поставки', '2023-10-18 13:25:40', 2200.00),
    ('Оплата заказа', '2023-10-19 17:10:10', 1300.00),
    ('Оплата бонусов сотрудникам', '2023-10-20 11:05:05', 4100.00),
    ('Оплата заказа', '2023-10-21 16:50:20', 1100.00),
    ('Оплата поставки', '2023-10-22 15:40:15', 2400.00),
    ('Оплата заказа', '2023-10-23 18:10:05', 1400.00),
    ('Оплата заказа', '2023-10-24 19:25:50', 760.00),
    ('Оплата бонусов сотрудникам', '2023-10-25 10:15:15', 5300.00),
    ('Оплата заказа', '2023-10-26 14:05:35', 900.00),
    ('Оплата заказа', '2023-10-27 20:30:40', 1150.00),
    ('Оплата поставки', '2023-10-28 12:45:25', 2900.00),
    ('Оплата заказа', '2023-10-29 17:55:55', 1550.00),
    ('Оплата бонусов сотрудникам', '2023-10-30 09:35:00', 4800.00);

-- Тестовые данные для таблицы dishes
INSERT INTO restaurant.dishes (name, price) VALUES
    ('Борщ', 300.00),
    ('Цезарь с курицей', 450.00),
    ('Картофельное пюре с котлетой', 350.00),
    ('Салат Оливье', 280.00),
    ('Куриный суп', 320.00),
    ('Паста Карбонара', 500.00),
    ('Стейк', 1200.00),
    ('Сырники', 250.00),
    ('Шоколадный торт', 450.00),
    ('Пельмени', 380.00);

-- Тестовые данные для таблицы products с указанием единиц измерения
INSERT INTO restaurant.products (name, stock_quantity, okei_id, price_per_unit, date_of_production, date_of_expiry) VALUES
    ('Капуста', 50, 'кг', 50.00, '2023-10-01', '2023-12-01'),    -- кг
    ('Свекла', 30, 'кг', 30.00, '2023-10-02', '2023-12-02'),     -- кг
    ('Морковь', 40, 'кг', 25.00, '2023-10-03', '2023-12-03'),    -- кг
    ('Картофель', 100, 'кг', 20.00, '2023-10-04', '2023-12-04'), -- кг
    ('Куриное филе', 20, 'кг', 200.00, '2023-10-05', '2023-11-05'), -- кг
    ('Говядина', 15, 'кг', 500.00, '2023-10-06', '2023-11-06'),  -- кг
    ('Сыр', 25, 'кг', 150.00, '2023-10-07', '2023-11-07'),       -- кг
    ('Яйца', 60, 'шт', 10.00, '2023-10-08', '2023-10-28'),       -- шт
    ('Мука', 50, 'кг', 30.00, '2023-10-09', '2023-12-09'),       -- кг
    ('Сахар', 30, 'кг', 40.00, '2023-10-10', '2024-01-10'),      -- кг
    ('Молоко', 40, 'л', 60.00, '2023-10-11', '2023-10-25'),      -- л
    ('Шоколад', 20, 'г', 80.00, '2023-10-12', '2023-12-12'),      -- г
    ('Сливки', 20, 'л', 100.00, '2023-10-13', '2023-11-13'),     -- л
    ('Сметана', 20, 'л', 75.00, '2023-10-14', '2023-10-28'),     -- л
    ('Лук', 40, 'кг', 15.00, '2023-10-15', '2023-11-15');        -- кг

    -- Тестовые данные для таблицы recipes
INSERT INTO restaurant.recipes (product_id, product_quantity, dish_id) VALUES
    -- Ингредиенты для Борща
    (1, 1, 1),  -- Капуста
    (2, 1, 1),  -- Свекла
    (3, 1, 1),  -- Морковь
    (4, 2, 1),  -- Картофель
    (15, 1, 1), -- Лук

    -- Ингредиенты для Цезаря с курицей
    (5, 1, 2),  -- Куриное филе
    (7, 1, 2),  -- Сыр
    (8, 1, 2),  -- Яйца

    -- Ингредиенты для Картофельного пюре с котлетой
    (4, 3, 3),  -- Картофель
    (6, 1, 3),  -- Говядина

    -- Ингредиенты для Салата Оливье
    (4, 2, 4),  -- Картофель
    (8, 2, 4),  -- Яйца
    (15, 1, 4), -- Лук

    -- Ингредиенты для Куриного супа
    (5, 1, 5),  -- Куриное филе
    (4, 2, 5),  -- Картофель
    (15, 1, 5), -- Лук

    -- Ингредиенты для Пасты Карбонара
    (7, 1, 6),  -- Сыр
    (8, 1, 6),  -- Яйца
    (13, 1, 6), -- Сливки

    -- Ингредиенты для Стейка
    (6, 1, 7),  -- Говядина

    -- Ингредиенты для Сырников
    (9, 1, 8),  -- Мука
    (10, 1, 8), -- Сахар
    (11, 1, 8), -- Молоко

    -- Ингредиенты для Шоколадного торта
    (9, 2, 9),  -- Мука
    (10, 1, 9), -- Сахар
    (12, 1, 9), -- Шоколад

    -- Ингредиенты для Пельменей
    (6, 1, 10), -- Говядина
    (9, 1, 10), -- Мука
    (15, 1, 10); -- Лук

-- Тестовые данные для таблицы suppliers
INSERT INTO restaurant.suppliers (company_name, chief_name, email, address) VALUES
    ('ОАО "Продукты питания"', 'Иванов Иван', 'ivanov@products.ru', 'Москва, ул. 1, д. 1'),
    ('ЗАО "Фрукты и Овощи"', 'Петров Петр', 'petrov@fruits.ru', 'Санкт-Петербург, ул. 2, д. 2'),
    ('ООО "Мясная фабрика"', 'Сидоров Сидор', 'sidorov@meat.ru', 'Казань, ул. 3, д. 3'),
    ('ИП "Молочные продукты"', 'Алексеев Алексей', 'alekseev@dairy.ru', 'Екатеринбург, ул. 4, д. 4'),
    ('ООО "Кулинарные деликатесы"', 'Борисов Борис', 'borisov@delicacies.ru', 'Нижний Новгород, ул. 5, д. 5'),
    ('ЗАО "Сельские продукты"', 'Семенов Семен', 'semenov@rural.ru', 'Челябинск, ул. 6, д. 6'),
    ('ООО "Специи и травы"', 'Степанов Степан', 'stepanov@spices.ru', 'Новосибирск, ул. 7, д. 7'),
    ('ЗАО "Крупы и мука"', 'Николаев Николай', 'nikolaev@grains.ru', 'Ростов-на-Дону, ул. 8, д. 8'),
    ('ООО "Гастрономические продукты"', 'Григорьев Григорий', 'grigoryev@gastronomy.ru', 'Краснодар, ул. 9, д. 9'),
    ('ИП "Кулинарные рецепты"', 'Федоров Федор', 'fedorov@recipes.ru', 'Волгоград, ул. 10, д. 10');

-- Тестовые данные для таблицы supplies
INSERT INTO restaurant.supplies (supplier_id, transaction_id, date_from, price) VALUES
    (1, 1, '2023-11-01 10:00:00', 1500.00),
    (2, 2, '2023-11-02 11:00:00', 2500.00),
    (3, 3, '2023-11-03 12:00:00', 1800.00),
    (4, 4, '2023-11-04 13:00:00', 2200.00),
    (5, 5, '2023-11-05 14:00:00', 1200.00),
    (6, 6, '2023-11-06 15:00:00', 1600.00),
    (7, 7, '2023-11-07 16:00:00', 2000.00),
    (8, 8, '2023-11-08 17:00:00', 1900.00),
    (9, 9, '2023-11-09 18:00:00', 2100.00),
    (10, 10, '2023-11-10 19:00:00', 2400.00),
    (1, 11, '2023-11-11 10:00:00', 1300.00),
    (2, 12, '2023-11-12 11:00:00', 1700.00),
    (3, 13, '2023-11-13 12:00:00', 1400.00),
    (4, 14, '2023-11-14 13:00:00', 1600.00),
    (5, 15, '2023-11-15 14:00:00', 1500.00);

-- Тестовые данные для таблицы supply_product
INSERT INTO restaurant.supply_product (quantity, product_id, supply_id) VALUES
    (100, 1, 1),  -- 100 кг капусты для поставки 1
    (50, 2, 1),   -- 50 кг свеклы для поставки 1
    (200, 3, 2),  -- 200 кг моркови для поставки 2
    (150, 4, 2),  -- 150 кг картофеля для поставки 2
    (80, 5, 3),   -- 80 кг куриного филе для поставки 3
    (60, 6, 3),   -- 60 кг говядины для поставки 3
    (30, 7, 4),   -- 30 кг сыра для поставки 4
    (200, 8, 4),  -- 200 штук яиц для поставки 4
    (50, 9, 5),   -- 50 кг муки для поставки 5
    (40, 10, 5),  -- 40 кг сахара для поставки 5
    (30, 11, 6),  -- 30 л молока для поставки 6
    (10, 12, 7),  -- 10 кг шоколада для поставки 7
    (20, 13, 8),  -- 20 л сливок для поставки 8
    (25, 14, 9),  -- 25 л сметаны для поставки 9
    (30, 15, 10); -- 30 кг лука для поставки 10

-- Тестовые данные для таблицы orders
INSERT INTO restaurant.orders (customer_id, transaction_id, created_at, price) VALUES
    (1, 1, '2023-10-01 12:45:00', 1500.00),
    (2, 3, '2023-10-02 13:15:00', 870.00),
    (3, 5, '2023-10-03 18:00:00', 1200.00),
    (4, 7, '2023-10-04 12:30:00', 950.00),
    (5, 8, '2023-10-05 14:45:00', 2100.75),
    (6, 11, '2023-10-06 19:30:00', 1600.00),
    (7, 12, '2023-10-07 11:25:00', 750.00),
    (8, 14, '2023-10-08 14:55:00', 1850.00),
    (9, 16, '2023-10-09 20:15:00', 930.00),
    (10, 17, '2023-10-10 16:30:00', 1450.00),
    (11, 19, '2023-10-11 17:05:00', 1300.00),
    (12, 21, '2023-10-12 15:45:00', 1100.00),
    (13, 23, '2023-10-13 19:20:00', 1400.00),
    (14, 24, '2023-10-14 10:50:00', 760.00),
    (15, 26, '2023-10-15 12:25:00', 900.00),
    (16, 27, '2023-10-16 15:30:00', 1150.00),
    (17, 29, '2023-10-17 18:45:00', 1550.00),
    (18, 2, '2023-10-18 14:10:00', 2300.50),
    (19, 4, '2023-10-19 11:30:00', 5000.00),
    (20, 6, '2023-10-20 13:40:00', 3400.00),
    (21, 9, '2023-10-21 09:15:00', 4500.00),
    (22, 10, '2023-10-22 13:50:00', 1800.00),
    (23, 13, '2023-10-23 18:55:00', 2700.00),
    (24, 15, '2023-10-24 09:40:00', 5200.00),
    (25, 18, '2023-10-25 13:25:00', 2200.00),
    (26, 20, '2023-10-26 10:10:00', 4100.00),
    (27, 22, '2023-10-27 16:20:00', 2400.00),
    (28, 25, '2023-10-28 11:35:00', 5300.00),
    (29, 28, '2023-10-29 12:50:00', 2900.00),
    (30, 30, '2023-10-30 09:55:00', 4800.00);


    -- Тестовые данные для таблицы order_dish
INSERT INTO restaurant.order_dish (dish_id, order_id, dish_quantity, dish_price) VALUES
    (1, 1, 2, 300.00),  -- Борщ
    (2, 1, 1, 400.00),  -- Цезарь с курицей
    (3, 2, 1, 250.00),  -- Картофельное пюре с котлетой
    (4, 2, 3, 600.00),  -- Салат Оливье
    (5, 3, 2, 500.00),  -- Куринный суп
    (6, 3, 1, 700.00),  -- Паста Карбонара
    (7, 4, 1, 800.00),  -- Стейк
    (8, 4, 5, 1500.00), -- Сырники
    (9, 5, 1, 350.00),  -- Шоколадный торт
    (10, 5, 2, 200.00),  -- Пельмени
    (1, 6, 2, 300.00),   -- Борщ
    (3, 6, 2, 500.00),   -- Картофельное пюре с котлетой
    (2, 7, 1, 400.00),   -- Цезарь с курицей
    (4, 8, 3, 900.00),   -- Салат Оливье
    (5, 9, 1, 300.00);   -- Куринный суп

    -- Тестовые данные для таблицы payments
INSERT INTO restaurant.payments (employee_id, transaction_id, date_from, bonus) VALUES
    (1, 1, '2023-11-01 10:30:00', 100),  -- Сотрудник 1, транзакция 1
    (2, 2, '2023-11-02 11:30:00', 150),  -- Сотрудник 2, транзакция 2
    (3, 3, '2023-11-03 12:30:00', 200),  -- Сотрудник 3, транзакция 3
    (1, 4, '2023-11-04 13:30:00', 50),   -- Сотрудник 1, транзакция 4
    (2, 5, '2023-11-05 14:30:00', 75),   -- Сотрудник 2, транзакция 5
    (3, 6, '2023-11-06 15:30:00', 120),  -- Сотрудник 3, транзакция 6
    (1, 7, '2023-11-07 16:30:00', 90),   -- Сотрудник 1, транзакция 7
    (2, 8, '2023-11-08 17:30:00', 110),  -- Сотрудник 2, транзакция 8
    (3, 9, '2023-11-09 18:30:00', 130),  -- Сотрудник 3, транзакция 9
    (1, 10, '2023-11-10 19:30:00', 140); -- Сотрудник 1, транзакция 10

    -- Тестовые данные для таблицы order_employee
INSERT INTO restaurant.order_employee (employee_id, order_id) VALUES
    (1, 1),  -- Сотрудник 1, заказ 1
    (2, 1),  -- Сотрудник 2, заказ 1
    (1, 2),  -- Сотрудник 1, заказ 2
    (3, 3),  -- Сотрудник 3, заказ 3
    (2, 3),  -- Сотрудник 2, заказ 3
    (1, 4),  -- Сотрудник 1, заказ 4
    (3, 4),  -- Сотрудник 3, заказ 4
    (2, 5),  -- Сотрудник 2, заказ 5
    (1, 6),  -- Сотрудник 1, заказ 6
    (3, 7);  -- Сотрудник 3, заказ 7
    
