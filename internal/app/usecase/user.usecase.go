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

func (u *UserUsecase) GetAllUserData(ctx context.Context) ([]entity.User, error) {
	users, err := u.UserRepo.GetUserData(ctx)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[GetAllUserData] failed on GetUserData")

		return users, err
	}

	return users, nil
}
