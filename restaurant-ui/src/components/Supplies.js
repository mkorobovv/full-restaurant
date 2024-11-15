import React, { useState } from "react";
import { getSuppliersByProduct } from "../api";
import ErrorModal from "./ErrorModal";

const Supplies = () => {
    const [productName, setProductName] = useState("");
    const [suppliers, setSuppliers] = useState([]);
    const [error, setError] = useState(null);

    const handleSearch = async (e) => {
        e.preventDefault();
        setError(null);
        if (!productName) {
            setError("Пожалуйста, введите название продукта.");
            return;
        }

        try {
            const data = await getSuppliersByProduct(productName);
            console.log(data);
            setSuppliers(data);
        } catch (err) {
            setError("Ошибка загрузки поставщиков: " + err.message);
        }
    };

    const closeErrorModal = () => {
        setError(null);
    };

    return (
        <div className="content">
            <h1>Поставщики продукта</h1>

            {/* Форма для ввода названия продукта */}
            <form onSubmit={handleSearch} className="ingredient-search">
                <input
                    type="text"
                    value={productName}
                    onChange={(e) => setProductName(e.target.value)}
                    placeholder="Введите название продукта"
                />
                <button type="submit" className="fetch-button">Поиск</button>
            </form>

            {/* Модальное окно с ошибкой */}
            {error && <ErrorModal message={error} onClose={closeErrorModal} />}

            {/* Список поставщиков */}
            {suppliers.length > 0 ? (
                <div className="suppliers-list">
                    {suppliers.map((supplier) => (
                        <div key={supplier.supplier_id} className="supplier-card">
                            <h2>{supplier.company_name}</h2>
                            <p><strong>Руководитель:</strong> {supplier.chief_name}</p>
                            <p><strong>Адрес:</strong> {supplier.address}</p>
                            <p><strong>Email:</strong> {supplier.email}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>Нет данных для отображения.</p>
            )}
        </div>
    );
};

export default Supplies;
