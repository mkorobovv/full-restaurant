import React, { useState } from "react";
import { getCustomerOrderHistory } from "../api"; // Importing API method
import ErrorModal from "./ErrorModal"; // Error modal component

const Customer = () => {
    const [customerId, setCustomerId] = useState("");
    const [dateFrom, setDateFrom] = useState("");
    const [dateTo, setDateTo] = useState("");
    const [orderHistory, setOrderHistory] = useState([]);
    const [error, setError] = useState(null);

    // Fetch order history from the API
    const fetchOrderHistory = async () => {
        try {
            const data = await getCustomerOrderHistory(customerId, dateFrom, dateTo);
            console.log("Fetched data: ", data);  // Check what is being returned from the API
            setOrderHistory(data);  // Save the fetched data
        } catch (err) {
            setError("Ошибка загрузки данных: " + err.message);  // Show error message if fetching fails
        }
    };

    // Handle form submission
    const handleSubmit = (e) => {
        e.preventDefault();
        fetchOrderHistory(); // Fetch the order history for the given customer and dates
    };

    // Close the error modal
    const closeErrorModal = () => {
        setError(null);
    };

    // Check if the client exists
    const isCustomerFound = orderHistory && orderHistory.first_name && orderHistory.first_name.trim() !== "";

    return (
        <div className="content">
            <h1>История заказов клиента</h1>

            {/* Form to input customer ID and date range */}
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

            {/* Error Modal */}
            {error && <ErrorModal message={error} onClose={closeErrorModal}/>}

            {/* Displaying Order History */}
            <div>
                {/* Check if customer exists */}
                {isCustomerFound ? (
                    <div>
                        <h2>Информация о клиенте</h2>
                        <p><strong>Имя:</strong> {orderHistory.first_name} {orderHistory.last_name}</p>
                        <p><strong>Телефон:</strong> {orderHistory.phone_number}</p>
                        <p><strong>Email:</strong> {orderHistory.email}</p>
                        <p><strong>Скидка:</strong> {orderHistory.discount}%</p>

                        <h2>История заказов</h2>
                        {/* Check if orders exist and if they have any data */}
                        {orderHistory.orders && orderHistory.orders.length > 0 ? (
                            <div className="orders-container">
                                {orderHistory.orders.map((order) => (
                                    <div className="order-block" key={order.order_id}>
                                        <h3>Заказ №{order.order_id}</h3>
                                        <p><strong>Дата заказа:</strong> {new Date(order.created_at).toLocaleDateString()}</p>

                                        <h4>Товары:</h4>
                                        <div className="order-dishes">
                                            {order.dishes.map((dish, index) => (
                                                <div className="dish-item" key={index}>
                                                    <strong>{dish.name}</strong> ({dish.dish_quantity} шт.)
                                                    - {dish.dish_price}₽
                                                </div>
                                            ))}
                                        </div>

                                        <p><strong>Сумма:</strong> {order.price}₽</p>
                                    </div>
                                ))}
                            </div>
                        ) : (
                            <p>У клиента нет заказов для выбранного периода.</p>
                        )}
                    </div>
                ) : (
                    <p>Клиент не найден. Пожалуйста, проверьте введенные данные.</p>
                )}
            </div>
        </div>
    );
};

export default Customer;
