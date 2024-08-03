package storage

import (
	"authentication-service/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type MainStorage interface {
	NewUserStorage() UserStorage
}

type MainStorageImpl struct {
	db *sqlx.DB
}

func (m *MainStorageImpl) NewUserStorage() UserStorage {
	return postgres.NewUserRepo(m.db)
}

func NewMainStorage(db *sqlx.DB) MainStorage {
	return &MainStorageImpl{db}
}
