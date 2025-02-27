package achievement_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/achievement"
	"go-tsukamoto/internal/app/models"
	mockAchievementRepo "go-tsukamoto/internal/app/repository/achievement"
	achievementService "go-tsukamoto/internal/app/service/achievement"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &achievement.CreateAchievementRequest{
			UserID:      1,
			Title:       "Achievement Title",
			Certificate: true,
			Rank:        1,
			Level:       "National",
			Year:        2023,
		}

		mockRepo.EXPECT().CreateAchievement(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, achievement *models.Achievement) error {
			achievement.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateAchievement(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.UserID, response.UserID)
		assert.Equal(t, req.Title, response.Title)
		assert.Equal(t, req.Certificate, response.Certificate)
		assert.Equal(t, req.Rank, response.Rank)
		assert.Equal(t, req.Level, response.Level)
		assert.Equal(t, req.Year, response.Year)
	})

	t.Run("Repository Error", func(t *testing.T) {
		req := &achievement.CreateAchievementRequest{
			UserID: 1,
		}

		mockRepo.EXPECT().CreateAchievement(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateAchievement(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAchievementByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		achievementID := 1
		achievementModel := &models.Achievement{
			ID:          achievementID,
			UserID:      1,
			Title:       "Achievement Title",
			Certificate: true,
			Rank:        1,
			Level:       models.Level("National"),
			Year:        2023,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(achievementModel, nil)

		response, err := service.GetAchievementByID(ctx, achievementID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, achievementModel.ID, response.ID)
		assert.Equal(t, achievementModel.UserID, response.UserID)
		assert.Equal(t, achievementModel.Title, response.Title)
		assert.Equal(t, achievementModel.Certificate, response.Certificate)
		assert.Equal(t, achievementModel.Rank, response.Rank)
		assert.Equal(t, string(achievementModel.Level), response.Level) // Perubahan di sini
		assert.Equal(t, achievementModel.Year, response.Year)
		assert.Equal(t, achievementModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, achievementModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("Achievement Not Found", func(t *testing.T) {
		achievementID := 999

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(nil, nil)

		response, err := service.GetAchievementByID(ctx, achievementID)

		assert.Error(t, err)
		assert.Equal(t, "achievement not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		achievementID := 1

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(nil, errors.New("database error"))

		response, err := service.GetAchievementByID(ctx, achievementID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAchievementsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
		achievementModels := []*models.Achievement{
			{
				ID:          1,
				UserID:      userID,
				Title:       "Achievement Title 1",
				Certificate: true,
				Rank:        1,
				Level:       models.Level("National"),
				Year:        2023,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          2,
				UserID:      userID,
				Title:       "Achievement Title 2",
				Certificate: true,
				Rank:        2,
				Level:       models.Level("International"),
				Year:        2023,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		mockRepo.EXPECT().GetAchievementsByUserID(ctx, userID).Return(achievementModels, nil)

		response, err := service.GetAchievementsByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, achievementModels[0].ID, response[0].ID)
		assert.Equal(t, achievementModels[0].UserID, response[0].UserID)
		assert.Equal(t, achievementModels[0].Title, response[0].Title)
		assert.Equal(t, achievementModels[0].Certificate, response[0].Certificate)
		assert.Equal(t, achievementModels[0].Rank, response[0].Rank)
		assert.Equal(t, string(achievementModels[0].Level), response[0].Level)
		assert.Equal(t, achievementModels[0].Year, response[0].Year)

		// Verify second record
		assert.Equal(t, achievementModels[1].ID, response[1].ID)
		assert.Equal(t, achievementModels[1].UserID, response[1].UserID)
		assert.Equal(t, achievementModels[1].Title, response[1].Title)
		assert.Equal(t, achievementModels[1].Certificate, response[1].Certificate)
		assert.Equal(t, achievementModels[1].Rank, response[1].Rank)
		assert.Equal(t, string(achievementModels[1].Level), response[1].Level)
		assert.Equal(t, achievementModels[1].Year, response[1].Year)
	})

	t.Run("Empty Result", func(t *testing.T) {
		userID := 999
		var emptyResult []*models.Achievement = []*models.Achievement{}

		mockRepo.EXPECT().GetAchievementsByUserID(ctx, userID).Return(emptyResult, nil)

		response, err := service.GetAchievementsByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetAchievementsByUserID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetAchievementsByUserID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAllAchievements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		achievementModels := []*models.Achievement{
			{
				ID:          1,
				UserID:      1,
				Title:       "Achievement Title 1",
				Certificate: true,
				Rank:        1,
				Level:       models.Level("National"),
				Year:        2023,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          2,
				UserID:      2,
				Title:       "Achievement Title 2",
				Certificate: true,
				Rank:        2,
				Level:       models.Level("International"),
				Year:        2023,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		mockRepo.EXPECT().GetAllAchievements(ctx).Return(achievementModels, nil)

		response, err := service.GetAllAchievements(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, achievementModels[0].ID, response[0].ID)
		assert.Equal(t, achievementModels[0].UserID, response[0].UserID)
		assert.Equal(t, achievementModels[0].Title, response[0].Title)
		assert.Equal(t, achievementModels[0].Certificate, response[0].Certificate)
		assert.Equal(t, achievementModels[0].Rank, response[0].Rank)
		assert.Equal(t, string(achievementModels[0].Level), response[0].Level)
		assert.Equal(t, achievementModels[0].Year, response[0].Year)

		// Verify second record
		assert.Equal(t, achievementModels[1].ID, response[1].ID)
		assert.Equal(t, achievementModels[1].UserID, response[1].UserID)
		assert.Equal(t, achievementModels[1].Title, response[1].Title)
		assert.Equal(t, achievementModels[1].Certificate, response[1].Certificate)
		assert.Equal(t, achievementModels[1].Rank, response[1].Rank)
		assert.Equal(t, string(achievementModels[1].Level), response[1].Level)
		assert.Equal(t, achievementModels[1].Year, response[1].Year)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetAllAchievements(ctx).Return(nil, errors.New("database error"))

		response, err := service.GetAllAchievements(ctx)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		achievementID := 1

		req := &achievement.UpdateAchievementRequest{
			Title:       "Updated Title",
			Certificate: true,
			Rank:        2,
			Level:       "International",
			Year:        2024,
		}

		achievementModel := &models.Achievement{
			ID:          achievementID,
			UserID:      1,
			Title:       "Achievement Title",
			Certificate: true,
			Rank:        1,
			Level:       models.Level("National"),
			Year:        2023,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(achievementModel, nil)
		mockRepo.EXPECT().UpdateAchievement(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, achievement *models.Achievement) error {
			assert.Equal(t, req.Title, achievement.Title)
			assert.Equal(t, req.Certificate, achievement.Certificate)
			assert.Equal(t, req.Rank, achievement.Rank)
			assert.Equal(t, models.Level(req.Level), achievement.Level)
			assert.Equal(t, req.Year, achievement.Year)
			assert.True(t, achievement.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateAchievement(ctx, achievementID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, achievementID, response.ID)
		assert.Equal(t, req.Title, response.Title)
		assert.Equal(t, req.Certificate, response.Certificate)
		assert.Equal(t, req.Rank, response.Rank)
		assert.Equal(t, req.Level, response.Level)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("Achievement Not Found", func(t *testing.T) {
		achievementID := 999
		req := &achievement.UpdateAchievementRequest{}

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(nil, nil)

		response, err := service.UpdateAchievement(ctx, achievementID, req)

		assert.Error(t, err)
		assert.Equal(t, "achievement not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		achievementID := 1
		req := &achievement.UpdateAchievementRequest{}

		achievementModel := &models.Achievement{
			ID: achievementID,
		}

		mockRepo.EXPECT().GetAchievementByID(ctx, achievementID).Return(achievementModel, nil)
		mockRepo.EXPECT().UpdateAchievement(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateAchievement(ctx, achievementID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteAchievement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	service := achievementService.NewAchievementService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		achievementID := 1

		mockRepo.EXPECT().DeleteAchievement(ctx, achievementID).Return(nil)

		err := service.DeleteAchievement(ctx, achievementID)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		achievementID := 1

		mockRepo.EXPECT().DeleteAchievement(ctx, achievementID).Return(errors.New("database error"))

		err := service.DeleteAchievement(ctx, achievementID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
