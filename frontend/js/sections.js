document.addEventListener('DOMContentLoaded', function () {
    const buttonsContainer = document.querySelector('[data-btn-container]');
    const exitButtons = document.querySelectorAll('.exit');

    // Обработка кнопок "Назад" и "Выход"
    if (exitButtons.length >= 2) {
        const backButton = exitButtons[0];
        const logoutButton = exitButtons[1];

        backButton.addEventListener('click', function () {
            window.history.back();
        });

        logoutButton.addEventListener('click', function () {
            if (confirm('Вы уверены, что хотите выйти?')) {
                window.location.href = '/index.html'; // редирект на страницу входа
            }
        });
    }

    async function loadSections() {
        try {
            const response = await fetch('http://localhost:8080/sport/sections', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error(`Ошибка HTTP: ${response.status}`);
            }

            const sections = await response.json();
            renderSections(sections);
        } catch (error) {
            console.error('Ошибка при получении секций:', error);
            buttonsContainer.innerHTML = '<p>Не удалось загрузить секции</p>';
        }
    }

    function renderSections(sections) {
        // Очищаем контейнер, но не удаляем его
        buttonsContainer.innerHTML = '';

        if (!Array.isArray(sections) || sections.length === 0) {
            buttonsContainer.textContent = 'Секции отсутствуют';
            return;
        }

        sections.forEach(section => {
            const btn = document.createElement('button');
            btn.classList.add('button');

            // Определяем имя секции
            const sectionName = section;
            btn.textContent = sectionName;

            // При клике переходим на description.html с параметром name
            btn.addEventListener('click', () => {
                const encodedName = encodeURIComponent(sectionName);
                window.location.href = `description.html?name=${encodedName}`;
            });

            buttonsContainer.appendChild(btn);
        });
    }

    loadSections();
});
