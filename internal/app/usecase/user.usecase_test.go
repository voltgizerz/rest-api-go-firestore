package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/voltgizerz/rest-api-go-firestore/config"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	"github.com/voltgizerz/rest-api-go-firestore/internal/app/mocks"
)

func TestGetAllUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	uc := UserUsecase{UserRepo: mockRepo}

	ctx := context.Background()

	mockUsers := []entity.User{
		{DocRefID: "1", FirstName: "Alice"},
		{DocRefID: "2", FirstName: "Bob"},
	}

	tests := []struct {
		name          string
		mockSetup     func()
		expectedUsers []entity.User
		expectedErr   error
	}{
		{
			name: "success returns user list",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAllUserData(gomock.Any()).
					Return(mockUsers, nil).
					Times(1)
			},
			expectedUsers: mockUsers,
			expectedErr:   nil,
		},
		{
			name: "repository error",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAllUserData(gomock.Any()).
					Return([]entity.User{}, errors.New("db failure")).
					Times(1)
			},
			expectedUsers: []entity.User{},
			expectedErr:   errors.New("db failure"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			users, err := uc.GetAllUserData(ctx)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Empty(t, users)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUsers, users)
			}
		})
	}
}

func TestGetUserDataByDocRefID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	uc := UserUsecase{UserRepo: mockRepo}

	ctx := context.Background()
	mockUser := entity.User{DocRefID: "123", FirstName: "Felix"}

	tests := []struct {
		name        string
		mockSetup   func()
		expected    *entity.User
		expectedErr error
	}{
		{
			name: "success returns user",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUserDataByDocRefID(gomock.Any(), "123").
					Return(mockUser, nil).
					Times(1)
			},
			expected:    &mockUser,
			expectedErr: nil,
		},
		{
			name: "repository error",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUserDataByDocRefID(gomock.Any(), "123").
					Return(entity.User{}, errors.New("not found")).
					Times(1)
			},
			expected:    nil,
			expectedErr: errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			user, err := uc.GetUserDataByDocRefID(ctx, "123")

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}

func TestInsertUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	ctx := context.Background()

	user := entity.User{FirstName: "Felix"}

	tests := []struct {
		name        string
		appCfg      config.App
		mockSetup   func()
		expectedID  string
		expectedErr error
	}{
		{
			name:   "success insert",
			appCfg: config.App{IsUseFakeData: false},
			mockSetup: func() {
				mockRepo.EXPECT().
					InsertUserData(gomock.Any(), gomock.Any()).
					Return("doc123", nil).
					Times(1)
			},
			expectedID:  "doc123",
			expectedErr: nil,
		},
		{
			name:   "repository error",
			appCfg: config.App{IsUseFakeData: false},
			mockSetup: func() {
				mockRepo.EXPECT().
					InsertUserData(gomock.Any(), gomock.Any()).
					Return("", errors.New("insert failed")).
					Times(1)
			},
			expectedID:  "",
			expectedErr: errors.New("insert failed"),
		},
		{
			name:   "use fake data",
			appCfg: config.App{IsUseFakeData: true},
			mockSetup: func() {
				mockRepo.EXPECT().
					InsertUserData(gomock.Any(), gomock.Any()).
					Return("fakeDoc", nil).
					Times(1)
			},
			expectedID:  "fakeDoc",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			uc := UserUsecase{UserRepo: mockRepo, AppConfig: tt.appCfg}

			docID, err := uc.InsertUserData(ctx, user)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Empty(t, docID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, docID)
			}
		})
	}
}

func TestDeleteUserDataByDocRefID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	uc := UserUsecase{UserRepo: mockRepo}
	ctx := context.Background()

	tests := []struct {
		name        string
		mockSetup   func()
		expectedOK  bool
		expectedErr error
	}{
		{
			name: "success delete",
			mockSetup: func() {
				mockRepo.EXPECT().
					DeleteUserDataByDocRefID(gomock.Any(), "id123").
					Return(true, nil).
					Times(1)
			},
			expectedOK:  true,
			expectedErr: nil,
		},
		{
			name: "repository error",
			mockSetup: func() {
				mockRepo.EXPECT().
					DeleteUserDataByDocRefID(gomock.Any(), "id123").
					Return(false, errors.New("delete failed")).
					Times(1)
			},
			expectedOK:  false,
			expectedErr: errors.New("delete failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			ok, err := uc.DeleteUserDataByDocRefID(ctx, "id123")

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.False(t, ok)
			} else {
				assert.NoError(t, err)
				assert.True(t, ok)
			}
		})
	}
}

func TestUpdateUserDataByDocRefID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	uc := UserUsecase{UserRepo: mockRepo}
	ctx := context.Background()
	user := entity.User{FirstName: "Felix"}

	tests := []struct {
		name        string
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "success update",
			mockSetup: func() {
				mockRepo.EXPECT().
					UpdateUserData(gomock.Any(), "id123", user).
					Return(nil).
					Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "repository error",
			mockSetup: func() {
				mockRepo.EXPECT().
					UpdateUserData(gomock.Any(), "id123", user).
					Return(errors.New("update failed")).
					Times(1)
			},
			expectedErr: errors.New("update failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := uc.UpdateUserDataByDocRefID(ctx, "id123", user)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
