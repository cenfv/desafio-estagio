package controllers

import (
	"net/http"

	"github.com/desafio-estagio/src/services"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string
	Password string
}

// @Summary Login
// @Description Authenticate user credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200
// @Failure 401
// @Router /login [post]
func Login(context *gin.Context) {
	var credentials Credentials

	err := context.Bind(&credentials)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Authenticate(credentials.Email, credentials.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization is required"})
			c.Abort()
			return
		}

		email, err := services.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Next()
	}
}
