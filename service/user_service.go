package service

import (
	"authentication-service/generated/auth"
	"authentication-service/storage"
	"context"
	"fmt"
	"time"
)

type UserService struct {
	userStorage storage.UserManagementStorage
	auth.UnimplementedAuthServiceServer
}

func (a *UserService) GetUserProfile(ctx context.Context, req *auth.GetUserProfileRequest) (*auth.UserResponse, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("InvalidIDError")
	}

	user, err := a.userStorage.GetUser(&req.UserId, nil, nil)
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

func (a *UserService) UpdateUserProfile(ctx context.Context, req *auth.UpdateUserProfileRequest) (*auth.UserResponse, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("InvalidIDError")
	}

	user, err := a.userStorage.GetUser(&req.UserId, nil, nil)

	if err != nil {
		return nil, err
	}

	user.FullName = req.GetFullName()
	user.NativeLanguage = req.GetNativeLanguage()

	updatedUser, err := a.userStorage.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &auth.UserResponse{
		Id:             updatedUser.ID.String(),
		Username:       updatedUser.Username,
		Email:          updatedUser.Email,
		FullName:       updatedUser.FullName,
		NativeLanguage: updatedUser.NativeLanguage,
		CreatedAt:      updatedUser.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (a *UserService) ChangeUserPasswordRequest(req *auth.ChangeUserPasswordRequest) (*auth.ChangeUserPasswordResponse, error) {
	return nil, nil
}
