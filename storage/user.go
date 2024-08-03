package storage

import pb "authentication-service/generated/user"

type UserStorage interface {
	CreateUser(in *pb.CreateUserRequest) (*pb.UserResponse, error)
	GetUserProfile(in *pb.UserId) (*pb.UserResponse, error)
	GetAllUsers(in *pb.FilterRequest) (*pb.UsersResponse, error)
	UpdateUserProfile(in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error)
	GetPassword(in *pb.ChangePasswordRequest) (string, error)
	ChangePassword(in *pb.ChangePasswordRequest) error
	DeleteUser(in *pb.UserId) error
}
