package interfaces

import (
	"context"

	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
)

type UserRepositoryInterface interface {
	GetAllUserData(ctx context.Context) ([]entity.User, error)
	InsertUserData(ctx context.Context, data entity.User) (string, error)
	DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error)
	GetUserDataByDocRefID(ctx context.Context, docRefID string) (entity.User, error)
}

type UserUsecaseInterface interface {
	GetAllUserData(ctx context.Context) ([]entity.User, error)
	InsertUserData(ctx context.Context, data entity.User) (string, error)
	GetUserDataByDocRefID(ctx context.Context, docRefID string) (*entity.User, error)
	DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error)
}
