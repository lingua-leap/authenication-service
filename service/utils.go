package service

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPasswordHash(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil

}

func ValidateEmail(email string) (string, error) {
	var err error
	_, err = regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func ValidatePassword(password string) (string, error) {
	var err error
	_, err = regexp.MatchString(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`, password)
	if err != nil {
		return "", err
	}
	return password, nil
}
