package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/interfaces"
	"github.com/voltgizerz/rest-api-go-firestore/logger"
)

type UserUsecase struct {
	UserRepo interfaces.UserRepositoryInterface
}

func NewUserUsecase(userRepo interfaces.UserRepositoryInterface) interfaces.UserUsecaseInterface {
	return &UserUsecase{
		UserRepo: userRepo,
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
