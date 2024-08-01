package storage

import "authentication-service/models"

type AuthenticationStorage interface {
	Register(user *models.CreateUser) (*models.User, error)
	Login(email, Password string) (*models.AuthToken, error)
	Logout(userID string) error
	VerifyToken(token string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
