package achievement_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/achievement"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateAchievement(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	achievement := &models.Achievement{}

	err := mockRepo.CreateAchievement(ctx, achievement)
	assert.NoError(t, err)
}

func TestGetAchievementByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetAchievementByID(gomock.Any(), 1).Return(&models.Achievement{ID: 1}, nil)

	ctx := context.Background()
	achievement, err := mockRepo.GetAchievementByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, achievement)
	assert.Equal(t, 1, achievement.ID)
}

func TestGetAchievementsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetAchievementsByUserID(gomock.Any(), 1).Return([]*models.Achievement{{ID: 1}}, nil)

	ctx := context.Background()
	achievements, err := mockRepo.GetAchievementsByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, achievements)
	assert.Len(t, achievements, 1)
	assert.Equal(t, 1, achievements[0].ID)
}

func TestUpdateAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateAchievement(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	achievement := &models.Achievement{ID: 1}

	err := mockRepo.UpdateAchievement(ctx, achievement)
	assert.NoError(t, err)
}

func TestDeleteAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteAchievement(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteAchievement(ctx, 1)
	assert.NoError(t, err)
}

func TestGetByStudentID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByStudentID(gomock.Any(), 1).Return([]models.Achievement{{ID: 1}}, nil)

	ctx := context.Background()
	achievements, err := mockRepo.GetByStudentID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, achievements)
	assert.Len(t, achievements, 1)
	assert.Equal(t, 1, achievements[0].ID)
}

func TestGetByStudentID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := achievement.NewMockAchievementRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByStudentID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	achievements, err := mockRepo.GetByStudentID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, achievements)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
