package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/repository"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/router"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/usecase"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
)

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	userData, err := userUsecase.GetAllUserData(context.Background())
	if err != nil {
		log.Fatal("Failed to fetch user data:", err)
	}
	log.Println(userData)

	// Initialize a channel to listen for interrupt signals (e.g., Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start the web server in a goroutine
	go router.StartWebServer()

	// Wait for the interrupt signal
	<-quit
	logger.Log.Warn("Shutting down the server...")
	logger.Log.Info("Server gracefully stopped.")
}
