package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api/middleware"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
)

const (
	defaultPort = "8080"
)

type Router struct {
	GinEngine  *gin.Engine
	APIHandler interfaces.APIHandlerInterface
}

func NewRouter(apiHandler interfaces.APIHandlerInterface) *Router {
	r := &Router{
		GinEngine:  gin.Default(),
		APIHandler: apiHandler,
	}

	r.GinEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	r.GinEngine.GET("/api/token", r.APIHandler.GenerateToken)

	r.userRouter()

	return r
}

func RunAPIServer(r *Router) {
	// Read the port number from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Run the server on the specified port
	r.GinEngine.Run(":" + port)
}

func (r *Router) userRouter() {
	apiGroup := r.GinEngine.Group("/api")

	// Use the JWTMiddleware only for the protected routes
	apiGroup.Use(middleware.JWTMiddleware())


	// Define the user-related routes within the apiGroup
	apiGroup.GET("/users/:docRefID", r.APIHandler.GetUserByID)
	apiGroup.GET("/users", r.APIHandler.GetAllUsers)
	apiGroup.POST("/users", r.APIHandler.InsertUser)
	apiGroup.DELETE("/users/:docRefID", r.APIHandler.DeleteUser)
}
