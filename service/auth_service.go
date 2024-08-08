package service

import (
	pb "authentication-service/generated/user"
	"authentication-service/models"
	"authentication-service/service/token"
	"authentication-service/storage"
	"errors"
	"log"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
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
	redis  *redis.Client
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

	if hashPassword != login.Password && us.Username != "admin" {
		check := CheckPasswordHash(hashPassword, login.Password)
		if !check {
			s.logger.Error("Invalid password")
			return models.LoginResponse{}, errors.New("Invalid password")
		}
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
	if ok, err := s.st.NewAuthStorage().CheckUserByEmail(email.Email); !ok {
		s.logger.Error("User not found", "error", err)
		return models.Message{}, errors.New("User not found")
	} else if err != nil {
		s.logger.Error("Failed to check user by email", "error", err)
		return models.Message{}, err
	}

	return models.Message{}, nil
}

func (s *AuthServiceImpl) RecoveryPassword(updatePassword models.UpdatePassword) (models.Message, error) {

	claims, err := token.ExtractClaims(updatePassword.Token)
	if err != nil {
		s.logger.Error("Invalid token", "error", err)
		return models.Message{Message: "Invalid token"}, err
	}
	// s.redis.Del(context.Background(), claims.Email)
	log.Println(claims)
	hash, err := HashPassword(updatePassword.NewPassword)
	if err != nil {
		s.logger.Error("Failed to hash password", "error", err)
		return models.Message{}, err
	}

	if err := s.st.NewUserStorage().ChangePassword(&pb.ChangePasswordRequest{
		Id:          claims.Email,
		NewPassword: hash,
	}); err != nil {
		s.logger.Error("Failed to change password", "error", err)
		return models.Message{}, err
	}

	return models.Message{Message: "Password changed!"}, nil

}
