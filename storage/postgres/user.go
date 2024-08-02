package postgres

import (
	"authentication-service/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserManagementSQLStorage struct {
	db *sqlx.DB
}

func NewUserManagementStorage(db *sqlx.DB) *UserManagementSQLStorage {
	return &UserManagementSQLStorage{db}
}

func (s *UserManagementSQLStorage) GetUser(id *string, email *string, username *string) (*models.User, error) {
	var user models.User
	var query string
	var err error

	switch {
	case id != nil:
		query = `
			SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
			FROM users
			WHERE id = $1 AND deleted_at IS NULL
			LIMIT 1
		`
		err = s.db.Get(&user, query, *id)
	case email != nil:
		query = `
			SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
			FROM users
			WHERE email = $1 AND deleted_at IS NULL
			LIMIT 1
		`
		err = s.db.Get(&user, query, email)
	case username != nil:
		query = `
			SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
			FROM users
			WHERE username = $1 AND deleted_at IS NULL
			LIMIT 1
		`
		err = s.db.Get(&user, query, *username)
	default:
		return nil, fmt.Errorf("no criteria provided for GetUser")
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	return &user, nil
}

func (s *UserManagementSQLStorage) UpdateUser(user *models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET full_name = $1, native_language = $2, updated_at = NOW()
		WHERE id = $3 AND deleted_at = 0
		RETURNING id, username, email, full_name, native_language, created_at, updated_at
	`

	row := s.db.QueryRow(query, user.FullName, user.NativeLanguage, user.ID)
	var updatedUser models.User
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email,
		&updatedUser.FullName, &updatedUser.NativeLanguage,
		&updatedUser.CreatedAt, &updatedUser.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return &updatedUser, nil
}

func (s *UserManagementSQLStorage) DeleteUser(id string) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found or already deleted")
	}

	return nil
}

func (s *UserManagementSQLStorage) GetAllUsers(page, limit int, query interface{}) ([]models.User, error) {
	offset := (page - 1) * limit

	queryString := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var users []models.User
	err := s.db.Select(&users, queryString, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}

	return users, nil
}

func (s *UserManagementSQLStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE email = $1 AND deleted_at = 0
		LIMIT 1
	`

	err := s.db.Get(&user, query, email)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("Error fetching user by email: %v", err)
	}

	return &user, nil
}
