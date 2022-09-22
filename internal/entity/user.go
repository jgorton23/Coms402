package entity

import "time"

type User struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`

	Password             string `json:"-"`
	PasswordConfirmation string `json:"-"`
}
