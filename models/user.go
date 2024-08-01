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
	ID             uuid.UUID
	Username       string
	Email          string
	PasswordHash   string
	FullName       string
	NativeLanguage string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      int64
}

type AuthToken struct {
	Token string
}
