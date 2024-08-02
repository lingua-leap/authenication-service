package handler

import "github.com/gin-gonic/gin"

type AuthenticationHandler interface {
	RegisterHandler(*gin.Context)
	LoginHandler(*gin.Context)
	VerifyTokenHandler(*gin.Context)
	LogoutHandler(*gin.Context)
	ForgotPasswordHandler(*gin.Context)
	ResetPasswordHandler(*gin.Context)
}
