package postgres

import (
	pb "authentication-service/generated/user"
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

func (u *UserManagementSQLStorage) CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := &pb.CreateUserResponse{}

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

func (u *UserManagementSQLStorage) GetUserProfile(in *pb.UserId) (*pb.UserProfile, error) {
	res := &pb.UserProfile{}

	err := u.db.Get(res, "SELECT * FROM users WHERE id = $1", in.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserManagementSQLStorage) UpdateUserProfile(in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res := &pb.UpdateUserPResponse{}

	err := u.db.QueryRow(`UPDATE users set username=$1, full_name=$2, native_language=$3
             WHERE id = $4 RETURNING email`,
		in.Username, in.FullName, in.NativeLanguage, in.Id).Scan(res.Email)
	if err != nil {
		return nil, err
	}

	res.Username = in.Username
	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage
	res.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return res, nil
}

func (u *UserManagementSQLStorage) ChangePassword(in *pb.ChangePasswordRequest) (*pb.Success, error) {
	var res = &pb.Success{}
	var hashPassword string

	err := u.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", in.Id).Scan(&hashPassword)
	if err != nil {
		return nil, err
	}

	check := help.VerifyPassword(hashPassword, in.OldPassword)
	if !check {
		res.Message = "Invalid old password"
		return res, nil
	}

	hash, err := help.HashPassword(in.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = u.db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2",
		hash, in.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
