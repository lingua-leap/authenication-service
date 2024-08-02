package service

import (
	pb "authentication-service/generated/user"
	"authentication-service/storage"
	"context"
)

type UserService struct {
	storage storage.UserManagementStorage
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage storage.UserManagementStorage) *UserService {
	return &UserService{storage: storage}
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := s.storage.CreateUser(req)
	return res, err
}
func (s *UserService) GetUserProfile(ctx context.Context, req *pb.UserId) (*pb.UserProfile, error) {
	res, err := s.storage.GetUserProfile(req)
	return res, err
}
func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res, err := s.storage.UpdateUserProfile(req)
	return res, err
}
func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Success, error) {
	res, err := s.storage.ChangePassword(req)
	return res, err
}
