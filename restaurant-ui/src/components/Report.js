import React, { useState } from "react";
import { getReportDownload } from "../api"; // Импортируем API метод для получения отчета
import ErrorModal from "./ErrorModal"; // Модалка ошибки

const Report = () => {
    const [dateFrom, setDateFrom] = useState(""); // Начальная дата
    const [dateTo, setDateTo] = useState(""); // Конечная дата
    const [error, setError] = useState(null); // Ошибка
    const [isLoading, setIsLoading] = useState(false); // Состояние загрузки

    const fetchReport = async () => {
        if (!dateFrom || !dateTo) {
            setError("Пожалуйста, укажите обе даты."); // Ошибка если даты не указаны
            return;
        }

        setIsLoading(true); // Начинаем загрузку
        setError(null); // Сбрасываем ошибку

        try {
            const data = await getReportDownload(dateFrom, dateTo); // Получаем отчет от API

            // Проверяем, получили ли мы данные в правильном формате (XLSX)
            if (!data) {
                throw new Error("Не удалось получить данные для отчета.");
            }

            // Инициализируем скачивание
            const blob = new Blob([data], { type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" });
            const link = document.createElement("a");
            link.href = URL.createObjectURL(blob);
            link.download = `report_${dateFrom}_${dateTo}.xlsx`; // Название файла
            link.click();
        } catch (err) {
            setError(err.message); // Устанавливаем ошибку
        } finally {
            setIsLoading(false); // Заканчиваем загрузку
        }
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        fetchReport(); // Вызов API для скачивания отчета
    };

    const closeErrorModal = () => {
        setError(null); // Закрываем модальное окно с ошибкой
    };

    return (
        <div className="content">
            <h1>Скачать отчет за период</h1>

            {/* Форма для ввода дат */}
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
                <button type="submit" disabled={isLoading}>
                    {isLoading ? "Загружаю..." : "Скачать отчет"}
                </button>
            </form>

            {/* Ошибка, если есть */}
            {error && <ErrorModal message={error} onClose={closeErrorModal} />}

            {/* Информация по отчёту */}
            <div className="report-info">
                {!error && !isLoading && (
                    <p>Введите диапазон дат и нажмите "Скачать отчет", чтобы получить файл.</p>
                )}
            </div>
        </div>
    );
};

export default Report;
