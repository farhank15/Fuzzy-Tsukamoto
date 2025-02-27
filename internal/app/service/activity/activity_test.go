package activity_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/activity"
	"go-tsukamoto/internal/app/models"
	mockActivityRepo "go-tsukamoto/internal/app/repository/activity"
	activityService "go-tsukamoto/internal/app/service/activity"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &activity.CreateActivityRequest{
			UserID:       1,
			Organization: "Organization Name",
			Year:         2023,
		}

		mockRepo.EXPECT().CreateActivity(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, activity *models.Activity) error {
			activity.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateActivity(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.UserID, response.UserID)
		assert.Equal(t, req.Organization, response.Organization)
		assert.Equal(t, req.Year, response.Year)
	})

	t.Run("Repository Error", func(t *testing.T) {
		req := &activity.CreateActivityRequest{
			UserID: 1,
		}

		mockRepo.EXPECT().CreateActivity(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateActivity(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetActivityByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		activityID := 1
		activityModel := &models.Activity{
			ID:           activityID,
			UserID:       1,
			Organization: "Organization Name",
			Year:         2023,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(activityModel, nil)

		response, err := service.GetActivityByID(ctx, activityID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, activityModel.ID, response.ID)
		assert.Equal(t, activityModel.UserID, response.UserID)
		assert.Equal(t, activityModel.Organization, response.Organization)
		assert.Equal(t, activityModel.Year, response.Year)
		assert.Equal(t, activityModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, activityModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("Activity Not Found", func(t *testing.T) {
		activityID := 999

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(nil, nil)

		response, err := service.GetActivityByID(ctx, activityID)

		assert.Error(t, err)
		assert.Equal(t, "activity not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		activityID := 1

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(nil, errors.New("database error"))

		response, err := service.GetActivityByID(ctx, activityID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetActivitiesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
		activityModels := []*models.Activity{
			{
				ID:           1,
				UserID:       userID,
				Organization: "Organization Name 1",
				Year:         2023,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			{
				ID:           2,
				UserID:       userID,
				Organization: "Organization Name 2",
				Year:         2023,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		}

		mockRepo.EXPECT().GetActivitiesByUserID(ctx, userID).Return(activityModels, nil)

		response, err := service.GetActivitiesByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, activityModels[0].ID, response[0].ID)
		assert.Equal(t, activityModels[0].UserID, response[0].UserID)
		assert.Equal(t, activityModels[0].Organization, response[0].Organization)
		assert.Equal(t, activityModels[0].Year, response[0].Year)

		// Verify second record
		assert.Equal(t, activityModels[1].ID, response[1].ID)
		assert.Equal(t, activityModels[1].UserID, response[1].UserID)
		assert.Equal(t, activityModels[1].Organization, response[1].Organization)
		assert.Equal(t, activityModels[1].Year, response[1].Year)
	})

	t.Run("Empty Result", func(t *testing.T) {
		userID := 999
		var emptyResult []*models.Activity

		mockRepo.EXPECT().GetActivitiesByUserID(ctx, userID).Return(emptyResult, nil)

		response, err := service.GetActivitiesByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response) // Should return empty slice, not nil
		assert.Len(t, response, 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetActivitiesByUserID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetActivitiesByUserID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAllActivities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		activityModels := []*models.Activity{
			{
				ID:           1,
				UserID:       1,
				Organization: "Organization Name 1",
				Year:         2023,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			{
				ID:           2,
				UserID:       2,
				Organization: "Organization Name 2",
				Year:         2023,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		}

		mockRepo.EXPECT().GetAllActivities(ctx).Return(activityModels, nil)

		response, err := service.GetAllActivities(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, activityModels[0].ID, response[0].ID)
		assert.Equal(t, activityModels[0].UserID, response[0].UserID)
		assert.Equal(t, activityModels[0].Organization, response[0].Organization)
		assert.Equal(t, activityModels[0].Year, response[0].Year)

		// Verify second record
		assert.Equal(t, activityModels[1].ID, response[1].ID)
		assert.Equal(t, activityModels[1].UserID, response[1].UserID)
		assert.Equal(t, activityModels[1].Organization, response[1].Organization)
		assert.Equal(t, activityModels[1].Year, response[1].Year)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetAllActivities(ctx).Return(nil, errors.New("database error"))

		response, err := service.GetAllActivities(ctx)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		activityID := 1

		req := &activity.UpdateActivityRequest{
			Organization: "Updated Organization",
			Year:         2024,
		}

		activityModel := &models.Activity{
			ID:           activityID,
			UserID:       1,
			Organization: "Organization Name",
			Year:         2023,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(activityModel, nil)
		mockRepo.EXPECT().UpdateActivity(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, activity *models.Activity) error {
			assert.Equal(t, req.Organization, activity.Organization)
			assert.Equal(t, req.Year, activity.Year)
			assert.True(t, activity.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateActivity(ctx, activityID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, activityID, response.ID)
		assert.Equal(t, req.Organization, response.Organization)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("Activity Not Found", func(t *testing.T) {
		activityID := 999
		req := &activity.UpdateActivityRequest{}

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(nil, nil)

		response, err := service.UpdateActivity(ctx, activityID, req)

		assert.Error(t, err)
		assert.Equal(t, "activity not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		activityID := 1
		req := &activity.UpdateActivityRequest{}

		activityModel := &models.Activity{
			ID: activityID,
		}

		mockRepo.EXPECT().GetActivityByID(ctx, activityID).Return(activityModel, nil)
		mockRepo.EXPECT().UpdateActivity(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateActivity(ctx, activityID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteActivity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	service := activityService.NewActivityService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		activityID := 1

		mockRepo.EXPECT().DeleteActivity(ctx, activityID).Return(nil)

		err := service.DeleteActivity(ctx, activityID)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		activityID := 1

		mockRepo.EXPECT().DeleteActivity(ctx, activityID).Return(errors.New("database error"))

		err := service.DeleteActivity(ctx, activityID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
