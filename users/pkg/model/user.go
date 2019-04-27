package model

import (
	"time"
)

// User Entity
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Company   string    `json:"company" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Token     string    `json:"token"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Token Entity
type Token struct {
	Token     string    `json:"token" validate:"required"`
	Valid     bool		`json:"valid"`
}
