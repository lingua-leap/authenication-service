package storage

import "authentication-service/models"

type AuthStorage interface {
	Register(user models.CreateUser) (models.User, error)
	Login(login models.LoginRequest) (models.User, string, error)
	CheckUserByEmail(email string) (bool, error)
}
