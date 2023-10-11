package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/database"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
	"google.golang.org/api/iterator"
)

const (
	USER_COLLECTION_NAME = "users"
)

type UserRepository struct {
	DB *database.Database
}

func NewUserRepository(db *database.Database) interfaces.UserRepositoryInterface {
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

// TODO handle collection not found
// TODO NOTE : Using Snapshot To Validate
func (u *UserRepository) DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error) {
	docRef := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Doc(docRefID)
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		return false, nil
	}
	if snapshot.Exists() {
		_, err := docRef.Delete(ctx)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil

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
	docData := map[string]interface{}{
		"firstname": data.FirstName,
		"lastname":  data.LastName,
		"username":  data.Username,
		"email":     data.Email,
		"cc_num":    data.CCNumber,
		"cc_type":   data.CCType,
		"country":   data.Country,
		"city":      data.City,
		"currency":  data.Currency,
	}

	docRef, _, err := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Add(ctx, docData)
	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

func (u *UserRepository) UpdateUserData(ctx context.Context, docRefID string, data entity.User) error {
	updateData := make(map[string]interface{})

	if data.FirstName != "" {
		updateData["firstname"] = data.FirstName
	}
	if data.LastName != "" {
		updateData["lastname"] = data.LastName
	}
	if data.Username != "" {
		updateData["username"] = data.Username
	}
	if data.Email != "" {
		updateData["email"] = data.Email
	}
	if data.CCNumber != "" {
		updateData["cc_num"] = data.CCNumber
	}
	if data.CCType != "" {
		updateData["cc_type"] = data.CCType
	}
	if data.Country != "" {
		updateData["country"] = data.Country
	}
	if data.City != "" {
		updateData["city"] = data.City
	}
	if data.Currency != "" {
		updateData["currency"] = data.Currency
	}

	_, err := u.DB.FirestoreClient.Collection(USER_COLLECTION_NAME).Doc(docRefID).Set(ctx, updateData, firestore.MergeAll)

	return err
}
