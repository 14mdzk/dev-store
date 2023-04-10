CREATE TABLE IF NOT EXISTS shopping_cart_items(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    shopping_cart_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,

    FOREIGN KEY(shopping_cart_id) REFERENCES shopping_carts(id),
    FOREIGN KEY(product_id) REFERENCES products(id)
)