package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	auth := &Auth{
		SecretKey: "your_secret_key",
	}

	clientID := "test_client"
	token, err := auth.generateToken(clientID)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}