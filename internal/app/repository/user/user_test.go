package user_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	user := &models.Users{}

	err := mockRepo.CreateUser(ctx, user)
	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetUserByID(gomock.Any(), 1).Return(&models.Users{ID: 1}, nil)

	ctx := context.Background()
	user, err := mockRepo.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
}

func TestGetUserByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetUserByID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	user, err := mockRepo.GetUserByID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetUserByUsername(gomock.Any(), "example").Return(&models.Users{ID: 1, Username: "example"}, nil)

	ctx := context.Background()
	user, err := mockRepo.GetUserByUsername(ctx, "example")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "example", user.Username)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetUserByUsername(gomock.Any(), "example").Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	user, err := mockRepo.GetUserByUsername(ctx, "example")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	user := &models.Users{ID: 1}

	err := mockRepo.UpdateUser(ctx, user)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteUser(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteUser(ctx, 1)
	assert.NoError(t, err)
}
