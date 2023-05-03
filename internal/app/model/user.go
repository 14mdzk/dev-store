package model

import "time"

type User struct {
	ID        int `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type UserAddress struct {
	ID          int        `db:"id"`
	UserId      int        `db:"user_id"`
	AddressLine string     `db:"address_line"`
	Country     string     `db:"country"`
	City        string     `db:"city"`
	PostalCode  string     `db:"postal_code"`
	Phone       string     `db:"phone"`
	Note        string     `db:"note"`
	IsMain      bool       `db:"is_main"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
