// Sidebar.js
import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = () => {
    const location = useLocation(); // Получаем текущий путь

    return (
        <div className="sidebar">
            <h1>Университетское кафе</h1>
            <ul>
                <li>
                    <Link to="/dishes" className={location.pathname === "/dishes" ? "active" : ""}>Меню</Link>
                </li>
                <li>
                    <Link to="/employees" className={location.pathname === "/employees" ? "active" : ""}>Сотрудники</Link>
                </li>
                <li>
                    <Link to="/products" className={location.pathname === "/products" ? "active" : ""}>Поставки</Link>
                </li>
                <li>
                    <Link to="/customers" className={location.pathname === "/customers" ? "active" : ""}>Клиенты</Link>
                </li>
            </ul>
        </div>
    );
};

export default Sidebar;
