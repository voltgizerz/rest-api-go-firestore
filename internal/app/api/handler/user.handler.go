package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
)

func (a *APIHandler) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	docRefID := c.Param("docRefID")

	user, err := a.APInteractor.UserInteractor.GetUserDataByDocRefID(ctx, docRefID)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed retrieve data "+docRefID, nil)
		return
	}

	api.JSONResponse(c, http.StatusOK, "User data retrieved successfully", user)
}

func (a *APIHandler) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := a.APInteractor.UserInteractor.GetAllUserData(ctx)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed retrieved data users", nil)
		return
	}

	api.JSONResponse(c, http.StatusOK, "Users data retrieved successfully", users)
}

func (a *APIHandler) InsertUser(c *gin.Context) {
	ctx := c.Request.Context()

	var user entity.User
	// Uncomment and use c.ShouldBindJSON(&user) to bind JSON data to user struct

	docRefID, err := a.APInteractor.UserInteractor.InsertUserData(ctx, user)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed insert data user", nil)
		return
	}

	api.JSONResponse(c, http.StatusCreated, "User data inserted successfully", map[string]string{
		"DocRefID": docRefID,
	})
}

func (a *APIHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	docRefID := c.Param("docRefID")

	success, err := a.APInteractor.UserInteractor.DeleteUserDataByDocRefID(ctx, docRefID)
	if err != nil {
		api.JSONResponse(c, http.StatusInternalServerError, "Failed delete data user", nil)
		return
	}

	if success {
		api.JSONResponse(c, http.StatusOK, "User data deleted successfully", nil)
	} else {
		api.JSONResponse(c, http.StatusNotFound, "User data not found", nil)
	}
}
