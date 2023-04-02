CREATE TABLE IF NOT EXISTS user_addresses(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    address_line TEXT,
    country VARCHAR(50),
    city VARCHAR(50),
    postal_code VARCHAR(20),
    phone VARCHAR(20),
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,

    FOREIGN KEY (user_id) REFERENCES users(id)
)