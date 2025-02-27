package thesis_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/thesis"
	"go-tsukamoto/internal/app/models"
	mockThesisRepo "go-tsukamoto/internal/app/repository/thesis"
	mockUserRepo "go-tsukamoto/internal/app/repository/user"
	thesisService "go-tsukamoto/internal/app/service/thesis"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	mockUserRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, mockUserRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &thesis.CreateThesisRequest{
			UserID:   1,
			Title:    "Thesis Title",
			Year:     2023,
			Semester: 1,
			Value:    "A",
			Level:    "National",
		}

		mockUserRepo.EXPECT().GetUserByID(ctx, req.UserID).Return(&models.Users{ID: req.UserID}, nil)
		mockRepo.EXPECT().CreateThesis(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, thesis *models.Thesis) error {
			thesis.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateThesis(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.UserID, response.UserID)
		assert.Equal(t, req.Title, response.Title)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, req.Semester, response.Semester)
		assert.Equal(t, req.Value, response.Value)
		assert.Equal(t, req.Level, response.Level)
	})

	t.Run("User Not Found", func(t *testing.T) {
		req := &thesis.CreateThesisRequest{
			UserID: 1,
		}

		mockUserRepo.EXPECT().GetUserByID(ctx, req.UserID).Return(nil, nil)

		response, err := service.CreateThesis(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		req := &thesis.CreateThesisRequest{
			UserID: 1,
		}

		mockUserRepo.EXPECT().GetUserByID(ctx, req.UserID).Return(&models.Users{ID: req.UserID}, nil)
		mockRepo.EXPECT().CreateThesis(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateThesis(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetThesisByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, nil)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		thesisID := 1
		thesisModel := &models.Thesis{
			ID:        thesisID,
			UserID:    1,
			Title:     "Thesis Title",
			Year:      2023,
			Semester:  1,
			Value:     "A",
			Level:     "National",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(thesisModel, nil)

		response, err := service.GetThesisByID(ctx, thesisID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, thesisModel.ID, response.ID)
		assert.Equal(t, thesisModel.UserID, response.UserID)
		assert.Equal(t, thesisModel.Title, response.Title)
		assert.Equal(t, thesisModel.Year, response.Year)
		assert.Equal(t, thesisModel.Semester, response.Semester)
		assert.Equal(t, thesisModel.Value, response.Value)
		assert.Equal(t, thesisModel.Level, response.Level)
		assert.Equal(t, thesisModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, thesisModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("Thesis Not Found", func(t *testing.T) {
		thesisID := 999

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(nil, nil)

		response, err := service.GetThesisByID(ctx, thesisID)

		assert.Error(t, err)
		assert.Equal(t, "thesis not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		thesisID := 1

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(nil, errors.New("database error"))

		response, err := service.GetThesisByID(ctx, thesisID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetThesesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, nil)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
		thesisModels := []*models.Thesis{
			{
				ID:        1,
				UserID:    userID,
				Title:     "Thesis Title 1",
				Year:      2023,
				Semester:  1,
				Value:     "A",
				Level:     "National",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				UserID:    userID,
				Title:     "Thesis Title 2",
				Year:      2023,
				Semester:  2,
				Value:     "B",
				Level:     "International",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		mockRepo.EXPECT().GetThesesByUserID(ctx, userID).Return(thesisModels, nil)

		response, err := service.GetThesesByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, thesisModels[0].ID, response[0].ID)
		assert.Equal(t, thesisModels[0].UserID, response[0].UserID)
		assert.Equal(t, thesisModels[0].Title, response[0].Title)
		assert.Equal(t, thesisModels[0].Year, response[0].Year)
		assert.Equal(t, thesisModels[0].Semester, response[0].Semester)
		assert.Equal(t, thesisModels[0].Value, response[0].Value)
		assert.Equal(t, thesisModels[0].Level, response[0].Level)
		assert.Equal(t, thesisModels[0].CreatedAt, response[0].CreatedAt)
		assert.Equal(t, thesisModels[0].UpdatedAt, response[0].UpdatedAt)

		// Verify second record
		assert.Equal(t, thesisModels[1].ID, response[1].ID)
		assert.Equal(t, thesisModels[1].UserID, response[1].UserID)
		assert.Equal(t, thesisModels[1].Title, response[1].Title)
		assert.Equal(t, thesisModels[1].Year, response[1].Year)
		assert.Equal(t, thesisModels[1].Semester, response[1].Semester)
		assert.Equal(t, thesisModels[1].Value, response[1].Value)
		assert.Equal(t, thesisModels[1].Level, response[1].Level)
		assert.Equal(t, thesisModels[1].CreatedAt, response[1].CreatedAt)
		assert.Equal(t, thesisModels[1].UpdatedAt, response[1].UpdatedAt)
	})

	t.Run("Empty Result", func(t *testing.T) {
		userID := 999
		var emptyResult []*models.Thesis

		mockRepo.EXPECT().GetThesesByUserID(ctx, userID).Return(emptyResult, nil)

		response, err := service.GetThesesByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response) // Should return empty slice, not nil
		assert.Len(t, response, 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetThesesByUserID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetThesesByUserID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAllTheses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, nil)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		thesisModels := []*models.Thesis{
			{
				ID:        1,
				UserID:    1,
				Title:     "Thesis Title 1",
				Year:      2023,
				Semester:  1,
				Value:     "A",
				Level:     "National",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				UserID:    2,
				Title:     "Thesis Title 2",
				Year:      2023,
				Semester:  2,
				Value:     "B",
				Level:     "International",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		mockRepo.EXPECT().GetAllTheses(ctx).Return(thesisModels, nil)

		response, err := service.GetAllTheses(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, thesisModels[0].ID, response[0].ID)
		assert.Equal(t, thesisModels[0].UserID, response[0].UserID)
		assert.Equal(t, thesisModels[0].Title, response[0].Title)
		assert.Equal(t, thesisModels[0].Year, response[0].Year)
		assert.Equal(t, thesisModels[0].Semester, response[0].Semester)
		assert.Equal(t, thesisModels[0].Value, response[0].Value)
		assert.Equal(t, thesisModels[0].Level, response[0].Level)
		assert.Equal(t, thesisModels[0].CreatedAt, response[0].CreatedAt)
		assert.Equal(t, thesisModels[0].UpdatedAt, response[0].UpdatedAt)

		// Verify second record
		assert.Equal(t, thesisModels[1].ID, response[1].ID)
		assert.Equal(t, thesisModels[1].UserID, response[1].UserID)
		assert.Equal(t, thesisModels[1].Title, response[1].Title)
		assert.Equal(t, thesisModels[1].Year, response[1].Year)
		assert.Equal(t, thesisModels[1].Semester, response[1].Semester)
		assert.Equal(t, thesisModels[1].Value, response[1].Value)
		assert.Equal(t, thesisModels[1].Level, response[1].Level)
		assert.Equal(t, thesisModels[1].CreatedAt, response[1].CreatedAt)
		assert.Equal(t, thesisModels[1].UpdatedAt, response[1].UpdatedAt)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetAllTheses(ctx).Return(nil, errors.New("database error"))

		response, err := service.GetAllTheses(ctx)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	mockUserRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, mockUserRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		thesisID := 1

		req := &thesis.UpdateThesisRequest{
			Title:    "Updated Title",
			Year:     2024,
			Semester: 2,
			Value:    "B",
			Level:    "International",
		}

		thesisModel := &models.Thesis{
			ID:        thesisID,
			UserID:    1,
			Title:     "Thesis Title",
			Year:      2023,
			Semester:  1,
			Value:     "A",
			Level:     "National",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(thesisModel, nil)
		mockUserRepo.EXPECT().GetUserByID(ctx, thesisModel.UserID).Return(&models.Users{ID: thesisModel.UserID}, nil)
		mockRepo.EXPECT().UpdateThesis(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, thesis *models.Thesis) error {
			assert.Equal(t, req.Title, thesis.Title)
			assert.Equal(t, req.Year, thesis.Year)
			assert.Equal(t, req.Semester, thesis.Semester)
			assert.Equal(t, req.Value, thesis.Value)
			assert.Equal(t, req.Level, thesis.Level)
			assert.True(t, thesis.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateThesis(ctx, thesisID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, thesisID, response.ID)
		assert.Equal(t, req.Title, response.Title)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, req.Semester, response.Semester)
		assert.Equal(t, req.Value, response.Value)
		assert.Equal(t, req.Level, response.Level)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("Thesis Not Found", func(t *testing.T) {
		thesisID := 999
		req := &thesis.UpdateThesisRequest{}

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(nil, nil)

		response, err := service.UpdateThesis(ctx, thesisID, req)

		assert.Error(t, err)
		assert.Equal(t, "thesis not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("User Not Found", func(t *testing.T) {
		thesisID := 1
		req := &thesis.UpdateThesisRequest{}

		thesisModel := &models.Thesis{
			ID: thesisID,
		}

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(thesisModel, nil)
		mockUserRepo.EXPECT().GetUserByID(ctx, thesisModel.UserID).Return(nil, nil)

		response, err := service.UpdateThesis(ctx, thesisID, req)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		thesisID := 1
		req := &thesis.UpdateThesisRequest{}

		thesisModel := &models.Thesis{
			ID: thesisID,
		}

		mockRepo.EXPECT().GetThesisByID(ctx, thesisID).Return(thesisModel, nil)
		mockUserRepo.EXPECT().GetUserByID(ctx, thesisModel.UserID).Return(&models.Users{ID: thesisModel.UserID}, nil)
		mockRepo.EXPECT().UpdateThesis(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateThesis(ctx, thesisID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	service := thesisService.NewThesisService(mockRepo, nil)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		thesisID := 1

		mockRepo.EXPECT().DeleteThesis(ctx, thesisID).Return(nil)

		err := service.DeleteThesis(ctx, thesisID)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		thesisID := 1

		mockRepo.EXPECT().DeleteThesis(ctx, thesisID).Return(errors.New("database error"))

		err := service.DeleteThesis(ctx, thesisID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
