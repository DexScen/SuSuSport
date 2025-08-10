document.querySelector('.login-dialog').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const form = e.target;
    const button = form.querySelector('.button-primary');
    const login = form.querySelector('input[type="text"]').value.trim();
    const password = form.querySelector('input[type="password"]').value.trim();

    if (!login || !password) {
        alert('Заполните все поля');
        return;
    }

    button.disabled = true;
    button.textContent = 'Отправка...';

    try {
        const response = await fetch('http://localhost:8081', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username: login, password })
        });

        if (!response.ok) throw new Error('Ошибка сервера');
        const data = await response.json();
        console.log('Успех:', data);
        
    } catch (error) {
        console.error('Ошибка:', error);
        alert('Ошибка авторизации');
    } finally {
        button.disabled = false;
        button.textContent = 'Вход';
    }
});