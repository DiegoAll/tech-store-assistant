-- Habilitar pgvector (por si acaso, aunque ya lo hace la imagen)
CREATE EXTENSION IF NOT EXISTS vector;

-- 1. Tabla de Pedidos (Caso de uso 1)
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(50) PRIMARY KEY,
    status VARCHAR(100) NOT NULL,
    items TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO orders (id, status, items) VALUES 
('1234', 'En camino. El repartidor está a 2 paradas de distancia.', 'SmartHome Hub X1, Cámara SecureCam Pro'),
('5678', 'Entregado el 24 de Octubre.', 'Termostato Smart'),
('9999', 'En preparación en bodega.', 'Suscripción Anual de Monitoreo')
ON CONFLICT (id) DO NOTHING;

-- 2. Tabla de Inventario/Stock (Caso de uso 2)
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    stock INTEGER NOT NULL DEFAULT 0,
    price NUMERIC(10, 2) NOT NULL
);

INSERT INTO products (name, stock, price) VALUES
('SmartHome Hub X1', 15, 150.00),
('Cámara SecureCam Pro', 0, 89.00),
('Termostato Smart', 5, 120.00),
('Sensor de Movimiento', 30, 25.00)
ON CONFLICT (name) DO NOTHING;

-- 3. Tabla de Tickets de Soporte (Caso de uso 3)
CREATE TABLE IF NOT EXISTS support_tickets (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    issue TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'OPEN',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Tabla de Documentos RAG (Caso de uso 5)
CREATE TABLE IF NOT EXISTS tienda_docs (
    id SERIAL PRIMARY KEY,
    content TEXT,
    embedding vector(1536)
);

-- Insertamos los documentos base (el embedding se generará desde el backend en Go)
INSERT INTO tienda_docs (content) VALUES
('El SmartHome Hub X1 es compatible con Alexa y Google Home. Cuesta $150 USD.'),
('La Cámara SecureCam Pro se conecta por WiFi de 2.4GHz. Si no conecta, reinicia el router.'),
('Ofrecemos servicio de instalación premium por $50 USD adicionales.'),
('El soporte técnico está disponible 24/7 a través de nuestro chat web o enviando un correo a soporte@tienda-tech.com.'),
('Política de devoluciones: Aceptamos devoluciones dentro de los primeros 30 días.')
ON CONFLICT DO NOTHING;