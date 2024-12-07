document.addEventListener("DOMContentLoaded", function() {
    const registerForm = document.getElementById("registerForm");
    const loginForm = document.getElementById("loginForm");

    // Обработчик для формы регистрации
    registerForm.addEventListener("submit", function(event) {
        event.preventDefault();  // Отменяем стандартную отправку формы
        console.log("Register form submitted");
        
        const login = document.getElementById("registerLogin").value;
        const password = document.getElementById("registerPassword").value;
        const role_id = document.getElementById("registerRole").value;
    
        const data = {
            login: login,
            password: password,
            role_id: parseInt(role_id)  // Преобразуем роль в число
        };
        console.log("Sending registration data:", data);
    
        fetch("http://localhost:8080/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Registration response:", data);
            if (data.token) {
                localStorage.setItem("token", data.token);
                alert("Registration successful!");
                // Перенаправляем на главную страницу после успешной регистрации
                window.location.href = "main.html"; // Замените на путь к вашей главной странице
            } else {
                alert("Error registering: " + data.error);
            }
        })
        .catch(error => {
            console.error("Error during registration:", error);
            alert("An error occurred during registration");
        });
    });
    

    // Обработчик для формы логина
    loginForm.addEventListener("submit", function(event) {
        event.preventDefault();  // Отменяем стандартную отправку формы
        console.log("Login form submitted");

        const login = document.getElementById("loginLogin").value;
        const password = document.getElementById("loginPassword").value;

        const data = {
            login: login,
            password: password
        };

        console.log("Sending login data:", data);

        fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            console.log("Login response:", data);
            if (data.token && data.client_id) {
                // Сохраняем token и client_id в localStorage
                localStorage.setItem("token", data.token);
                localStorage.setItem("client_id", data.client_id);  // Сохраняем client_id
                alert("Login successful!");
                window.location.href = "main.html";  // Перенаправление на главную страницу после логина
            } else {
                alert("Error logging in: " + data.error);
            }
        })
        
        .catch(error => {
            console.error("Error during login:", error);
            alert("An error occurred during login");
        });
    });
});

document.getElementById('loadServicesBtn').addEventListener('click', loadServices);

function loadServices() {
    fetch('http://localhost:8080/services', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token'), // Токен, если он хранится в localStorage
        }
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Ошибка при загрузке услуг');
    })
    .then(services => {
        renderServicesTable(services);
    })
    .catch(error => {
        console.error(error);
        alert('Не удалось загрузить услуги');
    });
}

function renderServicesTable(services) {
    const tableBody = document.querySelector('#servicesTable tbody');
    tableBody.innerHTML = ''; // Очищаем таблицу перед добавлением новых данных
    
    services.forEach(service => {
        const row = document.createElement('tr');
        
        row.innerHTML = `
            <td>${service.id}</td>
            <td>${service.service_type}</td>
            <td>${service.duration}</td>
            <td><button onclick="bookService(${service.id})">Записаться</button></td>
        `;
        
        tableBody.appendChild(row);
    });
    
    // Показываем таблицу с услугами
    document.getElementById('servicesTableContainer').style.display = 'block';
}

function bookService(serviceId) {
    // Предположим, что у нас есть токен и id клиента в localStorage
    const clientId = localStorage.getItem('clientId'); 
    if (!clientId) {
        alert('Вы должны быть авторизованы как клиент');
        return;
    }

    fetch('http://localhost:8080/service_requests', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token'),
        },
        body: JSON.stringify({
            client_id: clientId,
            service_id: serviceId,
            status: 0,  // Статус "в процессе"
            request_date: new Date().toISOString().split('T')[0],  // Текущая дата
        })
    })
    .then(response => {
        if (response.ok) {
            alert('Вы успешно записались на услугу!');
        } else {
            throw new Error('Не удалось записаться на услугу');
        }
    })
    .catch(error => {
        console.error(error);
        alert('Ошибка при записи на услугу');
    });
}
