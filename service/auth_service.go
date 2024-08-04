package service

import (
	"authentication-service/models"
	"authentication-service/service/token"
	"authentication-service/storage"
	"errors"
	"log/slog"
	"time"
)

type AuthService interface {
	Register(user models.CreateUser) (models.User, error)
	Login(login models.LoginRequest) (models.LoginResponse, error)
	RefreshToken(claims *token.Claims) (models.RefreshResponse, error)
	ResetTokenToEmail(email models.Email) (models.Message, error)
	RecoveryPassword(password models.UpdatePassword) (models.Message, error)
}

type AuthServiceImpl struct {
	logger *slog.Logger
	st     storage.MainStorage
}

func NewAuthService(logger *slog.Logger, st storage.MainStorage) AuthService {
	return &AuthServiceImpl{logger: logger, st: st}
}

func (s *AuthServiceImpl) Register(user models.CreateUser) (models.User, error) {
	hash, err := HashPassword(user.HashedPassword)
	if err != nil {
		s.logger.Error("Failed to hash password", "error", err)
		return models.User{}, err
	}

	user.HashedPassword = hash

	res, err := s.st.NewAuthStorage().Register(user)
	if err != nil {
		s.logger.Error("Failed to register user", "error", err)
		return models.User{}, err
	}
	return res, nil
}

func (s *AuthServiceImpl) Login(login models.LoginRequest) (models.LoginResponse, error) {
	us, hashPassword, err := s.st.NewAuthStorage().Login(login)
	if err != nil {
		s.logger.Error("Failed to login", "error", err)
		return models.LoginResponse{}, err
	}

	check := CheckPasswordHash(hashPassword, login.Password)
	if !check {
		s.logger.Error("Invalid password")
		return models.LoginResponse{}, errors.New("Invalid password")
	}

	accessToken, err := token.GenerateAccessToken(us)
	if err != nil {
		s.logger.Error("Failed to generate access token", "error", err)
		return models.LoginResponse{}, err
	}

	refreshToken, err := token.GenerateRefreshToken(us)
	if err != nil {
		s.logger.Error("Failed to generate refresh token", "error", err)
		return models.LoginResponse{}, err
	}

	loginRespnse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Minute * 10),
	}

	return loginRespnse, nil
}

func (s *AuthServiceImpl) RefreshToken(claims *token.Claims) (models.RefreshResponse, error) {
	usr := models.User{
		ID:       claims.ID,
		Username: claims.Username,
		Email:    claims.Username,
	}

	tk, err := token.GenerateAccessToken(usr)
	if err != nil {
		s.logger.Error("Failed to generate access token", "error", err)
		return models.RefreshResponse{}, err
	}

	refreshResponse := models.RefreshResponse{
		AccessToken: tk,
	}

	return refreshResponse, nil
}
func (s *AuthServiceImpl) ResetTokenToEmail(email models.Email) (models.Message, error) {
	return models.Message{}, nil
}

func (s *AuthServiceImpl) RecoveryPassword(password models.UpdatePassword) (models.Message, error) {
	return models.Message{}, nil
}
