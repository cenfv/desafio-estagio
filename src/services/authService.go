package services

import (
	"errors"
	"strings"
	"time"

	"github.com/desafio-estagio/database"
	"github.com/desafio-estagio/src/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authenticate(email, password string) (string, error) {
	var player models.Player

	database.DB.Where("email = ?", email).First(&player)

	if player.Email == "" {
		return "", errors.New("invalid email or password")
	}

	if player.Password != password {
		return "", errors.New("invalid email or password")
	}

	token, tokenErr := GenerateToken(email)
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func ValidateToken(tokenString string) (string, error) {

	splitToken := strings.Split(tokenString, "Bearer ")
	tokenString = strings.TrimSpace(splitToken[1])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email")
	}

	return email, nil
}
