package model

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
