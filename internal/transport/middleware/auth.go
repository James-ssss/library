package middleware

import (
	"crud/internal/core/model"
	"crud/internal/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid (no token)"})
		return
	}

	splitted := strings.Split(auth, " ")
	if len(splitted) != 2 || splitted[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token format"})
		return
	}

	login, err := parseToken(splitted[1])

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	c.Set("user", login)
	c.Next()
}

func parseToken(token string) (string, error) {
	at, err := jwt.ParseWithClaims(token, &service.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid method")
			}
			return []byte(model.SignInKey), nil
		})

	if err != nil {
		return "", err
	}

	claims, ok := at.Claims.(*service.TokenClaims)

	if !ok || !at.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	return claims.Login, nil
}
