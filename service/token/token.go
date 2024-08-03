package token

import (
	"authentication-service/config"
	"authentication-service/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToken(in models.User) (string, error) {
	claims := Claims{
		Username: in.Username,
		ID:       in.ID,
		Role:     "user",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Load().SECRET_KEY_ACCESS))
}

func GenerateRefreshToken(in models.User) (string, error) {
	claims := Claims{
		Username: in.Username,
		ID:       in.ID,
		Role:     "user",
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
