import React, { useState } from "react";
import { getEmployeeOrderCount } from "../api"; // Импортируем API метод
import ErrorModal from "./ErrorModal"; // Ошибка модалка

const Employee = () => {
    const [dateFrom, setDateFrom] = useState("");
    const [dateTo, setDateTo] = useState("");
    const [employees, setEmployees] = useState([]);
    const [error, setError] = useState(null);

    const fetchOrderCount = async () => {
        try {
            const data = await getEmployeeOrderCount(dateFrom, dateTo);
            console.log(data);
            setEmployees(data); // Assuming the data is an array of employee objects
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
            <form onSubmit={handleSubmit} className="date-form">
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
            <div className="employee-list">
                {employees.length > 0 ? (
                    <div className="employee-grid">
                        {employees.map((employee) => (
                            <div className="employee-card" key={employee.employee_id}>
                                <h3>{employee.first_name} {employee.last_name}</h3>
                                <p><strong>Должность:</strong> {employee.position_name}</p>
                                <p><strong>Телефон:</strong> {employee.phone_number}</p>
                                <p><strong>Зарплата:</strong> {employee.salary} ₽</p>
                                <p><strong>Количество заказов:</strong> {employee.orders_count}</p>
                            </div>
                        ))}
                    </div>
                ) : (
                    <p>Введите диапазон дат для получения данных.</p>
                )}
            </div>
        </div>
    );
};

export default Employee;
