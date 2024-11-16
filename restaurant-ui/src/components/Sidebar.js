// Sidebar.js
import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = () => {
    const location = useLocation(); // Получаем текущий путь

    return (
        <div className="sidebar">
            <div className="logo-container">
                <img src="https://mai.ru/press/brand/download/Default/RU/Default.svg" alt="Логотип" className="logo"/>
            </div>
            <h1>Университетское кафе</h1>
            <ul>
                <li>
                    <Link to="/dishes" className={location.pathname === "/dishes" ? "active" : ""}>Блюда</Link>
                </li>
                <li>
                    <Link to="/employees"
                          className={location.pathname === "/employees" ? "active" : ""}>Сотрудники</Link>
                </li>
                <li>
                    <Link to="/supplies" className={location.pathname === "/supplies" ? "active" : ""}>Поставки</Link>
                </li>
                <li>
                    <Link to="/customers" className={location.pathname === "/customers" ? "active" : ""}>Клиенты</Link>
                </li>
                <li>
                    <Link to="/report" className={location.pathname === "/report" ? "active" : ""}>Отчет</Link>
                </li>
            </ul>
        </div>
    );
};

export default Sidebar;
