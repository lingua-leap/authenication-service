package service

import (
	pb "authentication-service/generated/user"
	"authentication-service/storage"
	"context"
	"errors"
	"log"
	"log/slog"
)

var user storage.MainStorage

type UserService struct {
	pb.UnimplementedUserServiceServer
	log *slog.Logger
	st  storage.MainStorage
}

func NewUserService(log *slog.Logger, storage storage.MainStorage) *UserService {
	return &UserService{log: log}
}

func (u *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashedPassword, err := HashPassword(in.Password)
	if err != nil {
		u.log.Error("Error in hashing password", "error", err)
		return nil, err
	}

	in.Password = hashedPassword
	log.Println("HELLO WORLD 111111111111111111")
	res, err := u.st.NewUserStorage().CreateUser(in)
	log.Println("HELLO WORLD 222222222222222222222")

	if err != nil {
		u.log.Error("Failed to create user", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) GetUserProfile(ctx context.Context, in *pb.UserId) (*pb.UserResponse, error) {
	res, err := u.st.NewUserStorage().GetUserProfile(in)
	if err != nil {
		u.log.Error("Failed to get user profile", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, in *pb.FilterRequest) (*pb.UsersResponse, error) {
	res, err := u.st.NewUserStorage().GetAllUsers(in)
	if err != nil {
		u.log.Error("Failed to get all users", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res, err := u.st.NewUserStorage().UpdateUserProfile(in)
	if err != nil {
		u.log.Error("Failed to update user profile", "error", err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.Success, error) {
	hash, err := u.st.NewUserStorage().GetPassword(in)
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

	err = u.st.NewUserStorage().ChangePassword(in)
	if err != nil {
		u.log.Error("Failed to update user profile", "error", err)
		return nil, err
	}

	return &pb.Success{Message: "Password updated"}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Success, error) {
	err := u.st.NewUserStorage().DeleteUser(in)
	if err != nil {
		u.log.Error("Failed to delete user", "error", err)
		return nil, err
	}

	return &pb.Success{Message: "User deleted"}, nil
}
