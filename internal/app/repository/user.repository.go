package repository

import (
	"context"

	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
	"google.golang.org/api/iterator"

	"github.com/voltgizerz/rest-api-go-firestore/config"
)

const (
	USER_COLLECTION_NAME = "users"
)

type UserRepositoryInterface interface {
	GetUserData(ctx context.Context) ([]entity.User, error)
	InsertUserData(ctx context.Context, data entity.User) (string, error)
}

type UserRepository struct {
	DB *config.Database
}

func NewUserRepository(db *config.Database) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetUserData(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	iter := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[GetUserData] failed to iterate")

			return nil, err
		}

		var user entity.User
		err = doc.DataTo(&user)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[GetUserData] failed to read document data")

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) InsertUserData(ctx context.Context, data entity.User) (string, error) {
	docRef, _, err := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Add(ctx, map[string]interface{}{
		"firstname": fake.FirstName(),
		"lastname":  fake.LastName(),
		"username":  fake.UserName(),
		"email":     fake.EmailAddress(),
		"cc_num":    fake.CreditCardNum(""),
		"cc_type":   fake.CreditCardType(),
		"country":   fake.Country(),
		"city":      fake.City(),
		"currency":  fake.Currency(),
	})

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[InsertUserData] failed on insert user collection")

		return "", err
	}

	return docRef.ID, nil
}
