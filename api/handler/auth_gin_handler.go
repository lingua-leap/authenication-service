package handler

import (
	"authentication-service/models"
	"authentication-service/service"
	"authentication-service/service/token"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// "github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

type AuthenticationGINHandler struct {
	authService service.AuthService
	userService service.UserService
	redis       *redis.Client
}

func NewAuthenticationHandler(authService service.AuthService, userService service.UserService, reConn *redis.Client) AuthenticationHandler {
	return &AuthenticationGINHandler{authService, userService, reConn}
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
		c.JSON(400, gin.H{"error-1": err.Error()})
		return
	}

	user, err := a.authService.Register(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error-2": err.Error()})
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
		c.JSON(http.StatusBadRequest, models.Errors{Error: err.Error()})
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
		c.JSON(http.StatusBadRequest, models.Errors{Error: err.Error()})
		return
	}

	value, check := c.Get("claims")
	if !check {
		c.JSON(http.StatusBadRequest, models.Errors{Error: err.Error()})
		return
	}

	claims, ok := value.(*token.Claims)
	if !ok {
		c.JSON(http.StatusBadRequest, models.Errors{Error: err.Error()})
		return
	}

	res, err := a.authService.RefreshToken(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Errors{Error: err.Error()})
		return
	}

	res.RefreshToken = refreshToken
	res.ExpiresAt = time.Now().Add(time.Minute * 10)

	c.SetCookie("access_token", res.AccessToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, res)
}

// @Summary Forgot Password Request
// @Description Sending a code to the Email
// @Tags Auth
// @Accept json
// @Produce json
// @Param Email body models.Email true "send a secret token to the email"
// @Failure 400 {object} models.Errors
// @Failure 404 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /forgot-password [post]
func (a *AuthenticationGINHandler) ForgotPasswordHandler(c *gin.Context) {

	var email models.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	valideEmail, err := service.ValidateEmail(email.Email)
	if valideEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := a.authService.ResetTokenToEmail(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go func() {
		subject := "Password Recovery"
		secretToken, err := token.GenerateToken(valideEmail, "", "")
		if err != nil {
			log.Println("Failed to generate token", err)
			return
		}

		message := "Secure password recovery"

		msg, err := service.SendEmail([]string{valideEmail}, subject, secretToken, message)
		if err != nil {
			log.Println("Failed to send email", "error", err)
			return
		}

		a.redis.Set(context.Background(), valideEmail, secretToken, time.Minute*10)

		log.Printf("Email sent: %s", msg)

	}()

	c.JSON(http.StatusOK, gin.H{"message": "Password reset instructions sent to your email"})
}

// @Summary Verify the password reset instructions
// @Description Verify the password reset instructions
// @Tags Auth
// @Accept json
// @Produce json
// @Param UpdatePassword body models.UpdatePassword true "password reset instructions"
// @Failure 400 {object} models.Errors
// @Failure 404 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /reset-password [post]
func (a *AuthenticationGINHandler) ResetPasswordHandler(c *gin.Context) {
	var updatePassword models.UpdatePassword
	if err := c.ShouldBindJSON(&updatePassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_1": err.Error()})
		return
	}

	_, err := a.authService.RecoveryPassword(updatePassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
