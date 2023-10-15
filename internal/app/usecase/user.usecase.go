package usecase

import (
	"context"

	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

type UserUsecase struct {
	AppConfig config.App
	UserRepo  interfaces.UserRepositoryInterface
}

func NewUserUsecase(cfgApp config.App, userRepo interfaces.UserRepositoryInterface) interfaces.UserUsecaseInterface {
	return &UserUsecase{
		AppConfig: cfgApp,
		UserRepo:  userRepo,
	}
}

func (u *UserUsecase) GetUserDataByDocRefID(ctx context.Context, docRefID string) (*entity.User, error) {
	user, err := u.UserRepo.GetUserDataByDocRefID(ctx, docRefID)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"doc_ref_if": docRefID,
			"error":      err.Error(),
		}).Error("[GetUserDataByDocRefID] failed on GetUserDataByDocRefID")

		return nil, err
	}

	return &user, nil
}

func (u *UserUsecase) GetAllUserData(ctx context.Context) ([]entity.User, error) {
	users, err := u.UserRepo.GetAllUserData(ctx)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[GetAllUserData] failed on GetUserData")

		return users, err
	}

	return users, nil
}

func (u *UserUsecase) InsertUserData(ctx context.Context, data entity.User) (string, error) {
	if u.AppConfig.IsUseFakeData {
		// Modify data here to use fake data if needed
		data.FirstName = fake.FirstName()
		data.LastName = fake.LastName()
		data.Username = fake.UserName()
		data.Email = fake.EmailAddress()
		data.CCNumber = fake.CreditCardNum("")
		data.CCType = fake.CreditCardType()
		data.Country = fake.Country()
		data.City = fake.City()
		data.Currency = fake.Currency()
	}

	docRefID, err := u.UserRepo.InsertUserData(ctx, data)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[InsertUserData] failed on InsertUserData")

		return "", err
	}

	return docRefID, nil
}

func (u *UserUsecase) DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error) {
	success, err := u.UserRepo.DeleteUserDataByDocRefID(ctx, docRefID)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"doc_ref_if": docRefID,
			"error":      err.Error(),
		}).Error("[DeleteUserDataByDocRefID] failed on DeleteUserDataByDocRefID")

		return false, err
	}

	return success, nil
}

func (u *UserUsecase) UpdateUserDataByDocRefID(ctx context.Context, docRefID string, data entity.User) error {
	err := u.UserRepo.UpdateUserData(ctx, docRefID, data)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"doc_ref_id": docRefID,
			"error":      err.Error(),
		}).Error("[UpdateUserDataByDocRefID] failed on UpdateUserData")

		return err
	}

	return nil
}
