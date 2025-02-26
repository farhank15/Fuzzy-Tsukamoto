package activity_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/activity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateActivity(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	activity := &models.Activity{}

	err := mockRepo.CreateActivity(ctx, activity)
	assert.NoError(t, err)
}

func TestGetActivityByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetActivityByID(gomock.Any(), 1).Return(&models.Activity{ID: 1}, nil)

	ctx := context.Background()
	activity, err := mockRepo.GetActivityByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, activity)
	assert.Equal(t, 1, activity.ID)
}

func TestGetActivitiesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetActivitiesByUserID(gomock.Any(), 1).Return([]*models.Activity{{ID: 1}}, nil)

	ctx := context.Background()
	activities, err := mockRepo.GetActivitiesByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, activities)
	assert.Len(t, activities, 1)
	assert.Equal(t, 1, activities[0].ID)
}

func TestUpdateActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateActivity(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	activity := &models.Activity{ID: 1}

	err := mockRepo.UpdateActivity(ctx, activity)
	assert.NoError(t, err)
}

func TestDeleteActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteActivity(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteActivity(ctx, 1)
	assert.NoError(t, err)
}

func TestGetByStudentID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByStudentID(gomock.Any(), 1).Return([]models.Activity{{ID: 1}}, nil)

	ctx := context.Background()
	activities, err := mockRepo.GetByStudentID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, activities)
	assert.Len(t, activities, 1)
	assert.Equal(t, 1, activities[0].ID)
}

func TestGetByStudentID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByStudentID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	activities, err := mockRepo.GetByStudentID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, activities)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := activity.NewMockActivityRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetUserByID(gomock.Any(), 1).Return(&models.Users{ID: 1}, nil)

	ctx := context.Background()
	user, err := mockRepo.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
}
