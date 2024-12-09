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
                //window.location.href = "main.html"; // Замените на путь к вашей главной странице
            } else {
                alert("Error registering: " + (data.error || 'Unknown error'));
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
            // user_role: user_role // Не нужно передавать здесь user_role, так как он будет возвращен сервером
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
            
            // Проверим, что в ответе есть нужные данные
            if (data.token && data.user_id && data.user_role) {
                // Сохраняем token, user_id и user_role в localStorage
                localStorage.setItem("token", data.token);
                localStorage.setItem("user_id", data.user_id);  
                localStorage.setItem("user_role", data.user_role.toString());  // Сохраняем user_role как строку

                // В зависимости от роли выполняем соответствующие действия
                if (data.user_role === 3) { // Пользователь — клиент
                    getClientIdFromUserId(data.user_id)
                    .then(clientId => {
                        localStorage.setItem("client_id", clientId);
                        alert("Login successful!");
                        window.location.href = "main.html";
                    })
                    .catch(error => {
                        console.error("Error fetching client_id:", error);
                        alert("Не удалось получить client_id");
                    });
                } else if (data.user_role === 1 || data.user_role === 2) { // Пользователь — сотрудник или администратор
                    getEmployeeIdFromUserId(data.user_id)
                    .then(employeeId => {
                        localStorage.setItem("employee_id", employeeId);
                        alert("Login successful!");
                        window.location.href = "main.html";
                    })
                    .catch(error => {
                        console.error("Error fetching employee_id:", error);
                        alert("Не удалось получить employee_id");
                    });
                } else {
                    alert("Неизвестная роль пользователя");
                }
            } else {
                alert("Error logging in: " + (data.error || 'Missing required data'));
            }
        })
        .catch(error => {
            console.error("Error during login:", error);
            alert("An error occurred during login");
        });
    });
});

// Получение client_id по user_id
function getClientIdFromUserId(userId) {
    return fetch(`http://localhost:8080/getClientId?user_id=${userId}`, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
    })
    .then(response => response.json())
    .then(data => {
        if (data.client_id) {
            return data.client_id;
        } else {
            throw new Error('Client not found');
        }
    })
    .catch(error => {
        console.error('Error fetching client_id:', error);
        throw error;
    });
}

function getEmployeeIdFromUserId(userId) {
    return fetch(`http://localhost:8080/getEmployeeId?user_id=${userId}`, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
    })
    .then(response => response.json())
    .then(data => {
        if (data.employee_id) {
            return data.employee_id;
        } else {
            throw new Error('Employee not found');
        }
    })
    .catch(error => {
        console.error('Error fetching client_id:', error);
        throw error;
    });
}
