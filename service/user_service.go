package service

import (
	pb "authentication-service/generated/user"
	"authentication-service/storage"
	"context"
	"errors"
	"log/slog"
)

var user storage.MainStorage

type UserService struct {
	pb.UnimplementedUserServiceServer
	log *slog.Logger
	storage.MainStorage
}

func (u *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashedPassword, err := HashPassword(in.Password)
	if err != nil {
		u.log.Error("Error in hashing password", "error", err)
		return nil, err
	}

	in.Password = hashedPassword

	res, err := u.NewUserStorage().CreateUser(in)
	if err != nil {
		u.log.Error("Failed to create user", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) GetUserProfile(ctx context.Context, in *pb.UserId) (*pb.UserResponse, error) {
	res, err := u.NewUserStorage().GetUserProfile(in)
	if err != nil {
		u.log.Error("Failed to get user profile", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, in *pb.FilterRequest) (*pb.UsersResponse, error) {
	res, err := u.NewUserStorage().GetAllUsers(in)
	if err != nil {
		u.log.Error("Failed to get all users", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res, err := u.NewUserStorage().UpdateUserProfile(in)
	if err != nil {
		u.log.Error("Failed to update user profile", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.Success, error) {
	hash, err := u.NewUserStorage().GetPassword(in)
	if err != nil {
		u.log.Error("Failed to get password", "error", err)
		return nil, err
	}

	check := CheckPasswordHash(hash, in.OldPassword)
	if !check {
		u.log.Warn("Incorrect Current password")
		return nil, errors.New("Incorrect Current password")
	}

	hashedPassword, err := HashPassword(in.NewPassword)
	if err != nil {
		u.log.Error("Failed to hash password", "error", err)
		return nil, err
	}

	in.NewPassword = hashedPassword

	err = u.NewUserStorage().ChangePassword(in)
	if err != nil {
		u.log.Error("Failed to update user profile", "error", err)
		return nil, err
	}

	return &pb.Success{Message: "Password updated"}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Success, error) {
	err := u.NewUserStorage().DeleteUser(in)
	if err != nil {
		u.log.Error("Failed to delete user", "error", err)
		return nil, err
	}

	return &pb.Success{Message: "User deleted"}, nil
}
