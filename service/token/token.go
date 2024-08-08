package token

import (
	"authentication-service/config"
	"authentication-service/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToken(in models.User) (string, error) {
	role := "user"
	if in.Username == "admin" {
		role = "admin"
	}

	claims := Claims{
		Username: in.Username,
		ID:       in.ID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Load().SECRET_KEY_ACCESS))
}

func GenerateRefreshToken(in models.User) (string, error) {
	role := "user"
	if in.Username == "admin" {
		role = "admin"
	}

	claims := Claims{
		Username: in.Username,
		ID:       in.ID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Load().SECRET_KEY_ACCESS))
}

func ExtractAccessClaims(token string) (*Claims, error) {
	tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SECRET_KEY_ACCESS), nil
	})
	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := tk.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractRefreshClaims(token string) (*Claims, error) {
	tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SECRET_KEY_REFRESH), nil
	})
	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := tk.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

type GClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(email, id, username string) (string, error) {
	claims := GClaims{
		Email:    email,
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := config.Load()
	strToken, err := token.SignedString([]byte(config.SECRET_KEY_ACCESS))
	return strToken, err
}

func ExtractClaims(token string) (*GClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &GClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SECRET_KEY_ACCESS), nil
	})
	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, tk.Claims.Valid()
	}

	claims, ok := tk.Claims.(*GClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
