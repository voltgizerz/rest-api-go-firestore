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
	Auth       interfaces.AuthInterface
	APIHandler interfaces.APIHandlerInterface
}

func NewRouter(dataRouter Router) *Router {
	r := &Router{
		GinEngine:  dataRouter.GinEngine,
		Auth:       dataRouter.Auth,
		APIHandler: dataRouter.APIHandler,
	}

	gin.DefaultWriter = logger.Log.Writer()

	r.generalRouter()

	apiBaseRouter := r.GinEngine.Group("/api")
	r.userRouter(apiBaseRouter)

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

	r.GinEngine.GET("/api/token", r.Auth.GenerateToken)
}

func (r *Router) userRouter(baseRouter *gin.RouterGroup) {
	usersRouter := baseRouter.Group("/users").Use(middleware.JWTMiddleware()) // Use the JWTMiddleware only for the protected routes

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
