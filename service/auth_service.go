package service

import "authentication-service/models"

type AuthService interface {
	Register(user *models.CreateUser) (*models.User, error)
	Login(email, Password string) (*models.AuthToken, error)
	VerifyToken(token string) (*models.User, error)
	ResetTokenToEmail(email string) (*models.Message, error)
	UpdatePasswordByToken(res *models.UpdatePassword) (*models.Message, error)
}
