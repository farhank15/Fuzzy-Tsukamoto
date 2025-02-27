package academic_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/academic"
	"go-tsukamoto/internal/app/models"
	mockAcademicRepo "go-tsukamoto/internal/app/repository/academic"
	mockPredicateRepo "go-tsukamoto/internal/app/repository/predicate"
	mockUserRepo "go-tsukamoto/internal/app/repository/user"
	academicService "go-tsukamoto/internal/app/service/academic"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &academic.CreateAcademicRequest{
			UserID:          1,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
		}

		// Mock user repository to return a user
		mockUserRepository.EXPECT().GetUserByID(ctx, req.UserID).Return(&models.Users{ID: req.UserID}, nil)

		// Mock academic repository to successfully create an academic record
		mockRepo.EXPECT().CreateAcademic(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, academic *models.Academic) error {
			academic.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateAcademic(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.UserID, response.UserID)
		assert.Equal(t, req.Ipk, response.Ipk)
		assert.Equal(t, req.RepeatedCourses, response.RepeatedCourses)
		assert.Equal(t, req.Semester, response.Semester)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, &req.PredicateID, response.PredicateID)
	})

	t.Run("User Not Found", func(t *testing.T) {
		req := &academic.CreateAcademicRequest{
			UserID: 999,
		}

		// Mock user repository to return nil, indicating user not found
		mockUserRepository.EXPECT().GetUserByID(ctx, req.UserID).Return(nil, nil)

		response, err := service.CreateAcademic(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("User Repository Error", func(t *testing.T) {
		req := &academic.CreateAcademicRequest{
			UserID: 1,
		}

		// Mock user repository to return an error
		mockUserRepository.EXPECT().GetUserByID(ctx, req.UserID).Return(nil, errors.New("database error"))

		response, err := service.CreateAcademic(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Academic Repository Error", func(t *testing.T) {
		req := &academic.CreateAcademicRequest{
			UserID: 1,
		}

		// Mock user repository to return a user
		mockUserRepository.EXPECT().GetUserByID(ctx, req.UserID).Return(&models.Users{ID: req.UserID}, nil)

		// Mock academic repository to return an error
		mockRepo.EXPECT().CreateAcademic(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateAcademic(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAcademicByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		academicID := 1
		academicModel := &models.Academic{
			ID:              academicID,
			UserID:          1,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(academicModel, nil)

		response, err := service.GetAcademicByID(ctx, academicID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, academicModel.ID, response.ID)
		assert.Equal(t, academicModel.UserID, response.UserID)
		assert.Equal(t, academicModel.Ipk, response.Ipk)
		assert.Equal(t, academicModel.RepeatedCourses, response.RepeatedCourses)
		assert.Equal(t, academicModel.Semester, response.Semester)
		assert.Equal(t, academicModel.Year, response.Year)
		assert.Equal(t, &academicModel.PredicateID, response.PredicateID)
		assert.Equal(t, academicModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, academicModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("Academic Not Found", func(t *testing.T) {
		academicID := 999

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(nil, nil)

		response, err := service.GetAcademicByID(ctx, academicID)

		assert.Error(t, err)
		assert.Equal(t, "academic record not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		academicID := 1

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(nil, errors.New("database error"))

		response, err := service.GetAcademicByID(ctx, academicID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAcademicsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		userID := 1
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
			{
				ID:              2,
				UserID:          userID,
				Ipk:             3.6,
				RepeatedCourses: 0,
				Semester:        2,
				Year:            2023,
				PredicateID:     1,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		}

		mockRepo.EXPECT().GetAcademicsByUserID(ctx, userID).Return(academicModels, nil)

		response, err := service.GetAcademicsByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, academicModels[0].ID, response[0].ID)
		assert.Equal(t, academicModels[0].UserID, response[0].UserID)
		assert.Equal(t, academicModels[0].Ipk, response[0].Ipk)
		assert.Equal(t, academicModels[0].Semester, response[0].Semester)

		// Verify second record
		assert.Equal(t, academicModels[1].ID, response[1].ID)
		assert.Equal(t, academicModels[1].UserID, response[1].UserID)
		assert.Equal(t, academicModels[1].Ipk, response[1].Ipk)
		assert.Equal(t, academicModels[1].Semester, response[1].Semester)
	})

	t.Run("Empty Result", func(t *testing.T) {
		userID := 999
		var emptyResult []*models.Academic

		mockRepo.EXPECT().GetAcademicsByUserID(ctx, userID).Return(emptyResult, nil)

		response, err := service.GetAcademicsByUserID(ctx, userID)

		assert.NoError(t, err)
		assert.NotNil(t, response) // Should return empty slice, not nil
		assert.Len(t, response, 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		userID := 1

		mockRepo.EXPECT().GetAcademicsByUserID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.GetAcademicsByUserID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetAllAcademics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		academicModels := []*models.Academic{
			{
				ID:              1,
				UserID:          1,
				Ipk:             3.5,
				RepeatedCourses: 0,
				Semester:        1,
				Year:            2023,
				PredicateID:     1,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			{
				ID:              2,
				UserID:          2,
				Ipk:             3.6,
				RepeatedCourses: 0,
				Semester:        2,
				Year:            2023,
				PredicateID:     1,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		}

		mockRepo.EXPECT().GetAllAcademics(ctx).Return(academicModels, nil)

		response, err := service.GetAllAcademics(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, academicModels[0].ID, response[0].ID)
		assert.Equal(t, academicModels[0].UserID, response[0].UserID)
		assert.Equal(t, academicModels[0].Ipk, response[0].Ipk)
		assert.Equal(t, academicModels[0].Semester, response[0].Semester)

		// Verify second record
		assert.Equal(t, academicModels[1].ID, response[1].ID)
		assert.Equal(t, academicModels[1].UserID, response[1].UserID)
		assert.Equal(t, academicModels[1].Ipk, response[1].Ipk)
		assert.Equal(t, academicModels[1].Semester, response[1].Semester)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetAllAcademics(ctx).Return(nil, errors.New("database error"))

		response, err := service.GetAllAcademics(ctx)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		academicID := 1
		userID := 1

		req := &academic.UpdateAcademicRequest{
			Ipk:             3.75,
			RepeatedCourses: 1,
			Semester:        2,
			Year:            2023,
			PredicateID:     2,
		}

		academicModel := &models.Academic{
			ID:              academicID,
			UserID:          userID,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(academicModel, nil)
		mockUserRepository.EXPECT().GetUserByID(ctx, userID).Return(&models.Users{ID: userID}, nil)
		mockRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, academic *models.Academic) error {
			assert.Equal(t, req.Ipk, academic.Ipk)
			assert.Equal(t, req.RepeatedCourses, academic.RepeatedCourses)
			assert.Equal(t, req.Semester, academic.Semester)
			assert.Equal(t, req.Year, academic.Year)
			assert.Equal(t, req.PredicateID, academic.PredicateID)
			assert.True(t, academic.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateAcademic(ctx, academicID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, academicID, response.ID)
		assert.Equal(t, userID, response.UserID)
		assert.Equal(t, req.Ipk, response.Ipk)
		assert.Equal(t, req.RepeatedCourses, response.RepeatedCourses)
		assert.Equal(t, req.Semester, response.Semester)
		assert.Equal(t, req.Year, response.Year)
		assert.Equal(t, &req.PredicateID, response.PredicateID)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("Academic Not Found", func(t *testing.T) {
		academicID := 999
		req := &academic.UpdateAcademicRequest{}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(nil, nil)

		response, err := service.UpdateAcademic(ctx, academicID, req)

		assert.Error(t, err)
		assert.Equal(t, "academic record not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("User Not Found", func(t *testing.T) {
		academicID := 1
		userID := 1
		req := &academic.UpdateAcademicRequest{}

		academicModel := &models.Academic{
			ID:     academicID,
			UserID: userID,
		}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(academicModel, nil)
		mockUserRepository.EXPECT().GetUserByID(ctx, userID).Return(nil, nil)

		response, err := service.UpdateAcademic(ctx, academicID, req)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Academic Repository Error", func(t *testing.T) {
		academicID := 1
		userID := 1
		req := &academic.UpdateAcademicRequest{}

		academicModel := &models.Academic{
			ID:     academicID,
			UserID: userID,
		}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(academicModel, nil)
		mockUserRepository.EXPECT().GetUserByID(ctx, userID).Return(&models.Users{ID: userID}, nil)
		mockRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateAcademic(ctx, academicID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})

	t.Run("User Repository Error", func(t *testing.T) {
		academicID := 1
		userID := 1
		req := &academic.UpdateAcademicRequest{}

		academicModel := &models.Academic{
			ID:     academicID,
			UserID: userID,
		}

		mockRepo.EXPECT().GetAcademicByID(ctx, academicID).Return(academicModel, nil)
		mockUserRepository.EXPECT().GetUserByID(ctx, userID).Return(nil, errors.New("database error"))

		response, err := service.UpdateAcademic(ctx, academicID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		academicID := 1

		mockRepo.EXPECT().DeleteAcademic(ctx, academicID).Return(nil)

		err := service.DeleteAcademic(ctx, academicID)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		academicID := 1

		mockRepo.EXPECT().DeleteAcademic(ctx, academicID).Return(errors.New("database error"))

		err := service.DeleteAcademic(ctx, academicID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestCreateAcademic_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()

	t.Run("Invalid UserID", func(t *testing.T) {
		req := &academic.CreateAcademicRequest{
			UserID:          0,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
		}

		mockUserRepository.EXPECT().GetUserByID(ctx, 0).Return(nil, nil)

		response, err := service.CreateAcademic(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestUpdateAcademic_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockUserRepository := mockUserRepo.NewMockUserRepositoryInterface(ctrl)
	mockPredicateRepository := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	service := academicService.NewAcademicService(mockRepo, mockUserRepository, mockPredicateRepository)
	ctx := context.Background()

	t.Run("Invalid Academic ID", func(t *testing.T) {
		req := &academic.UpdateAcademicRequest{
			Ipk:             3.75,
			RepeatedCourses: 1,
			Semester:        2,
			Year:            2023,
			PredicateID:     2,
		}

		// Karena service tidak melakukan validasi ID terlebih dahulu,
		// kita perlu mengharapkan pemanggilan GetAcademicByID
		mockRepo.EXPECT().GetAcademicByID(ctx, 0).Return(nil, nil)

		response, err := service.UpdateAcademic(ctx, 0, req)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
