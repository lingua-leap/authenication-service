package service

import (
	"authentication-service/genproto/auth"
	"authentication-service/storage"
	"context"
)

type MainService interface {
	AuthService() AuthenticationService
	UserService() UserManagementService
}

type AuthenticationService interface {
	RegisterUser(ctx context.Context, req *auth.RegisterUserRequest) (*auth.UserResponse, error)
	LoginUser(ctx context.Context, req *auth.LoginUserRequest) (*auth.LoginUserResponse, error)
	UpdateUserProfile(ctx context.Context, req *auth.UpdateUserProfileRequest) (*auth.UserResponse, error)
	RequestPasswordReset(ctx context.Context, req *auth.RequestPasswordResetRequest) (*auth.RequestPasswordResetResponse, error)
	ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
	LogoutUser(ctx context.Context, req *auth.LogoutUserRequest) (*auth.LogoutUserResponse, error)
}

type UserManagementService interface {
	UpdateUserProfile(ctx context.Context, req *auth.UpdateUserProfileRequest) (*auth.UserResponse, error)
	ChangeUserPassword(ctx context.Context, req *auth.ChangeUserPasswordRequest) (*auth.ChangeUserPasswordResponse, error)
	GetAllUsersProfile(ctx context.Context, req *auth.GetAllUsersProfileRequest) (*auth.UsersResponse, error)
	GetUserProfile(ctx context.Context, req *auth.GetUserProfileRequest) (*auth.UserResponse, error)
}

func NewMainService(mainStorage storage.MainStorage) MainService {
	return &mainService{mainStorage: mainStorage}
}

type mainService struct {
	mainStorage storage.MainStorage
}

func (s *mainService) AuthService() AuthenticationService {
	return &authenticationService{
		authStorage: s.mainStorage.AuthenticationStorage(),
	}
}

func (s *mainService) UserService() UserManagementService {
	return &UserService{
		userStorage: s.mainStorage.UserManagementStorage(),
	}
}
