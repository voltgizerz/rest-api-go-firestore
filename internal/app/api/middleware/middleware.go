package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
)

const (
	ErrAuthorizationMissing = "Authorization token missing"
	ErrInvalidToken         = "Invalid token"
	ErrInvalidTokenClaims   = "Invalid token claims"
)

func JWTMiddleware(cfgAuth config.Auth) gin.HandlerFunc {
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
			return []byte(cfgAuth.JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			api.JSONResponse(c, http.StatusUnauthorized, ErrInvalidToken, nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			api.JSONResponse(c, http.StatusUnauthorized, ErrInvalidTokenClaims, nil)
			c.Abort()
			return
		}

		// Set data in the context
		c.Set("userID", claims["client_id"])

		c.Next()
	}
}
