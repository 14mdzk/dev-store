package model

import "time"

type Product struct {
	ID          int        `db:"id"`
	CategoryId  int        `db:"category_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Currency    string     `db:"currency"`
	Price       float32    `db:"price"`
	Stock       int        `db:"stock"`
	IsActive    bool       `db:"is_active"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
