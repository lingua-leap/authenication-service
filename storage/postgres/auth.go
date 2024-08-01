package postgresStorage

import (
	"authentication-service/models"
	"authentication-service/storage"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationStorage struct {
	db *sqlx.DB
}

func NewAuthenticationStorage(db *sqlx.DB) storage.AuthenticationStorage {
	return &AuthenticationStorage{db}
}

func (a *AuthenticationStorage) Register(user *models.CreateUser) (*models.User, error) {
	_, err := a.GetUserByEmail(user.Email)
	if err == nil {
		return nil, fmt.Errorf("EmailAlreadyExists")
	} else if err != sql.ErrNoRows {
		return nil, fmt.Errorf("Error registering user: %v", err)
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("Error hashing password: %v", err)
	}

	newUUID := uuid.NewString()
	query := `
		INSERT INTO users (
			id,
			username,
			email,
			password_hash,
			full_name,
			native_language
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, username, email, password_hash, full_name, native_language, created_at, updated_at
	`

	row := a.db.QueryRow(query, newUUID, user.Username, user.Email, hashedPassword, user.FullName, user.NativeLanguage)

	var createdUser models.User
	err = row.Scan(&createdUser.ID, &createdUser.Username, &createdUser.Email,
		&createdUser.PasswordHash, &createdUser.FullName, &createdUser.NativeLanguage,
		&createdUser.CreatedAt, &createdUser.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (a *AuthenticationStorage) Login(email, password string) (*models.AuthToken, error) {
	var user models.User
	query := `
	    SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE email = $1 AND deleted_at = 0
		LIMIT 1
	`

	err := a.db.Get(&user, query, email)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("InvalidEmail")
	} else if err != nil {
		return nil, fmt.Errorf("Error logging in: %v", err)
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("InvalidPassword")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("Error generating token: %v", err)
	}

	return &models.AuthToken{Token: token}, nil
}

func (a *AuthenticationStorage) Logout(userID string) error {
	// Implement logout logic, e.g., invalidate the token
	return nil
}

func (a *AuthenticationStorage) VerifyToken(token string) (*models.User, error) {
	userID, err := ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("Error verifying token: %v", err)
	}

	var user models.User
	query := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted_at = 0
		LIMIT 1
	`

	err = a.db.Get(&user, query, userID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("InvalidToken")
	} else if err != nil {
		return nil, fmt.Errorf("Error fetching user: %v", err)
	}

	return &user, nil
}

func (a *AuthenticationStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE email = $1 AND deleted_at = 0
		LIMIT 1
	`

	err := a.db.Get(&user, query, email)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("Error fetching user by email: %v", err)
	}

	return &user, nil
}

// Utility functions for password hashing and token management
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

func GenerateToken(userID uuid.UUID) (string, error) {
	// Implement token generation, e.g., using JWT
	return "", nil
}

func ParseToken(token string) (uuid.UUID, error) {
	// Implement token parsing
	return uuid.UUID{}, nil
}