package repository

import (
	"context"
	"math/rand"

	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
	"google.golang.org/api/iterator"

	"github.com/voltgizerz/rest-api-go-firestore/config"
)

type UserRepositoryInterface interface {
	GetUserData(ctx context.Context) error
	InsertUserData(ctx context.Context) error
}

type UserRepository struct {
	DB *config.Database
}

func NewUserRepository(db *config.Database) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetUserData(ctx context.Context) error {
	iter := u.DB.FirestoreClient.Collection("users").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[GetUserData] failed to iterate")

			return err
		}

		logger.Log.Println(doc.Data())
	}

	return nil
}

func (u *UserRepository) InsertUserData(ctx context.Context) error {
	_, _, err := u.DB.FirestoreClient.Collection("users").Add(ctx, map[string]interface{}{
		"first": fake.FirstName(),
		"last":  fake.LastName(),
		"born":  rand.Intn(10000),
	})

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[InsertUserData] failed on insert user collection")

		return err
	}

	return nil
}
