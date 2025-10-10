package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/constant"
)

func JWTMiddleware(cfgAuth config.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constant.HeaderAuthorization)
		if authHeader == "" {
			api.JSONResponse(c, http.StatusUnauthorized, constant.ErrAuthorizationMissing, nil)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != constant.HeaderAuthType {
			api.JSONResponse(c, http.StatusUnauthorized, constant.ErrInvalidToken, nil)
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfgAuth.JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			api.JSONResponse(c, http.StatusUnauthorized, constant.ErrInvalidToken, nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			api.JSONResponse(c, http.StatusUnauthorized, constant.ErrInvalidTokenClaims, nil)
			c.Abort()
			return
		}

		// Set data in the context
		c.Set("userID", claims["client_id"])

		c.Next()
	}
}
