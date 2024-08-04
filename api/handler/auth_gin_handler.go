package handler

import (
	"authentication-service/models"
	"authentication-service/service"
	"authentication-service/service/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthenticationGINHandler struct {
	authService service.AuthService
}

func NewAuthenticationHandler(authService service.AuthService) AuthenticationHandler {
	return &AuthenticationGINHandler{authService}
}

// @Summary SigUp user
// @Description Create user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.CreateUser true "create user"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.Errors
// @Failure 404 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /register [post]
func (a *AuthenticationGINHandler) RegisterHandler(c *gin.Context) {
	var createUser models.CreateUser

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authService.Register(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary Login user
// @Description Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Login body models.LoginRequest true "login user"
// @Failure 400 {object} models.Errors
// @Failure 404 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /login [post]
func (a *AuthenticationGINHandler) LoginHandler(c *gin.Context) {
	var login models.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := a.authService.Login(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
	}

	c.JSON(http.StatusOK, res)

	c.SetCookie("access_token", res.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "/", "", false, true)
}

// @Summary Get Access Token user
// @Description Get Access token by refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param Login body models.RefreshResponse true "get access token"
// @Failure 400 {object} models.Errors
// @Failure 404 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /verify-token [post]
func (a *AuthenticationGINHandler) VerifyTokenHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
		return
	}

	value, check := c.Get("claims")
	if !check {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
		return
	}

	claims, ok := value.(*token.Claims)
	if !ok {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
		return
	}

	res, err := a.authService.RefreshToken(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Errors{err.Error()})
		return
	}

	res.RefreshToken = refreshToken
	res.ExpiresAt = time.Now().Add(time.Minute * 10)

	c.SetCookie("access_token", res.AccessToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, res)
}

func (a *AuthenticationGINHandler) ForgotPasswordHandler(c *gin.Context) {
	var email models.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	mesaage, err := a.authService.ResetTokenToEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mesaage)
}

func (a *AuthenticationGINHandler) ResetPasswordHandler(c *gin.Context) {
	// Your handler logic here
}
