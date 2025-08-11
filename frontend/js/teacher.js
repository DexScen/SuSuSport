document.addEventListener('DOMContentLoaded', function () {
    const qrReader = new Html5Qrcode("qr-reader");
    const startButton = document.getElementById('startScanner');
    const stopButton = document.getElementById('stopScanner');
    const exitButtons = document.querySelectorAll('.exit');

    if (exitButtons.length >= 1) {
        exitButtons[0].addEventListener('click', function () {
            if (confirm('Вы уверены, что хотите выйти?')) {
                window.location.href = '/index.html';
            }
        });
    }

    startButton.addEventListener('click', function () {
        const config = { 
            fps: 10,
            qrbox: 250
        };

        qrReader.start(
            { facingMode: "environment" }, 
            config,
            function (decodedText) {
                alert(`Ученик успешно отмечен! QR-код: ${decodedText}`);
                qrReader.stop();
                startButton.disabled = false;
                stopButton.disabled = true;
            },
            function (errorMessage) {
                console.log(errorMessage);
            }
        ).then(() => {
            startButton.disabled = true;
            stopButton.disabled = false;
        }).catch(err => {
            alert(`Ошибка: ${err}`);
        });
    });

    stopButton.addEventListener('click', function () {
        qrReader.stop()
            .then(() => {
                startButton.disabled = false;
                stopButton.disabled = true;
            })
            .catch(err => {
                alert(`Ошибка при остановке: ${err}`);
            });
    });

     function handleSuccessfulScan(qrData) {
        try {
            const studentData = JSON.parse(qrData);
        
            const message = `
                Студент успешно отмечен!
                ФИО: ${studentData.surname} ${studentData.name} ${studentData.patronymic}
                Группа: ${studentData.student_group}
                Секция: ${studentData.section || 'Не указана'}
            `;
            
            showSuccessNotification(message);
            
            // Здесь можно отправить данные на сервер
            // sendAttendanceToServer(studentData);
            
        } catch (e) {
            alert('Ошибка: Неверный QR-код студента');
        }
    }

    function showSuccessNotification(message) {
        const notification = document.createElement('div');
        notification.className = 'scan-notification';
        notification.innerHTML = `
            <div class="notification-content">
                <img src="assets/success-icon.svg" alt="Успех">
                <p>${message}</p>
            </div>
        `;
        document.body.appendChild(notification);
        
        setTimeout(() => {
            notification.remove();
        }, 3000);
    }
});