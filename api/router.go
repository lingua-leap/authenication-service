package api

import (
	_ "authentication-service/api/docs"
	"authentication-service/api/handler"
	"authentication-service/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

// @title Authenfication service
// @version 1.0
// @description server for siginIn or signUp
// @host localhost:8081
// @BasePath /auth
// @schemes http
func (a *api) InitRoutes() {
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := a.router.Group("/auth")
	{
		auth.POST("/register", a.handler.AuthHandler().RegisterHandler)
		auth.POST("/login", a.handler.AuthHandler().LoginHandler)
		auth.POST("/verify-token", a.handler.AuthHandler().VerifyTokenHandler)
		auth.POST("/send/token", a.handler.AuthHandler().ForgotPasswordHandler)
		auth.POST("/reset-password", a.handler.AuthHandler().ResetPasswordHandler)
	}
}
