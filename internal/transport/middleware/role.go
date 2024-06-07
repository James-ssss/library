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

func RoleMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid (no role)")
		return
	}

	splitted := strings.Split(auth, " ")

	role, err := parseTokenRole(splitted[0])

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid role")
	}

	c.Set("role", role)

	if role != "admin" {
		return
	}
	c.Next()
}

func parseTokenRole(token string) (string, error) {
	// at - access token
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

	if !ok {
		return "", err
	}

	return claims.Role, nil
}
