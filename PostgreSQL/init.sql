CREATE TABLE IF NOT EXISTS sections (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    info TEXT NOT NULL,
    schedule TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    section VARCHAR(255),
    section_id INT,
    student_group VARCHAR(255),
    visits INT DEFAULT 0,
    paid BOOLEAN DEFAULT FALSE,
    last_scanned TIMESTAMP,
    qr_token TEXT UNIQUE,
    CONSTRAINT fk_users_section
        FOREIGN KEY (section_id)
        REFERENCES sections(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

INSERT INTO sections (name, info, schedule) VALUES
('Футбол', 'Тренер: Алексей Сидоров. Тренировки по футболу для всех возрастов', 'Пн, Ср, Пт 18:00-20:00'),
('Плавание', 'Тренер: Мария Петровна. Обучение плаванию и тренировки по технике', 'Вт, Чт 17:00-19:00'),
('Бокс', 'Тренер: Александр Самарцев. Тренировки по боксу для начинающих и профессионалов', 'Пн, Ср, Пт 19:00-21:00');

INSERT INTO users (login, name, surname, patronymic, password, role, section_id, student_group, visits, paid, last_scanned, qr_token) VALUES
('test', 'Иван', 'Иванов', 'Иванович', 'test', 'student', 1, 'КЭ-241', 5, TRUE, '2025-08-01 12:30:00', 'qr_001'),
('petrova2', 'Мария', 'Петрова', 'Сергеевна', 'pass456', 'student', 2, 'ЕТ-341', 3, FALSE, '2025-08-02 10:15:00', 'qr_002'),
('sidorov3', 'Алексей', 'Сидоров', NULL, 'pass789', 'teacher', 1, NULL, 0, TRUE, '2025-08-03 09:00:00', 'qr_003');
