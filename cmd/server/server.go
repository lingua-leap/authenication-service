package server

import (
	"authentication-service/api"
	"authentication-service/api/handler"
	"authentication-service/config"
	pb "authentication-service/generated/user"
	"authentication-service/service"
	"authentication-service/storage"
	redisr "authentication-service/storage/redis"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func GrpcService(logger *slog.Logger, db *sqlx.DB) {
	usr := service.NewUserService(logger, storage.NewMainStorage(db))
	serv := grpc.NewServer()

	pb.RegisterUserServiceServer(serv, usr)

	listener, err := net.Listen("tcp", config.Load().GRPC_SERVER_PORT)
	if err != nil {
		str := fmt.Sprintf("failed to listen %v", config.Load().GRPC_SERVER_PORT)
		logger.Error(str, "error", err)
		log.Fatalln(err)
	}

	str := fmt.Sprintf("gRPC server listening on port %v", config.Load().GRPC_SERVER_PORT)
	logger.Info(str, "port", listener.Addr().(*net.TCPAddr).Port)
	log.Printf("\nServer listening on port %s \n", config.Load().GRPC_SERVER_PORT)
	log.Fatal(serv.Serve(listener))
}

func GinService(logger *slog.Logger, db *sqlx.DB) {
	mainStorage := storage.NewMainStorage(db)
	authService := service.NewAuthService(logger, mainStorage)
	redisConn := redisr.ConnectRedis()
	userService := service.NewUserService(logger, mainStorage)
	handler1 := handler.NewMainHandler(authService, redisConn, *userService)

	Api := api.NewAPI(handler1)
	Api.InitRoutes()

	log.Fatalln(Api.Run())
}
