package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
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

	r.GinEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

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
	r.GinEngine.GET("api/users/:docRefID", func(c *gin.Context) {
		ctx := c.Request.Context()

		docRefID := c.Param("docRefID")

		user, err := r.APInteractor.UserInteractor.GetUserDataByDocRefID(ctx, docRefID)
		if err != nil {
			api.JSONResponse(c, http.StatusInternalServerError, "Failed retrieve data "+docRefID, nil)
			return
		}

		api.JSONResponse(c, http.StatusOK, "User data retrieved successfully", user)
	})

	r.GinEngine.GET("/api/users", func(c *gin.Context) {
		ctx := c.Request.Context()

		users, err := r.APInteractor.UserInteractor.GetAllUserData(ctx)
		if err != nil {
			api.JSONResponse(c, http.StatusInternalServerError, "Failed retrieved data users", nil)
			return
		}

		api.JSONResponse(c, http.StatusOK, "Users data retrieved successfully", users)
	})

	r.GinEngine.POST("api/users", func(c *gin.Context) {
		ctx := c.Request.Context()

		var user entity.User
		// * currently data user filled by faker
		// if err := c.ShouldBindJSON(&user); err != nil {
		// 	api.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		// 	return
		// }

		docRefID, err := r.APInteractor.UserInteractor.InsertUserData(ctx, user)
		if err != nil {
			api.JSONResponse(c, http.StatusInternalServerError, "Failed insert data user", nil)
			return
		}

		api.JSONResponse(c, http.StatusCreated, "User data inserted successfully", map[string]string{
			"DocRefID": docRefID,
		})
	})

	r.GinEngine.DELETE("api/users/:docRefID", func(c *gin.Context) {
		ctx := c.Request.Context()

		docRefID := c.Param("docRefID")

		success, err := r.APInteractor.UserInteractor.DeleteUserDataByDocRefID(ctx, docRefID)
		if err != nil {
			api.JSONResponse(c, http.StatusInternalServerError, "Failed delete data user", nil)
			return
		}

		if success {
			api.JSONResponse(c, http.StatusOK, "User data deleted successfully", nil)
		} else {
			api.JSONResponse(c, http.StatusNotFound, "User data not found", nil)
		}
	})
}
