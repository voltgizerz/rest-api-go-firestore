package main

import (
	"context"

	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/repository"
)

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)

	_ = userRepo.InsertUserData(context.Background())
	_ = userRepo.GetUserData(context.Background())
}
