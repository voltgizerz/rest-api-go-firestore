package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
)

func TestGetUserByID_Success(t *testing.T) {
	mock := InitTest(t)

	mock.MockUserInteractor.EXPECT().GetUserDataByDocRefID(gomock.Any(), gomock.Any()).Return(&entity.User{}, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = []gin.Param{{Key: "docRefID", Value: "123"}}

	mock.APIHandler.GetUserByID(c)

	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, c.Writer.Status())
	}
}
