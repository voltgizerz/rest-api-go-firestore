package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
)

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
)

const (
	ErrAuthorizationMissing = "Authorization token missing"
	ErrInvalidToken         = "Invalid token"
	ErrInvalidTokenClaims   = "Invalid token claims"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			api.JSONResponse(c, http.StatusUnauthorized, ErrAuthorizationMissing, nil)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			api.JSONResponse(c, http.StatusUnauthorized, ErrInvalidToken, nil)
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			api.JSONResponse(c, http.StatusUnauthorized, ErrInvalidToken, nil)
			c.Abort()
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			api.JSONResponse(c, http.StatusUnauthorized, ErrInvalidTokenClaims, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
