import React, { useState } from "react";
import { getReportDownload } from "../api"; // Импорт API метода
import ErrorModal from "./ErrorModal"; // Модальное окно для ошибок

const Report = () => {
    const [dateFrom, setDateFrom] = useState("");
    const [dateTo, setDateTo] = useState("");
    const [report, setReport] = useState(null); // Состояние для отчета
    const [error, setError] = useState(null);

    // Функция для получения отчета
    const fetchReport = async () => {
        try {
            const data = await getReportDownload(dateFrom, dateTo);
            console.log("Fetched report data: ", data);  // Проверка данных, полученных с API
            setReport(data);  // Сохранение данных отчета
        } catch (err) {
            setError("Ошибка загрузки данных: " + err.message);  // Показать сообщение об ошибке
        }
    };

    // Обработка отправки формы
    const handleSubmit = (e) => {
        e.preventDefault();
        fetchReport(); // Запросить отчет для выбранного диапазона дат
    };

    // Закрытие модального окна ошибки
    const closeErrorModal = () => {
        setError(null);
    };

    return (
        <div className="content">
            <h1>Отчет за период</h1>

            {/* Форма для выбора диапазона дат */}
            <form onSubmit={handleSubmit}>
                <div className="date-picker">
                    <label>
                        Дата от:
                        <input
                            type="date"
                            value={dateFrom}
                            onChange={(e) => setDateFrom(e.target.value)}
                        />
                    </label>
                </div>
                <div className="date-picker">
                    <label>
                        Дата по:
                        <input
                            type="date"
                            value={dateTo}
                            onChange={(e) => setDateTo(e.target.value)}
                        />
                    </label>
                </div>
                <button className="fetch-btn" type="submit">Получить отчет</button>
            </form>

            {/* Модальное окно с ошибкой */}
            {error && <ErrorModal message={error} onClose={closeErrorModal}/>}

            {/* Отображение отчета */}
            <div className="report-container">
                {report ? (
                    <div className="report-card">
                        <h2>Отчет с {new Date(dateFrom).toLocaleDateString()} по {new Date(dateTo).toLocaleDateString()}</h2>

                        {/* Формальный отчет в виде таблицы */}
                        <table className="report-table">
                            <thead>
                            <tr>
                                <th>Показатель</th>
                                <th>Значение</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td>Потери в доходах</td>
                                <td>{report.lost_revenue.toFixed(2)}₽</td>
                            </tr>
                            <tr>
                                <td>Чистая прибыль</td>
                                <td>{report.net_profit.toFixed(2)}₽</td>
                            </tr>
                            <tr>
                                <td>Средний чек поставки</td>
                                <td>{report.average_supply_check.toFixed(2)}₽</td>
                            </tr>
                            <tr>
                                <td>Средний чек заказа</td>
                                <td>{report.average_order_check.toFixed(2)}₽</td>
                            </tr>
                            <tr>
                                <td>Сумма затрат на поставку</td>
                                <td>{report.amount_supply_costs.toFixed(2)}₽</td>
                            </tr>
                            </tbody>
                        </table>

                    </div>
                ) : (
                    <p>Загрузите отчет для отображения данных.</p>
                )}
            </div>
        </div>
    );
};

export default Report;
