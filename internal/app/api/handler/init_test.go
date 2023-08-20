package handler

import (
	"testing"

	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interactor"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/mocks"
	gomock "go.uber.org/mock/gomock"
)

type TestMocks struct {
	APIHandler         *APIHandler
	MockUserInteractor *mocks.MockUserUsecaseInterface
}

func InitMockTest(t *testing.T) *TestMocks {
	ctrl := gomock.NewController(t)

	mockUserInteractor := mocks.NewMockUserUsecaseInterface(ctrl)

	interactor := interactor.APInteractor{
		UserInteractor: mockUserInteractor,
	}

	apiHandler := &APIHandler{
		APInteractor: interactor,
	}

	return &TestMocks{
		MockUserInteractor: mockUserInteractor,
		APIHandler:         apiHandler,
	}
}
