-- User Service
CREATE DATABASE user_db;
\c user_db

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    photo TEXT,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Product Service
CREATE DATABASE product_db;
\c product_db

CREATE TABLE productos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL,
    precio NUMERIC(10,2) NOT NULL,
    descripcion TEXT,
    foto TEXT,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Order Service
CREATE DATABASE order_db;
\c order_db

CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    estado VARCHAR(50) DEFAULT 'pendiente',
    total NUMERIC(10,2) NOT NULL,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pedido_productos (
    id SERIAL PRIMARY KEY,
    pedido_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL DEFAULT 1
);

-- Payment Service
CREATE DATABASE payment_db;
\c payment_db

CREATE TABLE facturacion (
    id SERIAL PRIMARY KEY,
    pedido_id INT NOT NULL,
    metodo_pago VARCHAR(50) NOT NULL,
    monto NUMERIC(10,2) NOT NULL,
    fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notification Service
CREATE DATABASE notification_db;
\c notification_db

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    mensaje TEXT NOT NULL,
    leido BOOLEAN DEFAULT FALSE,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);






















/* -- Crear base de datos
CREATE DATABASE BuyRush;

-- Conectarse a la base de datos
\c BuyRush;

-- Tabla Usuarios
CREATE TABLE usuarios (
                          id SERIAL PRIMARY KEY,
                          Name VARCHAR(100) NOT NULL,
                          Surname VARCHAR(100) NOT NULL,
                          Email VARCHAR(150) UNIQUE NOT NULL,
                          Password TEXT NOT NULL,
                          photo TEXT,
                          creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla Productos
CREATE TABLE productos (
                           id SERIAL PRIMARY KEY,
                           nombre VARCHAR(150) NOT NULL,
                           precio NUMERIC(10,2) NOT NULL,
                           descripcion TEXT,
                           foto TEXT,
                           creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla Carrito
CREATE TABLE carritos (
                          id SERIAL PRIMARY KEY,
                          usuario_id INT REFERENCES usuarios(id) ON DELETE CASCADE,
                          total NUMERIC(10,2) DEFAULT 0,
                          creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla Carrito_Productos (relación muchos a muchos)
CREATE TABLE carrito_productos (
                                   id SERIAL PRIMARY KEY,
                                   carrito_id INT REFERENCES carritos(id) ON DELETE CASCADE,
                                   producto_id INT REFERENCES productos(id) ON DELETE CASCADE,
                                   cantidad INT NOT NULL DEFAULT 1
);

-- Tabla Pedidos
CREATE TABLE pedidos (
                         id SERIAL PRIMARY KEY,
                         usuario_id INT REFERENCES usuarios(id) ON DELETE CASCADE,
                         estado VARCHAR(50) DEFAULT 'pendiente',
                         total NUMERIC(10,2) NOT NULL,
                         creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla Pedido_Productos (relación muchos a muchos)
CREATE TABLE pedido_productos (
                                  id SERIAL PRIMARY KEY,
                                  pedido_id INT REFERENCES pedidos(id) ON DELETE CASCADE,
                                  producto_id INT REFERENCES productos(id) ON DELETE CASCADE,
                                  cantidad INT NOT NULL DEFAULT 1
);

-- Tabla Facturación
CREATE TABLE facturacion (
                             id SERIAL PRIMARY KEY,
                             pedido_id INT REFERENCES pedidos(id) ON DELETE CASCADE,
                             metodo_pago VARCHAR(50) NOT NULL,
                             monto NUMERIC(10,2) NOT NULL,
                             fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); */
