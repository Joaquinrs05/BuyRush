CREATE DATABASE notification_db;
\c notification_db;

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    mensaje TEXT NOT NULL,
    leido BOOLEAN DEFAULT FALSE,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);