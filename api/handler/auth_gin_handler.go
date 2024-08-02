package handler

import (
	"authentication-service/models"
	"authentication-service/service"

	"github.com/gin-gonic/gin"
)

type AuthenticationGINHandler struct {
	authService service.AuthService
}

func NewAuthenticationHandler(authService service.AuthService) AuthenticationHandler {
	return &AuthenticationGINHandler{authService}
}

func (a *AuthenticationGINHandler) RegisterHandler(c *gin.Context) {
	var createUser *models.CreateUser

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authService.Register(createUser)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
	// Your handler logic here
}

func (a *AuthenticationGINHandler) LoginHandler(c *gin.Context) {
	// Your handler logic here
}

func (a *AuthenticationGINHandler) VerifyTokenHandler(c *gin.Context) {
	// Your handler logic here
}

func (a *AuthenticationGINHandler) LogoutHandler(c *gin.Context) {
	// Your handler logic here
}

func (a *AuthenticationGINHandler) ForgotPasswordHandler(c *gin.Context) {
	// Your handler logic here
}

func (a *AuthenticationGINHandler) ResetPasswordHandler(c *gin.Context) {
	// Your handler logic here
}
