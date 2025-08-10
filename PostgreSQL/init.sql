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