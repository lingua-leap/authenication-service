package api

import (
	"authentication-service/api/handler"

	"github.com/gin-gonic/gin"
)

type api struct {
	router  *gin.Engine
	handler handler.MainHandler
}

func (a *api) InitRoutes() {
	auth := a.router.Group("/auth")
	{
		auth.POST("/register", a.handler.AuthenticationHandler().RegisterHandler)
		auth.POST("/login", a.handler.AuthenticationHandler().LoginHandler)
		auth.GET("/verify-token", a.handler.AuthenticationHandler().VerifyTokenHandler)
		auth.POST("/logout", a.handler.AuthenticationHandler().LogoutHandler)
		auth.POST("/forgot-password", a.handler.AuthenticationHandler().ForgotPasswordHandler)
		auth.POST("/reset-password", a.handler.AuthenticationHandler().ResetPasswordHandler)
	}
}
