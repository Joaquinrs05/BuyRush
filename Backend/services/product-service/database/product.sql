CREATE DATABASE product_db;
\c product_db;

CREATE TABLE produsctos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL,
    precio NUMERIC(10,2) NOT NULL,
    descripcion TEXT,
    foto TEXT,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);