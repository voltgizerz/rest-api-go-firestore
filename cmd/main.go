package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/database"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api/auth"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api/handler"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/api/router"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interactor"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/repository"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/usecase"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/env"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

func main() {
	env.LoadENV()

	// initialize config
	cfg := config.NewConfig()

	// initialize database
	db := database.InitDB(cfg)

	userRepo := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize a channel to listen for interrupt signals (e.g., Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	interactorAPI := interactor.APInteractor{
		UserInteractor: userUsecase,
	}

	auth := auth.NewAuth()
	apiHandler := handler.NewAPIHandler(interactorAPI)

	dataRouter := router.Router{
		GinEngine:  gin.Default(),
		Auth:       auth,
		APIHandler: apiHandler,
	}

	r := router.NewRouter(dataRouter)
	go router.RunAPIServer(r) // Start the web server in a goroutine

	// Wait for the interrupt signal
	<-quit
	logger.Log.Warn("Shutting down the server...")
	logger.Log.Info("Server gracefully stopped.")
}
