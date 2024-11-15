// App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Sidebar from './components/Sidebar'; // Импортируем Sidebar
import Dishes from './components/Dishes';
import Employees from './components/Employees';
import Supplies from './components/Supplies';
import Customers from './components/Customers';
import Content from './components/Content';
import Report from './components/Report';

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
                        <Route path="/supplies" element={<Supplies />} />
                        <Route path="/customers" element={<Customers />} />
                        <Route path="/report" element={<Report />} />
                    </Routes>
                </Content>
            </div>
        </Router>
    );
};

export default App;
