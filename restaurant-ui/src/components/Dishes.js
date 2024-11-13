import React, { useState, useEffect } from "react";
import { getDishesWithIngredients, getDishesByIngredients, getMostPopularDishes, getExpiringProducts } from "../api";
import ErrorModal from "./ErrorModal";

const Dishes = () => {
    const [dishes, setDishes] = useState([]);
    const [expiringProducts, setExpiringProducts] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);
    const [activeTab, setActiveTab] = useState("dishes");
    const [ingredient, setIngredient] = useState("");
    const [activeDish, setActiveDish] = useState(null); // Для управления видимостью ингредиентов

    // Функция для переключения видимости ингредиентов
    const toggleIngredients = (dishId) => {
        setActiveDish(activeDish === dishId ? null : dishId);
    };

    const fetchDataWithRetry = async (fetchFunction, setData, retries = 3) => {
        setError(null);
        setLoading(true);
        while (retries > 0) {
            try {
                const data = await fetchFunction();
                console.log("Ответ от сервера:", data); // Логируем ответ
                setData(data);
                setLoading(false); // Завершаем загрузку
                return;
            } catch (err) {
                retries -= 1;
                if (retries === 0) {
                    if (err.message === "Failed to fetch") {
                        setError("Не удалось подключиться к серверу. Проверьте ваше интернет-соединение или доступность сервера.");
                    } else {
                        setError("Ошибка загрузки данных: " + err.message);
                    }
                    setLoading(false); // Завершаем загрузку при ошибке
                }
            }
        }
    };

    const handleIngredientSearch = async (e) => {
        e.preventDefault();
        if (!ingredient) {
            setError("Пожалуйста, введите название продукта.");
            return;
        }
        setLoading(true);
        setError(null);
        setDishes([]); // Очищаем блюда перед новым поиском

        try {
            const data = await getDishesByIngredients(ingredient);
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError("Ошибка поиска по ингредиенту: " + err.message);
        } finally {
            setLoading(false);
            setIngredient("");
        }
    };

    useEffect(() => {
        setDishes([]); // Очищаем блюда при смене вкладки
        setExpiringProducts([]); // Очищаем истекающие продукты при смене вкладки
        setError(null); // Сбрасываем ошибки

        if (activeTab === "dishes") {
            fetchDataWithRetry(getDishesWithIngredients, setDishes, 3);
        } else if (activeTab === "popular") {
            fetchDataWithRetry(getMostPopularDishes, setDishes, 3);
        } else if (activeTab === "expiring") {
            fetchDataWithRetry(getExpiringProducts, setExpiringProducts, 3);
        }
    }, [activeTab]);

    const closeErrorModal = () => setError(null);

    return (
        <div className="content">
            <h1>Блюда</h1>
            <div className="navbar">
                <button onClick={() => setActiveTab("dishes")} className={activeTab === "dishes" ? "active-tab" : ""}>Все блюда</button>
                <button onClick={() => setActiveTab("popular")} className={activeTab === "popular" ? "active-tab" : ""}>Популярные блюда</button>
                <button onClick={() => setActiveTab("byIngredient")} className={activeTab === "byIngredient" ? "active-tab" : ""}>Поиск по ингредиенту</button>
                <button onClick={() => setActiveTab("expiring")} className={activeTab === "expiring" ? "active-tab" : ""}>Истекающие продукты</button>
            </div>

            {activeTab === "byIngredient" && (
                <form onSubmit={handleIngredientSearch} className="ingredient-search">
                    <input
                        type="text"
                        value={ingredient}
                        onChange={(e) => setIngredient(e.target.value)}
                        placeholder="Введите ингредиент"
                        required
                    />
                    <button type="submit" className="fetch-button">Поиск</button>
                </form>
            )}

            {error && <ErrorModal message={error} onClose={closeErrorModal} />}
            {loading ? <p>Загрузка данных...</p> : (
                <div className="dishes-list">
                    {/* Для вкладки Истекающие продукты */}
                    {activeTab === "expiring" && expiringProducts.length > 0 ? (
                        expiringProducts.map((product) => (
                            <div
                                key={product.product_id}
                                className={`dish ${product.status === "expired" ? "expired" : "expiring-soon"}`}
                            >
                                <h2>{product.name}</h2>
                                <p>Срок годности: {new Date(product.date_of_expiry).toLocaleDateString("ru-RU")}</p>
                                <p>Статус: {product.status}</p>
                            </div>
                        ))
                    ) : activeTab === "popular" && dishes.length > 0 ? (
                        // Для вкладки Популярные блюда
                        dishes.map((dish, index) => (
                            <div key={dish.dish_id} className="dish">
                                <div className="dish-info">
                                    <h2>№{index + 1}. {dish.dish_name}</h2> {/* Индекс + 1 для отображения порядкового номера */}
                                    <p><strong>Цена:</strong> {dish.price}₽</p>
                                    <p><strong>Количество заказов:</strong> {dish.ordered_count}</p>
                                </div>
                            </div>
                        ))
                    ) : activeTab === "byIngredient" || activeTab === "dishes" && dishes.length > 0 ? (
                        dishes.map((dish) => (
                            <div key={dish.dish_id} className="dish">
                                {/* Название блюда */}
                                <h2>{dish.name}</h2> {/* Убедитесь, что "dish_name" существует в данных */}
                                <p>Цена: {dish.price}₽</p>

                                {/* Только в случае, если это не вкладка "Поиск по ингредиенту" показываем ингредиенты */}
                                {activeTab !== "byIngredient" && (
                                    <>
                                        {/* Стрелочка и подпись для переключения видимости ингредиентов */}
                                        <div
                                            className={`toggle-ingredients ${activeDish === dish.dish_id ? "open" : ""}`}
                                            onClick={() => toggleIngredients(dish.dish_id)}
                                        >
                                            <span className="arrow"></span>
                                            <span
                                                className="ingredients-label">Ингредиенты</span> {/* Подпись "Ингредиенты" */}
                                        </div>

                                        {/* Скрытые ингредиенты, отображаем их только если активен данный dish */}
                                        <div
                                            className={`ingredients-container ${activeDish === dish.dish_id ? "open" : ""}`}>
                                            <ul className="ingredients-list">
                                                {dish.ingredients && dish.ingredients.map((ingredient) => (
                                                    <li key={ingredient.id}>{ingredient.name}</li>
                                                ))}
                                            </ul>
                                        </div>
                                    </>
                                )}
                            </div>
                        ))
                    ) : (
                        <p>Нет данных для отображения.</p>
                    )}
                </div>
            )}
        </div>
    );
};

export default Dishes;
