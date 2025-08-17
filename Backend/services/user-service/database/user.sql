CREATE DATABASE user_db;
\c user_db;

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    photo TEXT,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);