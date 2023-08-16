package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
)

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
)

// Image this is a database
var MapClientSecret = map[string]string{
	"sample": "this-is-secret",
}

func (a *APIHandler) GenerateToken(c *gin.Context) {
	clientID := c.Query("client_id")
	clientSecret := c.Query("client_secret")

	// Validate client credentials
	secret, ok := MapClientSecret[clientID]
	if !ok || secret != clientSecret {
		api.JSONResponse(c, http.StatusUnauthorized, "Invalid client credentials", nil)
		return
	}

	token, err := generateToken(clientID)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	api.JSONResponse(c, http.StatusOK, "Token generated successfully", map[string]string{
		"token": token,
	})
}

func generateToken(clientID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
