// src/components/ErrorModal.js
import React, { useEffect, useState } from 'react';
import './ErrorModal.css';

const ErrorModal = ({ message, onClose }) => {
    const [show, setShow] = useState(false);

    useEffect(() => {
        if (message) {
            setShow(true);

            // Закрытие окна через 5 секунд (5000 мс)
            const timer = setTimeout(() => {
                setShow(false);
                onClose(); // вызываем onClose, чтобы сбросить ошибку в родительском компоненте
            }, 5000);

            return () => clearTimeout(timer); // Очистка таймера при размонтировании компонента
        }
    }, [message, onClose]);

    return (
        <div className={`error-modal ${show ? 'show' : ''}`}>
            <span className="close-btn" onClick={onClose}>&times;</span>
            <h2>Ошибка</h2>
            <p>{message}</p>
        </div>
    );
};

export default ErrorModal;
