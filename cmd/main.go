package main

import (
	configs "authentication-service/config"
	pb "authentication-service/generated/user"
	"authentication-service/logs"
	"authentication-service/service"
	"authentication-service/storage"
	"authentication-service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
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

	usr := service.NewUserService(logger, storage.NewMainStorage(db))
	serv := grpc.NewServer()

	pb.RegisterUserServiceServer(serv, usr)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error("Error listening on port 8080", "error", err)
		log.Fatalln(err)
	}
	log.Println("Listening on port 8080")
	log.Fatal(serv.Serve(listener))
}
