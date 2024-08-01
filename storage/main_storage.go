package storage

type MainStorage interface {
	UserManagementStorage() UserManagementStorage
	AuthenticationStorage() AuthenticationStorage

	CloseDB() error
}
