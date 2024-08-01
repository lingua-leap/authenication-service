package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateUser struct {
	Username       string
	Email          string
	Password       string
	FullName       string
	NativeLanguage string
}

type User struct {
	ID             uuid.UUID `db:"id"`
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	FullName       string    `db:"full_name"`
	NativeLanguage string    `db:"native_language"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      int64     `db:"deleted_at"`
}

type AuthToken struct {
	Token string
}
