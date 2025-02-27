package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
)

type Auth struct {
	SecretKey       string
	MapClientSecret map[string]string
}

func NewAuth(cfgAuth config.Auth) interfaces.AuthInterface {
	mapClientSecret := map[string]string{ // imagine this data got from database
		"sample": "BiquzG0JVY3pWPrh8xiVPkbNXyx20Gmn",
	}
	auth := &Auth{
		SecretKey:       cfgAuth.JWTSecretKey,
		MapClientSecret: mapClientSecret,
	}

	return auth
}

func (a *Auth) GenerateToken(c *gin.Context) {
	clientID := c.Query("client_id")
	clientSecret := c.Query("client_secret")

	// Validate client credentials
	secret, ok := a.MapClientSecret[clientID]
	if !ok || secret != clientSecret {
		api.JSONResponse(c, http.StatusUnauthorized, "Invalid client credentials", nil)
		return
	}

	token, err := a.generateToken(clientID)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	api.JSONResponse(c, http.StatusOK, "Token generated successfully", map[string]string{
		"type":  "Bearer",
		"token": token,
	})
}

func (a *Auth) generateToken(clientID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
