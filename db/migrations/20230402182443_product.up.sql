CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    category_id INT ,
    name VARCHAR(120) ,
    description TEXT,
    currency VARCHAR(30),
    price NUMERIC(17,2),
    stock INT,
    is_active BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (category_id) REFERENCES categories(id)
)