package handler

import (
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interactor"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
)

type APIHandler struct {
	APInteractor interactor.APInteractor
}

func NewAPIHandler(interactor interactor.APInteractor) interfaces.APIHandlerInterface {
	apiHandler := &APIHandler{
		APInteractor: interactor,
	}

	return apiHandler
}
