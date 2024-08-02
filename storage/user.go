package storage

import "authentication-service/models"

type UserManagementStorage interface {
	GetUser(id *string, email *string, username *string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	GetAllUsers(page, limit int, query interface{}) ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
