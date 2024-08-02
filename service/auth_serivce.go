package service

import (
	"authentication-service/generated/auth"
	"authentication-service/models"
	"authentication-service/storage"
	"context"
	"log/slog"
	"time"
)

type authenticationService struct {
	authStorage storage.AuthenticationStorage
	userStorage storage.UserManagementStorage
	auth.UnimplementedAuthServiceServer
	log *slog.Logger
}

func NewAuthenticationService(authStorage storage.AuthenticationStorage, userStorage storage.UserManagementStorage) *authenticationService {
	return &authenticationService{
		authStorage: authStorage,
		userStorage: userStorage,
	}
}

func (a *authenticationService) RegisterUser(ctx context.Context, req *auth.RegisterUserRequest) (*auth.UserResponse, error) {

	e, err := ValidateEmail(req.Email)
	if e == "" {
		a.log.Error("Invalid email address", err)
		return nil, err
	}

	p, err := ValidatePassword(req.Password)
	if p == "" {
		a.log.Error("Invalid password", err)
		return nil, err
	}

	hashedPassord, err := HashPassword(req.Password)
	if err != nil {
		a.log.Error(err.Error())
		return nil, err
	}

	createUser := models.CreateUser{
		Email:          req.GetEmail(),
		HashedPassword: hashedPassord,
		Username:       req.GetUsername(),
		FullName:       req.GetFullName(),
		NativeLanguage: req.GetNativeLanguage(),
	}

	user, err := a.authStorage.Register(&createUser)
	if err != nil {
		return nil, err
	}

	return &auth.UserResponse{
		Id:             user.ID.String(),
		Username:       user.Username,
		Email:          user.Email,
		FullName:       user.FullName,
		NativeLanguage: user.NativeLanguage,
		CreatedAt:      user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (a *authenticationService) LoginUser(ctx context.Context, req *auth.LoginUserRequest) (*auth.LoginUserResponse, error) {
	token, err := a.authStorage.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &auth.LoginUserResponse{
		Token: token.Token,
	}, nil
}

func (a *authenticationService) RequestPasswordReset(ctx context.Context, req *auth.RequestPasswordResetRequest) (*auth.RequestPasswordResetResponse, error) {
	// token, err := a.authStorage.RequestPasswordReset(req.GetEmail())
	// if err != nil {
	// 	return nil, err
	// }

	// return &auth.RequestPasswordResetResponse{
	// 	Token: token.Token,
	// }, nil
	return nil, nil
}

func (a *authenticationService) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	// err := a.authStorage.ResetPassword(req.GetEmail(), req.GetPassword())
	// if err != nil {
	// 	return nil, err
	// }

	// return &auth.ResetPasswordResponse{}, nil
	return nil, nil
}

func (a *authenticationService) LogoutUser(ctx context.Context, req *auth.LogoutUserRequest) (*auth.LogoutUserResponse, error) {
	// err := a.authStorage.(req.GetToken())
	// if err != nil {
	// 	return nil, err
	// }

	return &auth.LogoutUserResponse{}, nil
}
