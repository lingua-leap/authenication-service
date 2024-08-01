package storage

import "authentication-service/models"

type UserManagementStorage interface {
	CreateUser(user *models.CreateUser) (*models.User, error)
	GetUser(id *string, email string, username *string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	GetAllUsers(page, limit int, query interface{}) ([]models.User, error)
}
