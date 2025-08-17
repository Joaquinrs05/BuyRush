CREATE DATABASE review_db;
\c review_db;

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    producto_id INT NOT NULL,
    usuario_id INT NOT NULL,
    comentario TEXT,
    puntuacion INT CHECK (puntuacion >= 1 AND puntuacion <= 5),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);