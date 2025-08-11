document.addEventListener('DOMContentLoaded', function () {
    const exitButtons = document.querySelectorAll('.exit');
    const signButton = document.querySelector('.sign-button');
    const titleElement = document.querySelector('.title p');
    const descriptionElement = document.querySelector('[data-info]');
    const scheduleElement = document.querySelector('[data-schedule]');

    // Кнопки "Назад" и "Выход"
    if (exitButtons.length >= 2) {
        const backButton = exitButtons[0];
        const logoutButton = exitButtons[1];

        backButton.addEventListener('click', function () {
            window.history.back();
        });

        logoutButton.addEventListener('click', function () {
            if (confirm('Вы уверены, что хотите выйти?')) {
                window.location.href = '/index.html';
            }
        });
    }

    // Получаем параметр name из URL
    const urlParams = new URLSearchParams(window.location.search);
    const sectionId = urlParams.get('name');

    if (!sectionId) {
        titleElement.textContent = 'Секция не найдена';
        descriptionElement.textContent = 'Не указан идентификатор секции.';
        scheduleElement.textContent = 'Не указано расписание.';
        return;
    }

    async function loadSection() {
        try {
            const response = await fetch(`http://localhost:8080/sport/sections/${encodeURIComponent(sectionId)}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error(`Ошибка HTTP: ${response.status}`);
            }

            const b = await response.json();

            titleElement.textContent = b.name || 'Без названия';
            descriptionElement.textContent = b.info || 'Описание отсутствует';
            scheduleElement.textContent = b.schedule || 'Расписание отсутствует';
        } catch (error) {
            console.error('Ошибка загрузки секции:', error);
            titleElement.textContent = 'Ошибка загрузки';
            descriptionElement.textContent = 'Не удалось получить данные секции.';
            scheduleElement.textContent = 'Не удалось получить расписание секции.';
        }
    }

    loadSection();
});
