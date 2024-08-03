package api

import (
	"authentication-service/api/handler"
	"authentication-service/config"

	"github.com/gin-gonic/gin"
)

type ApiRouter interface {
	InitRoutes()
	Run() error
}

type api struct {
	router  *gin.Engine
	handler handler.MainHandler
}

func NewAPI(handler handler.MainHandler) ApiRouter {
	router := gin.Default()
	return &api{router: router, handler: handler}
}

func (api *api) Run() error {
	return api.router.Run(
		config.Load().GIN_SERVER_PORT)
}

func (a *api) InitRoutes() {
	auth := a.router.Group("/auth")
	{
		auth.POST("/register", a.handler.AuthHandler().RegisterHandler)
		auth.POST("/login", a.handler.AuthHandler().LoginHandler)
		auth.POST("/verify-token", a.handler.AuthHandler().VerifyTokenHandler)
		auth.POST("/send/token", a.handler.AuthHandler().ForgotPasswordHandler)
		auth.POST("/reset-password", a.handler.AuthHandler().ResetPasswordHandler)
	}
}
