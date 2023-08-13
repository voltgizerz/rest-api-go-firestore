package repository

import (
	"context"

	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
	"google.golang.org/api/iterator"

	"github.com/voltgizerz/rest-api-go-firestore/config"
)

const (
	USER_COLLECTION_NAME = "users"
)

type UserRepository struct {
	DB *config.Database
}

func NewUserRepository(db *config.Database) interfaces.UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetUserDataByDocRefID(ctx context.Context, docRefID string) (entity.User, error) {
	var user entity.User

	docRef := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Doc(docRefID)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return user, err
	}

	err = docSnapshot.DataTo(&user)
	if err != nil {
		return user, err
	}
	user.DocRefID = docRef.ID

	return user, nil
}

func (u *UserRepository) DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error) {
	docRef := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Doc(docRefID)
	_, err := docRef.Delete(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserRepository) GetAllUserData(ctx context.Context) ([]entity.User, error) {
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
		user.DocRefID = doc.Ref.ID

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) InsertUserData(ctx context.Context, data entity.User) (string, error) {
	// Currently using fake data
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
		return "", err
	}

	return docRef.ID, nil
}
