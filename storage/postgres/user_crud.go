package postgres

import (
	pb "authentication-service/generated/user"
	"authentication-service/storage"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(sqlx *sqlx.DB) storage.UserStorage {
	return &UserRepo{db: sqlx}
}

func (u *UserRepo) CreateUser(in *pb.CreateUserRequest) (*pb.UserResponse, error) {
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

func (u *UserRepo) GetUserProfile(in *pb.UserId) (*pb.UserResponse, error) {
	res := &pb.UserResponse{}

	err := u.db.Get(res, `SELECT * FROM users WHERE id = $1`, in.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserRepo) GetAllUsers(in *pb.FilterRequest) (*pb.UsersResponse, error) {
	res := pb.UsersResponse{}

	query := `SELECT id, username, email, full_name, native_language, created_at
from users where deleted_at = 0 `
	args := []interface{}{}

	if in.Limit == 0 {
		in.Limit = 10
	}
	if in.Offset == 0 {
		in.Offset = 0
	}
	if in.Native != "" {
		query += " AND native_language=$1"
		args = append(args, in.Native)
	}
	if in.Name != "" {
		query += " AND username LIKE $2"
		args = append(args, "%"+in.Name+"%")
	}

	err := u.db.Get(&res, query, args...)
	return &res, err
}

func (u *UserRepo) UpdateUserProfile(in *pb.UpdateUserPRequest) (*pb.UpdateUserPResponse, error) {
	res := &pb.UpdateUserPResponse{}

	err := u.db.QueryRow(`UPDATE users set full_name=$1, native_language=$2
             WHERE id = $3 RETURNING email`,
		in.FullName, in.NativeLanguage, in.Id).Scan(res.Email)
	if err != nil {
		return nil, err
	}

	res.FullName = in.FullName
	res.NativeLanguage = in.NativeLanguage

	return res, nil
}

func (u *UserRepo) GetPassword(in *pb.ChangePasswordRequest) (string, error) {

	var Password string

	err := u.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", in.Id).Scan(&Password)

	return Password, err

}

func (u *UserRepo) ChangePassword(in *pb.ChangePasswordRequest) error {

	_, err := u.db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2",
		in.NewPassword, in.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) DeleteUser(in *pb.UserId) error {

	_, err := u.db.Exec(`UPDATE users 
        SET  deleted_at=DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
        WHERE id=$1 AND deleted_at = 0`, in.Id)

	return err
}
