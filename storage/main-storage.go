package storage

import (
	"authentication-service/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type MainStorage interface {
	NewUserStorage() UserStorage
	NewAuthStorage() AuthStorage
}

type MainStorageImpl struct {
	db *sqlx.DB
}

func NewMainStorage(db *sqlx.DB) MainStorage {
	return &MainStorageImpl{db}
}

func (m *MainStorageImpl) NewUserStorage() UserStorage {

	return postgres.NewUserRepo(m.db)
}

func (m *MainStorageImpl) NewAuthStorage() AuthStorage {
	return postgres.NewAuthStorage(m.db)
}
