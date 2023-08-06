package main

import (
	"context"
	"log"

	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/repository"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/usecase"
)

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	userData, err := userUsecase.GetAllUserData(context.Background())
	log.Println(userData, err)
}
