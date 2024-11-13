// src/components/Customer.js
import React, { useState } from "react";
import { getCustomerOrderHistory } from "../api"; // Импортируем API метод
import ErrorModal from "./ErrorModal"; // Ошибка модалка

const Customer = () => {
    const [customerId, setCustomerId] = useState("");
    const [dateFrom, setDateFrom] = useState("");
    const [dateTo, setDateTo] = useState("");
    const [orderHistory, setOrderHistory] = useState([]);
    const [error, setError] = useState(null);

    const fetchOrderHistory = async () => {
        try {
            const data = await getCustomerOrderHistory(customerId, dateFrom, dateTo);
            setOrderHistory(data); // Сохраняем историю заказов
        } catch (err) {
            setError(err.message); // Устанавливаем ошибку
        }
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        fetchOrderHistory(); // Запрос на историю заказов
    };

    const closeErrorModal = () => {
        setError(null); // Закрытие модального окна с ошибкой
    };

    return (
        <div className="content">
            <h1>История заказов клиента</h1>

            {/* Форма для ввода данных */}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>
                        ID клиента:
                        <input
                            type="text"
                            value={customerId}
                            onChange={(e) => setCustomerId(e.target.value)}
                            required
                        />
                    </label>
                </div>
                <div>
                    <label>
                        Дата от:
                        <input
                            type="date"
                            value={dateFrom}
                            onChange={(e) => setDateFrom(e.target.value)}
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
                        />
                    </label>
                </div>
                <button type="submit">Получить историю</button>
            </form>

            {/* Ошибка, если есть */}
            {error && <ErrorModal message={error} onClose={closeErrorModal} />}

            {/* Результаты */}
            <div>
                {orderHistory.length > 0 ? (
                    <table>
                        <thead>
                        <tr>
                            <th>Дата</th>
                            <th>Товары</th>
                            <th>Сумма</th>
                        </tr>
                        </thead>
                        <tbody>
                        {orderHistory.map((order) => (
                            <tr key={order.order_id}>
                                <td>{order.date}</td>
                                <td>{order.items.join(", ")}</td>
                                <td>{order.total_amount}₽</td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                ) : (
                    <p>Нет заказов для выбранного периода.</p>
                )}
            </div>
        </div>
    );
};

export default Customer;
