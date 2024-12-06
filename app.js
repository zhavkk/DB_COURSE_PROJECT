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
            if (data.token) {
                localStorage.setItem("token", data.token);
                alert("Login successful!");
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
