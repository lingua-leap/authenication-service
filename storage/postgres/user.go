package postgres

import (
	pb "authentication-service/generated/auth"
	"authentication-service/help"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserManagementSQLStorage struct {
	db *sqlx.DB
}

func NewUserManagementStorage(db *sqlx.DB) *UserManagementSQLStorage {
	return &UserManagementSQLStorage{db}
}

func (u *UserManagementSQLStorage) CreateUser(in *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	res := &pb.UserResponse{}

	err := u.db.QueryRow(`INSERT INTO users
    (username, email, password_hash, full_name, native_language)
    VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		in.Username, in.Email, in.Password, in.FullName, in.NativeLanguage).Scan(res.Id)
	if err != nil {
		return nil, err
	}

	res.Username = in.Username
	res.Email = in.Email
	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage
	res.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	return res, nil
}

func (u *UserManagementSQLStorage) GetUserProfile(in *pb.GetUserProfileRequest) (*pb.UserResponse, error) {
	res := &pb.UserResponse{}

	err := u.db.Get(res, "SELECT * FROM users WHERE id = $1", in.UserId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserManagementSQLStorage) UpdateUserProfile(in *pb.UpdateUserProfileRequest) (*pb.UserResponse, error) {
	res := &pb.UserResponse{}

	err := u.db.QueryRow(`UPDATE users set full_name=$1, native_language=$2
             WHERE id = $3 RETURNING email`,
		in.FullName, in.NativeLanguage, in.UserId).Scan(res.Email)
	if err != nil {
		return nil, err
	}

	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage

	return res, nil
}

func (u *UserManagementSQLStorage) ChangePassword(in *pb.ChangeUserPasswordRequest) (*pb.ChangeUserPasswordResponse, error) {
	var res = &pb.ChangeUserPasswordResponse{}
	var hashPassword string

	err := u.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", in.UserId).Scan(&hashPassword)
	if err != nil {
		return nil, err
	}

	check := help.VerifyPassword(hashPassword, in.CurrentPassword)
	if !check {
		res.Message = "Invalid old password"
		return res, nil
	}

	hash, err := help.HashPassword(in.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = u.db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2",
		hash, in.UserId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// func (u *UserManagementSQLStorage) DeleteUser(in *pb.) error
