package models

import (
	"time"
)

type CreateUser struct {
	Username       string
	Email          string
	HashedPassword string
	FullName       string
	NativeLanguage string
}

type User struct {
	ID             string    `db:"id" json:"id"`
	Username       string    `db:"username" json:"username"`
	Email          string    `db:"email" json:"email"`
	FullName       string    `db:"full_name" json:"full_name"`
	NativeLanguage string    `db:"native_language" json:"native_language"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt      int64     `db:"deleted_at" json:"deleted_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Message struct {
	Message string `json:"message"`
}

type RefreshResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type UpdatePassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type Email struct {
	Email string `json:"email"`
}
