const API_URL = "http://localhost:8080";

// Общая функция для обработки запросов
const fetchData = async (url, options = {}) => {
    try {
        const response = await fetch(`${API_URL}${url}`, options);
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.message || "Ошибка запроса");
        }
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }
};

// Примеры функций API
export const getDishesWithIngredients = () =>
    fetchData("/api/dishes/get", {method: "GET"});

export const getUnorderedDishes = () =>
    fetchData("/api/dishes/get/unordered", {method: "GET"});

export const getDishesByIngredients = (ingredient) =>
    fetchData(`/api/dishes/get/by-ingredients?ingredient=${ingredient}`,{method: "GET"});

export const getEmployeeOrderCount = (dateFrom, dateTo) =>
    fetchData(`/api/employees/orders-count?date_from=${dateFrom}&date_to=${dateTo}`, {method: "GET"});

export const getExpiringProducts = () =>
    fetchData("/api/products/expiring-soon", {method: "GET"});

export const downloadReport = (dateFrom, dateTo) =>
    fetchData(`/api/report/download?date_from=${dateFrom}&date_to=${dateTo}`, { method: "GET" });

export const getCustomerOrderHistory = (customerId, dateFrom, dateTo) =>
    fetchData(`/api/customers/${customerId}/orders?date_from=${dateFrom}&date_to=${dateTo}`,{method: "GET"});

export const getMostPopularDishes = () =>
    fetchData("/api/dishes/popular",{method: "GET"});

export const getSuppliersByProduct = (productName) =>
    fetchData(`/api/products/${productName}/suppliers`, {method: "GET"});

export const getReportDownload = (dateFrom, dateTo) =>
        fetchData(`/api/report/download?date_from=${dateFrom}&date_to=${dateTo}`, {method: "GET"});


