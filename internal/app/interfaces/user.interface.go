package interfaces

import (
	"context"

	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
)

type UserRepositoryInterface interface {
	GetUserData(ctx context.Context) ([]entity.User, error)
	InsertUserData(ctx context.Context, data entity.User) (string, error)
}

type UserUsecaseInterface interface {
	GetAllUserData(ctx context.Context) ([]entity.User, error)
}
