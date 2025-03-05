package user_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/user"
	"go-tsukamoto/internal/app/models"
	mockAcademicRepo "go-tsukamoto/internal/app/repository/academic"
	mockAchievementRepo "go-tsukamoto/internal/app/repository/achievement"
	mockActivityRepo "go-tsukamoto/internal/app/repository/activity"
	mockThesisRepo "go-tsukamoto/internal/app/repository/thesis"
	mockUserRepo "go-tsukamoto/internal/app/repository/user"
	userService "go-tsukamoto/internal/app/service/user"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := userService.NewUserService(mockRepo, nil, nil, nil, nil)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &user.CreateUserRequest{
			Username:  "testuser",
			Name:      "Test User",
			Nim:       "123456789",
			Password:  "password",
			StartYear: 2023,
		}

		mockRepo.EXPECT().GetUserByNim(ctx, req.Nim).Return(nil, nil)
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, user *models.Users) error {
			user.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateUser(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.Username, response.Username)
		assert.Equal(t, req.Name, response.Name)
		assert.Equal(t, req.Nim, response.Nim)
		assert.Equal(t, req.StartYear, response.StartYear)
	})

	t.Run("NIM Already Exists", func(t *testing.T) {
		req := &user.CreateUserRequest{
			Username: "testuser",
			Nim:      "123456789",
		}

		mockRepo.EXPECT().GetUserByNim(ctx, req.Nim).Return(&models.Users{ID: 1}, nil)

		response, err := service.CreateUser(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "NIM already exists", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		req := &user.CreateUserRequest{
			Username: "testuser",
		}

		mockRepo.EXPECT().GetUserByNim(ctx, req.Nim).Return(nil, nil)
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateUser(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := userService.NewUserService(mockRepo, nil, nil, nil, nil)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
		userModel := &models.Users{
			ID:        userID,
			Username:  "testuser",
			Name:      "Test User",
			Nim:       "123456789",
			Password:  "hashedpassword",
			StartYear: 2023,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(userModel, nil)

		response, err := service.GetUserByID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, userModel.ID, response.ID)
		assert.Equal(t, userModel.Username, response.Username)
		assert.Equal(t, userModel.Name, response.Name)
		assert.Equal(t, userModel.Nim, response.Nim)
		assert.Equal(t, userModel.Password, response.Password)
		assert.Equal(t, userModel.StartYear, response.StartYear)
		assert.Equal(t, userModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, userModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("User Not Found", func(t *testing.T) {
		userID := 999

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(nil, nil)

		response, err := service.GetUserByID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetUserByID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := userService.NewUserService(mockRepo, nil, nil, nil, nil)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1

		req := &user.UpdateUserRequest{
			Username:  "updateduser",
			Name:      "Updated User",
			Nim:       "987654321",
			StartYear: 2024,
		}

		userModel := &models.Users{
			ID:        userID,
			Username:  "testuser",
			Name:      "Test User",
			Nim:       "123456789",
			Password:  "hashedpassword",
			StartYear: 2023,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(userModel, nil)
		mockRepo.EXPECT().UpdateUser(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, user *models.Users) error {
			assert.Equal(t, req.Username, user.Username)
			assert.Equal(t, req.Name, user.Name)
			assert.Equal(t, req.Nim, user.Nim)
			assert.Equal(t, req.StartYear, user.StartYear)
			assert.True(t, user.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateUser(ctx, userID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, userID, response.ID)
		assert.Equal(t, req.Username, response.Username)
		assert.Equal(t, req.Name, response.Name)
		assert.Equal(t, req.Nim, response.Nim)
		assert.Equal(t, req.StartYear, response.StartYear)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("User Not Found", func(t *testing.T) {
		userID := 999
		req := &user.UpdateUserRequest{}

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(nil, nil)

		response, err := service.UpdateUser(ctx, userID, req)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1
		req := &user.UpdateUserRequest{}

		userModel := &models.Users{
			ID: userID,
		}

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(userModel, nil)
		mockRepo.EXPECT().UpdateUser(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateUser(ctx, userID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	service := userService.NewUserService(mockRepo, nil, nil, nil, nil)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().DeleteUser(ctx, userID).Return(nil)

		err := service.DeleteUser(ctx, userID)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().DeleteUser(ctx, userID).Return(errors.New("database error"))

		err := service.DeleteUser(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestGetUserWithRelatedData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockAcademicRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockAchievementRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	mockActivityRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	mockThesisRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	service := userService.NewUserService(mockRepo, mockAcademicRepo, mockAchievementRepo, mockActivityRepo, mockThesisRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
		userModel := &models.Users{
			ID:        userID,
			Username:  "testuser",
			Name:      "Test User",
			Nim:       "123456789",
			Password:  "hashedpassword",
			StartYear: 2023,
			CreatedAt: now,
			UpdatedAt: now,
		}

		academicModels := []*models.Academic{
			{
				ID:              1,
				UserID:          userID,
				Ipk:             3.5,
				RepeatedCourses: 0,
				Semester:        1,
				Year:            2023,
				PredicateID:     1,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		}

		achievementModels := []*models.Achievement{
			{
				ID:          1,
				UserID:      userID,
				Title:       "Achievement Title",
				Certificate: true,
				Rank:        1,
				Level:       models.LevelInternasional,
				Year:        2023,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		activityModels := []*models.Activity{
			{
				ID:           1,
				UserID:       userID,
				Organization: "Organization Name",
				Year:         2023,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		}

		thesisModels := []*models.Thesis{
			{
				ID:     1,
				UserID: userID,
				Level:  "internasional",
			},
		}

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(userModel, nil)
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, userID).Return(academicModels, nil)
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, userID).Return(achievementModels, nil)
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, userID).Return(activityModels, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, userID).Return(thesisModels, nil)

		response, err := service.GetUserWithRelatedData(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, userModel.ID, response.ID)
		assert.Equal(t, userModel.Username, response.Username)
		assert.Equal(t, userModel.Name, response.Name)
		assert.Equal(t, userModel.Nim, response.Nim)
		assert.Equal(t, userModel.Password, response.Password)
		assert.Equal(t, userModel.StartYear, response.StartYear)
		assert.Equal(t, userModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, userModel.UpdatedAt, response.UpdatedAt)
		assert.Len(t, response.Academics, 1)
		assert.Len(t, response.Achievements, 1)
		assert.Len(t, response.Activities, 1)
		assert.Len(t, response.Theses, 1)
	})

	t.Run("User Not Found", func(t *testing.T) {
		userID := 999

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(nil, nil)

		response, err := service.GetUserWithRelatedData(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetUserByID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetUserWithRelatedData(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}
