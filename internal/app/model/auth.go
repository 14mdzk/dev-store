package model

import "time"

type Auth struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	AuthType  string    `db:"auth_type"`
	ExpiredAt time.Time `db:"expired_at"`
}
