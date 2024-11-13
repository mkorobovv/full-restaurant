// src/components/Employee.js
import React, { useState } from "react";
import { getEmployeeOrderCount } from "../api"; // Импортируем API метод
import ErrorModal from "./ErrorModal"; // Ошибка модалка

const Employee = () => {
    const [dateFrom, setDateFrom] = useState("");
    const [dateTo, setDateTo] = useState("");
    const [orderCount, setOrderCount] = useState(null);
    const [error, setError] = useState(null);

    const fetchOrderCount = async () => {
        try {
            const data = await getEmployeeOrderCount(dateFrom, dateTo);
            setOrderCount(data); // Сохраняем количество заказов
        } catch (err) {
            setError(err.message); // Устанавливаем ошибку
        }
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        fetchOrderCount(); // Запрос на количество заказов
    };

    const closeErrorModal = () => {
        setError(null); // Закрытие модального окна с ошибкой
    };

    return (
        <div className="content">
            <h1>Количество заказов сотрудника</h1>

            {/* Форма для ввода данных */}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>
                        Дата с:
                        <input
                            type="date"
                            value={dateFrom}
                            onChange={(e) => setDateFrom(e.target.value)}
                            required
                        />
                    </label>
                </div>
                <div>
                    <label>
                        Дата по:
                        <input
                            type="date"
                            value={dateTo}
                            onChange={(e) => setDateTo(e.target.value)}
                            required
                        />
                    </label>
                </div>
                <button type="submit">Получить количество заказов</button>
            </form>

            {/* Ошибка, если есть */}
            {error && <ErrorModal message={error} onClose={closeErrorModal} />}

            {/* Результаты */}
            <div>
                {orderCount !== null ? (
                    <div>
                        <h2>Количество заказов: {orderCount}</h2>
                    </div>
                ) : (
                    <p>Введите диапазон дат для получения данных.</p>
                )}
            </div>
        </div>
    );
};

export default Employee;
