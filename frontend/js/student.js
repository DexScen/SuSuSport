const data = JSON.parse(localStorage.getItem('myData'));

document.addEventListener('DOMContentLoaded', function () {
    
    // Элементы на странице
    const fullNameElem = document.querySelector('.student-info h1');
    const sectionElem = document.getElementById('section');
    const attendanceElem = document.getElementById('attendance');
    const scheduleElem = document.getElementById('schedule');
    const qrContainer = document.querySelector('.qr-code-container');
    const infoButton = document.querySelector('.info-button');
    const exitButtons = document.querySelectorAll('.exit');

    if (exitButtons.length >= 1) {
        const logoutButton = exitButtons[0];

        logoutButton.addEventListener('click', function () {
            if (confirm('Вы уверены, что хотите выйти?')) {
                window.location.href = '/index.html';
            }
        });
    }

    // Заполняем ФИО и группу
    fullNameElem.textContent = `${data.surname} ${data.name} ${data.patronymic}, ${data.student_group}`;

    // Заполняем остальные данные
    sectionElem.textContent = data.section || 'Футбол';
    attendanceElem.textContent = data.visits !== undefined ? data.visits : '-';
    scheduleElem.textContent = 'Пн, Ср, Пт 18:00-20:00'; // Можно потом заменить на реальное расписание

    // Генерируем QR-код из data.qrcode
    // Сначала удалим старое изображение, если есть
    qrContainer.innerHTML = '';

    // Создаем контейнер для QR кода
    const qrDiv = document.createElement('div');
    qrContainer.appendChild(qrDiv);

    // Генерируем QR код
    new QRCode(qrDiv, {
        text: data.qrcode,
        width: 128,
        height: 128
    });

    // Кнопка "Подробнее о секциях" ведет на страницу секций
    infoButton.addEventListener('click', () => {
        window.location.href = 'sections.html';
    });
});
