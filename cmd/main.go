package main

import (
	configs "authentication-service/config"
	"authentication-service/storage/postgres"
	"log"
)

func main() {
	dbConfig, err := configs.GetDatabaseCongig(".")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	db, err := postgres.ConnectPostgres(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	authSQLStorage := postgres.NewAuthenticationSQLStorage(db)
	userSQLStorage := postgres.NewUserManagementStorage(db)

	sqlStorage := postgres.NewMainSQLStorage(db, *userSQLStorage, *authSQLStorage)

	defer sqlStorage.CloseDB()
	// user := models.CreateUser{
	// 	Email:    "test@example.com",
	// 	Username: "test",
	// 	Password: "test123",
	// }

	cUser, err := sqlStorage.AuthenticationStorage().Login("test@example.com", "test123")
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}
	log.Println(cUser)
}
