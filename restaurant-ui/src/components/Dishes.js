// src/components/Dishes.js
import React, { useState, useEffect } from "react";
import { getDishesWithIngredients, getDishesByIngredients, getMostPopularDishes, getExpiringProducts } from "../api";
import ErrorModal from "./ErrorModal";

const Dishes = () => {
    const [dishes, setDishes] = useState([]);
    const [expiringProducts, setExpiringProducts] = useState([]);
    const [error, setError] = useState(null);
    const [activeTab, setActiveTab] = useState("dishes");
    const [ingredient, setIngredient] = useState("");

    const fetchDishes = async () => {
        setError(null); // Сброс ошибки перед запросом
        try {
            const data = await getDishesWithIngredients();
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError("Ошибка загрузки блюд: " + err.message);
        }
    };

    const fetchDishesByIngredient = async () => {
        setError(null);
        if (!ingredient) {
            setError("Пожалуйста, введите название продукта.");
            return;
        }

        try {
            const data = await getDishesByIngredients(ingredient);
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError("Ошибка поиска по ингредиенту: " + err.message);
        }
    };

    const fetchMostPopularDishes = async () => {
        setError(null); // Сброс ошибки перед запросом
        try {
            const data = await getMostPopularDishes();
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError("Ошибка загрузки популярных блюд: " + err.message);
        }
    };

    const fetchExpiringProducts = async () => {
        setError(null); // Сброс ошибки перед запросом
        try {
            const data = await getExpiringProducts();
            setExpiringProducts(data);
            setDishes([]);
        } catch (err) {
            setError("Ошибка загрузки продуктов с истекающим сроком годности: " + err.message);
        }
    };

    // При смене вкладки сбрасываем ошибку и запускаем соответствующую функцию
    useEffect(() => {
        setError(null); // Сброс ошибки при переключении вкладок
        if (activeTab === "dishes") fetchDishes();
        if (activeTab === "popular") fetchMostPopularDishes();
        if (activeTab === "expiring") fetchExpiringProducts();
    }, [activeTab]);

    const handleIngredientSearch = (e) => {
        e.preventDefault();
        fetchDishesByIngredient();
    };

    const closeErrorModal = () => {
        setError(null);
    };

    return (
        <div className="content">
            <h1>Блюда</h1>

            {/* Навигационная панель */}
            <div className="navbar">
                <button onClick={() => setActiveTab("dishes")} className={activeTab === "dishes" ? "active-tab" : ""}>Все блюда</button>
                <button onClick={() => setActiveTab("popular")} className={activeTab === "popular" ? "active-tab" : ""}>Популярные блюда</button>
                <button onClick={() => setActiveTab("byIngredient")} className={activeTab === "byIngredient" ? "active-tab" : ""}>Поиск по ингредиенту</button>
                <button onClick={() => setActiveTab("expiring")} className={activeTab === "expiring" ? "active-tab" : ""}>Истекающие продукты</button>
            </div>

            {/* Поле для ввода ингредиента, отображается только при выборе вкладки "Поиск по ингредиенту" */}
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

            {/* Модальное окно с ошибкой */}
            {error && <ErrorModal message={error} onClose={closeErrorModal} />}

            {/* Контент в зависимости от выбранной вкладки */}
            <div className="dishes-list">
                {activeTab === "expiring" && expiringProducts.length > 0 ? (
                    expiringProducts.map((product) => (
                        <div key={product.product_id} className="dish">
                            <h2>{product.name}</h2>
                            <p>Срок годности: {product.date_of_expiry}</p>
                            <p>Статус: {product.status}</p>
                        </div>
                    ))
                ) : activeTab !== "expiring" && dishes.length > 0 ? (
                    dishes.map((dish) => (
                        <div key={dish.dish_id} className="dish">
                            <h2>{dish.name}</h2>
                            <p>Цена: {dish.price}₽</p>
                            {dish.ingredients && <p>Ингредиенты: {dish.ingredients.join(", ")}</p>}
                        </div>
                    ))
                ) : (
                    <p>Нет данных для отображения.</p>
                )}
            </div>
        </div>
    );
};

export default Dishes;
