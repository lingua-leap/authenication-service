package postgres

import (
	"authentication-service/models"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{db}
}

func (a *AuthStorage) Register(user models.CreateUser) (models.User, error) {
	res := models.User{}

	query := `
		INSERT INTO users (
			username,
			email,
			password_hash,
			full_name,
			native_language
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := a.db.QueryRow(query, user.Username, user.Email, user.HashedPassword, user.FullName, user.NativeLanguage).Scan(&res.ID)
	if err != nil {
		return res, err
	}

	res.Username = user.Username
	res.FullName = user.FullName
	res.NativeLanguage = user.NativeLanguage
	res.Email = user.Email

	return res, nil
}

func (a *AuthStorage) Login(login models.LoginRequest) (models.User, string, error) {
	var user models.User
	var password string

	err := a.db.QueryRow("select id, username, password_hash from users where username = $1 and deleted_at = 0", login.Username).
		Scan(&user.ID, &user.Username, &password)

	if errors.Is(err, sql.ErrNoRows) {
		return user, "", errors.New("user not found")
	}

	if err != nil {
		return user, "", err
	}
	return user, password, nil
}

func (a *AuthStorage) CheckUserByEmail(email string) (bool, error) {
	var id string
	err := a.db.QueryRow("select id from users where email = $1 and deleted_at = 0", email).
		Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return id != "", err
}
