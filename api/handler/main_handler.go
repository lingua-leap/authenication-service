package handler

import "authentication-service/service"

type MainHandler interface {
	AuthHandler() AuthenticationHandler
}

type MainHandlerImpl struct {
	service service.AuthService
}

func NewMainHandler(service service.AuthService) MainHandler {
	return &MainHandlerImpl{service: service}
}

func (m *MainHandlerImpl) AuthHandler() AuthenticationHandler {
	return NewAuthenticationHandler(m.service)
}
