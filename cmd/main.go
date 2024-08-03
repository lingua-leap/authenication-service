package main

import (
	configs "authentication-service/config"
	postgres2 "authentication-service/storage/postgres"
	"authentication-service/storagev2/postgres"
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

	authSQLStorage := postgres2.NewAuthenticationSQLStorage(db)
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
