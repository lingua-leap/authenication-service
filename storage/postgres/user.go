package postgresStorage

import (
	"authentication-service/models"
	"authentication-service/storage"

	"github.com/jmoiron/sqlx"
)

type UserManagementStorage struct {
	db *sqlx.DB
}

func NewUserManagementStorage(db *sqlx.DB) storage.UserManagementStorage {
	return &UserManagementStorage{db}
}

func (s *UserManagementStorage) CreateUser(user *models.CreateUser) (*models.User, error) {
	return nil, nil
}
func (s *UserManagementStorage) GetUser(id *string, email string, username *string) (*models.User, error) {
	return nil, nil
}

func (s *UserManagementStorage) UpdateUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (s *UserManagementStorage) DeleteUser(id string) error {
	return nil
}

func (s *UserManagementStorage) GetAllUsers(page, limit int, query interface{}) ([]models.User, error) {
	return nil, nil
}
