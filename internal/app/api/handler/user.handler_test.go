package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	gomock "go.uber.org/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	testCases := []struct {
		name          string
		mockUser      *entity.User
		mockErr       error
		paramDocRefID string
		expectedCode  int
	}{
		{
			name:          "Success",
			mockUser:      &entity.User{},
			mockErr:       nil,
			paramDocRefID: "123",
			expectedCode:  http.StatusOK,
		},
		{
			name:          "Error_UserNotFound",
			mockUser:      nil,
			mockErr:       errors.New("some errors"),
			paramDocRefID: "456",
			expectedCode:  http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := InitMockTest(t)

			m.MockUserInteractor.EXPECT().
				GetUserDataByDocRefID(gomock.Any(), gomock.Any()).
				Return(tc.mockUser, tc.mockErr)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Params = []gin.Param{{Key: "docRefID", Value: tc.paramDocRefID}}

			m.APIHandler.GetUserByID(c)

			if c.Writer.Status() != tc.expectedCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedCode, c.Writer.Status())
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	testCases := []struct {
		name         string
		mockUsers    []entity.User
		mockErr      error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			mockUsers:    []entity.User{},
			mockErr:      nil,
			expectedCode: http.StatusOK,
			expectedBody: "Users data retrieved successfully",
		},
		{
			name:         "Error",
			mockUsers:    nil,
			mockErr:      errors.New("some errors"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed retrieved data users",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := InitMockTest(t)

			m.MockUserInteractor.EXPECT().
				GetAllUserData(gomock.Any()).
				Return(tc.mockUsers, tc.mockErr)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			m.APIHandler.GetAllUsers(c)

			if c.Writer.Status() != tc.expectedCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedCode, c.Writer.Status())
			}
		})
	}
}
