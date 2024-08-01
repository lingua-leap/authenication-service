package postgres

import (
	"authentication-service/models"

	"github.com/jmoiron/sqlx"
)

type UserManagementSQLStorage struct {
	db *sqlx.DB
}

func NewUserManagementStorage(db *sqlx.DB) *UserManagementSQLStorage {
	return &UserManagementSQLStorage{db}
}

func (s *UserManagementSQLStorage) CreateUser(user *models.CreateUser) (*models.User, error) {
	return nil, nil
}
func (s *UserManagementSQLStorage) GetUser(id *string, email string, username *string) (*models.User, error) {
	return nil, nil
}

func (s *UserManagementSQLStorage) UpdateUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (s *UserManagementSQLStorage) DeleteUser(id string) error {
	return nil
}

func (s *UserManagementSQLStorage) GetAllUsers(page, limit int, query interface{}) ([]models.User, error) {
	return nil, nil
}
