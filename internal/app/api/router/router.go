package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api/middleware"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
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

	gin.DefaultWriter = logger.Log.Writer()

	r.generalRouter()
	r.userRouter()

	return r
}

func RunAPIServer(r *Router) {
	gin.SetMode(getGinModeFromEnv())

	// Read the port number from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r.GinEngine.Run(":" + port)
}

func (r *Router) generalRouter() {
	r.GinEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	r.GinEngine.GET("/api/token", r.APIHandler.GenerateToken)
}

func (r *Router) userRouter() {
	baseRouter := r.GinEngine.Group("/api")
	usersRouter := baseRouter.Group("/users")

	// Use the JWTMiddleware only for the protected routes
	usersRouter.Use(middleware.JWTMiddleware())

	// Define the user-related routes within the usersRouter
	usersRouter.GET("/:docRefID", r.APIHandler.GetUserByID)
	usersRouter.GET("/", r.APIHandler.GetAllUsers)
	usersRouter.POST("/", r.APIHandler.InsertUser)
	usersRouter.DELETE("/:docRefID", r.APIHandler.DeleteUser)
	usersRouter.PATCH("/:docRefID", r.APIHandler.UpdateUser)
}

func getGinModeFromEnv() string {
	ginMode := os.Getenv("GIN_MODE")

	switch ginMode {
	case "release":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}
