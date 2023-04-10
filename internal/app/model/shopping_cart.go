package model

import "time"

type ShoppingCart struct {
	ID        int     `db:"id"`
	UserId    int     `db:"user_id"`
	Total     float32 `db:"total"`
	Items     []ShoppingCartItem
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type ShoppingCartItem struct {
	ID             int        `db:"id"`
	ShoppingCartId int        `db:"shopping_cart_id"`
	ProductId      int        `db:"product_id"`
	Quantity       int        `db:"quantity"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}
