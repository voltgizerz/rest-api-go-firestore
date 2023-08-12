package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interactor"
)

const (
	defaultPort = "8080"
)

type Router struct {
	GinEngine    *gin.Engine
	APInteractor interactor.APInteractor
}

func NewRouter(interactor interactor.APInteractor) *Router {
	r := &Router{
		GinEngine:    gin.Default(),
		APInteractor: interactor,
	}

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
	r.GinEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	r.GinEngine.GET("/api/users", func(c *gin.Context) {
		ctx := c.Request.Context()

		users, err := r.APInteractor.UserInteractor.GetAllUserData(ctx)
		if err != nil {
			api.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		api.JSONResponse(c, http.StatusOK, "User data retrieved successfully", users)
	})
}
