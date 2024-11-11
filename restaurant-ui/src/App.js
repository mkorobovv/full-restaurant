// App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Sidebar from './components/Sidebar'; // Импортируем Sidebar
import Dishes from './components/Dishes';
import Employees from './components/Employees';
import Products from './components/Products';
import Customers from './components/Customers';
import Content from './components/Content'; // Импортируем Content для отображения контента

const App = () => {
    return (
        <Router>
            <div className="app">
                <Sidebar /> {/* Боковая панель */}
                <Content>
                    <Routes>
                        <Route path="/" element={<h1>Система управления университетским кафе</h1>} />
                        <Route path="/dishes" element={<Dishes />} />
                        <Route path="/employees" element={<Employees />} />
                        <Route path="/products" element={<Products />} />
                        <Route path="/customers" element={<Customers />} />
                    </Routes>
                </Content>
            </div>
        </Router>
    );
};

export default App;
