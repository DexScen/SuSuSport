document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.querySelector('.login-dialog');
    const loginInput = document.querySelector('input[type="text"]');
    const passwordInput = document.querySelector('input[type="password"]');
    const submitButton = document.querySelector('.button-primary');

    loginForm.addEventListener('submit', async function(e) {
        e.preventDefault(); 
        
        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();

        if (!login || !password) {
            alert('Пожалуйста, заполните все поля');
            return;
        }

        const authData = {
            username: login,
            password: password
        };

        try {
            submitButton.disabled = true;
            submitButton.textContent = 'Отправка...';

            const response = await fetch('http://localhost:8081', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(authData)
            });

            if (!response.ok) {
                throw new Error(`Ошибка HTTP: ${response.status}`);
            }

            const data = await response.json();
            console.log('Успешная авторизация:', data);  
        } catch (error) {
            console.error('Ошибка при авторизации:', error);
            alert('Ошибка при авторизации. Проверьте логин и пароль.');
        } finally {
            submitButton.disabled = false;
            submitButton.textContent = 'Вход';
        }
    });

    submitButton.addEventListener('click', function() {
        loginForm.dispatchEvent(new Event('submit'));
    });
});