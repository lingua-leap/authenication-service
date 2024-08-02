package storage

import "authentication-service/models"

type AuthenticationStorage interface {
	Register(user *models.CreateUser) (*models.User, error)
	Login(email, Password string) (*models.AuthToken, error)
	VerifyToken(token string) (*models.User, error)
	ResetTokenToEmail(email string) (*models.Message, error)
	UpdatePasswordByToken(res *models.UpdatePassword) (*models.Message, error)
}
