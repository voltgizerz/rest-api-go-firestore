package interfaces

import "github.com/gin-gonic/gin"

type AuthInterface interface {
	GenerateToken(c *gin.Context)
}
