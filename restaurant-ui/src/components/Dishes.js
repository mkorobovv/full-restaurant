// src/components/Dishes.js
import React, { useState, useEffect } from "react";
import { getDishesWithIngredients, getDishesByIngredients, getMostPopularDishes, getExpiringProducts } from "../api";

const Dishes = () => {
    const [dishes, setDishes] = useState([]);
    const [expiringProducts, setExpiringProducts] = useState([]);
    const [error, setError] = useState(null);
    const [activeTab, setActiveTab] = useState("dishes");
    const [ingredient, setIngredient] = useState("");

    const fetchDishes = async () => {
        try {
            const data = await getDishesWithIngredients();
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError(err.message);
        }
    };

    const fetchDishesByIngredient = async () => {
        try {
            const data = await getDishesByIngredients(ingredient);
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError(err.message);
        }
    };

    const fetchMostPopularDishes = async () => {
        try {
            const data = await getMostPopularDishes();
            setDishes(data);
            setExpiringProducts([]);
        } catch (err) {
            setError(err.message);
        }
    };

    const fetchExpiringProducts = async () => {
        try {
            const data = await getExpiringProducts();
            setExpiringProducts(data);
            setDishes([]);
        } catch (err) {
            setError(err.message);
        }
    };

    useEffect(() => {
        if (activeTab === "dishes") fetchDishes();
    }, [activeTab]);

    const handleIngredientSearch = (e) => {
        e.preventDefault();
        fetchDishesByIngredient();
    };

    return (
        <div className="content">
            <h1>Меню</h1>

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
                    />
                    <button type="submit" className="fetch-button">Поиск</button>
                </form>
            )}

            {/* Ошибки, если есть */}
            {error && <div className="error">Ошибка: {error}</div>}

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
