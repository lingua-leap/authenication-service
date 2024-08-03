package main

import (
	"authentication-service/cmd/server"
	configs "authentication-service/config"
	"authentication-service/logs"
	"authentication-service/storage/postgres"
	"log"
)

func main() {
	logger := logs.InitLogger()

	config := configs.Load()

	db, err := postgres.ConnectPostgres(config)
	if err != nil {
		logger.Error("Error connecting to database", "error", err)
		log.Fatalln(err)
	}
	defer db.Close()

	var wait chan struct{}

	go server.GrpcService(logger, db)
	go server.GinService(logger, db)

	<-wait
}
