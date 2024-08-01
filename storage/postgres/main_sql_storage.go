package postgres

import (
	"authentication-service/storage"

	"github.com/jmoiron/sqlx"
)

type MainSQLStorage struct {
	userManagementSQLStorage UserManagementSQLStorage
	authenticationSQLStorage AuthenticationSQLStorage
	db                       *sqlx.DB
}

func NewMainSQLStorage(db *sqlx.DB, user UserManagementSQLStorage, auth AuthenticationSQLStorage) storage.MainStorage {
	return &MainSQLStorage{
		userManagementSQLStorage: user,
		authenticationSQLStorage: auth,
		db:                       db,
	}
}

func (s *MainSQLStorage) CloseDB() error {
	return s.db.Close()
}

func (s *MainSQLStorage) UserManagementStorage() storage.UserManagementStorage {
	return &s.userManagementSQLStorage
}

func (s *MainSQLStorage) AuthenticationStorage() storage.AuthenticationStorage {
	return &s.authenticationSQLStorage
}
