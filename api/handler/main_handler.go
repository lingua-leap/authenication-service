package handler

import (
	"authentication-service/service"

	"github.com/redis/go-redis/v9"
)

type MainHandler interface {
	AuthHandler() AuthenticationHandler
}

type MainHandlerImpl struct {
	service     service.AuthService
	userService service.UserService
	redisConn   *redis.Client
}

func NewMainHandler(service service.AuthService, redisConn *redis.Client, userService service.UserService) MainHandler {
	return &MainHandlerImpl{service, userService, redisConn}
}

func (m *MainHandlerImpl) AuthHandler() AuthenticationHandler {
	return NewAuthenticationHandler(m.service, m.userService, m.redisConn)
}
