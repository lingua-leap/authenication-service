package handler

import (
	"authentication-service/models"
	"authentication-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationGINHandler struct {
	authService service.AuthService
}

func NewAuthenticationHandler(authService service.AuthService) AuthenticationHandler {
	return &AuthenticationGINHandler{authService}
}

func (a *AuthenticationGINHandler) RegisterHandler(c *gin.Context) {
	var createUser models.CreateUser

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authService.Register(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *AuthenticationGINHandler) LoginHandler(c *gin.Context) {
	var login models.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := a.authService.Login(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
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
