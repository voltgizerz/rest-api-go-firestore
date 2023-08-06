package main

import (
	"context"
	"log"

	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/repository"
)

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)

	id, _ := userRepo.InsertUserData(context.Background(), entity.User{})
	log.Println(id)
	data, _ := userRepo.GetUserData(context.Background())
	log.Println(data)
}
