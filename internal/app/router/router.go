package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (
	defaultPort = "8080"
)

func initRouter(r *gin.Engine) *gin.Engine {
	// Define your routes here
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	r.GET("/users", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	return r
}

func StartWebServer() {
	r := gin.Default()
	r = initRouter(r)

	// Read the port number from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Run the server on the specified port
	r.Run(":" + port)
}
