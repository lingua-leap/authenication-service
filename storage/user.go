package storage

import (
	pb "authentication-service/generated/user"
	"authentication-service/help"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"time"
)

type UserManagementStorage interface {
	CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUserProfile(in *pb.UserId) (*pb.UserProfile, error)
	UpdateUserProfile(in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error)
	ChangePassword(in *pb.ChangePasswordRequest) (*pb.Success, error)
}

type UserRepo struct {
	Log slog.Logger
	db  *sqlx.DB
}

func NewUserRepo(log slog.Logger, db *sqlx.DB) UserManagementStorage {
	return &UserRepo{log, db}
}

func (u *UserRepo) CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := &pb.CreateUserResponse{}

	err := u.db.QueryRow(`INSERT INTO users
    (username, email, password_hash, full_name, native_language)
    VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		in.Username, in.Email, in.Password, in.FullName, in.NativeLanguage).Scan(res.Id)
	if err != nil {
		u.Log.Error("Error inserting user into database", "error", err)
		return nil, err
	}

	res.Username = in.Username
	res.Email = in.Email
	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage
	res.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	return res, nil
}

func (u *UserRepo) GetUserProfile(in *pb.UserId) (*pb.UserProfile, error) {
	res := &pb.UserProfile{}

	err := u.db.Get(res, "SELECT * FROM users WHERE id = $1", in.Id)
	if err != nil {
		u.Log.Error("Error getting user", "error", err)
		return nil, err
	}

	return res, nil
}

func (u *UserRepo) UpdateUserProfile(in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res := &pb.UpdateUserPResponse{}

	err := u.db.QueryRow(`UPDATE users set username=$1, full_name=$2, native_language=$3
             WHERE id = $4 RETURNING email`,
		in.Username, in.FullName, in.NativeLanguage, in.Id).Scan(res.Email)
	if err != nil {
		u.Log.Error("Error updating user", "error", err)
		return nil, err
	}

	res.Username = in.Username
	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage
	res.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return res, nil
}

func (u *UserRepo) ChangePassword(in *pb.ChangePasswordRequest) (*pb.Success, error) {
	var res = &pb.Success{}
	var hashPassword string

	err := u.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", in.Id).Scan(&hashPassword)
	if err != nil {
		u.Log.Error("Error getting user", "error", err)
		return nil, err
	}

	check := help.VerifyPassword(hashPassword, in.OldPassword)
	if !check {
		u.Log.Info("Invalid old password", "old_password", in.OldPassword)
		res.Message = "Invalid old password"
		return res, nil
	}

	hash, err := help.HashPassword(in.NewPassword)
	if err != nil {
		u.Log.Error("Error hashing password", "error", err)
		return nil, err
	}

	_, err = u.db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2",
		hash, in.Id)
	if err != nil {
		u.Log.Error("Error updating user", "error", err)
		return nil, err
	}

	return res, nil
}
