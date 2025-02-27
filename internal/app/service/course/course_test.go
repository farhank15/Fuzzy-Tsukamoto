package course_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/course"
	"go-tsukamoto/internal/app/models"
	mockCourseRepo "go-tsukamoto/internal/app/repository/course"
	courseService "go-tsukamoto/internal/app/service/course"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := &course.CreateCourseRequest{
			Code:         "CS101",
			CourseName:   "Introduction to Computer Science",
			CreditCourse: 3,
		}

		mockRepo.EXPECT().CreateCourse(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, course *models.Course) error {
			course.ID = 1 // Simulate ID generation
			return nil
		})

		response, err := service.CreateCourse(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, req.Code, response.Code)
		assert.Equal(t, req.CourseName, response.CourseName)
		assert.Equal(t, req.CreditCourse, response.CreditCourse)
	})

	t.Run("Repository Error", func(t *testing.T) {
		req := &course.CreateCourseRequest{
			Code: "CS101",
		}

		mockRepo.EXPECT().CreateCourse(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.CreateCourse(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetCourseByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		courseID := 1
		courseModel := &models.Course{
			ID:           courseID,
			Code:         "CS101",
			CourseName:   "Introduction to Computer Science",
			CreditCourse: 3,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(courseModel, nil)

		response, err := service.GetCourseByID(ctx, courseID)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, courseModel.ID, response.ID)
		assert.Equal(t, courseModel.Code, response.Code)
		assert.Equal(t, courseModel.CourseName, response.CourseName)
		assert.Equal(t, courseModel.CreditCourse, response.CreditCourse)
		assert.Equal(t, courseModel.CreatedAt, response.CreatedAt)
		assert.Equal(t, courseModel.UpdatedAt, response.UpdatedAt)
	})

	t.Run("Course Not Found", func(t *testing.T) {
		courseID := 999

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(nil, nil)

		response, err := service.GetCourseByID(ctx, courseID)

		assert.Error(t, err)
		assert.Equal(t, "course not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		courseID := 1

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(nil, errors.New("database error"))

		response, err := service.GetCourseByID(ctx, courseID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestGetCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		courseModels := []*models.Course{
			{
				ID:           1,
				Code:         "CS101",
				CourseName:   "Introduction to Computer Science",
				CreditCourse: 3,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			{
				ID:           2,
				Code:         "CS102",
				CourseName:   "Data Structures",
				CreditCourse: 4,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		}

		mockRepo.EXPECT().GetCourses(ctx).Return(courseModels, nil)

		response, err := service.GetCourses(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response, 2)

		// Verify first record
		assert.Equal(t, courseModels[0].ID, response[0].ID)
		assert.Equal(t, courseModels[0].Code, response[0].Code)
		assert.Equal(t, courseModels[0].CourseName, response[0].CourseName)
		assert.Equal(t, courseModels[0].CreditCourse, response[0].CreditCourse)

		// Verify second record
		assert.Equal(t, courseModels[1].ID, response[1].ID)
		assert.Equal(t, courseModels[1].Code, response[1].Code)
		assert.Equal(t, courseModels[1].CourseName, response[1].CourseName)
		assert.Equal(t, courseModels[1].CreditCourse, response[1].CreditCourse)
	})

	t.Run("Empty Result", func(t *testing.T) {
		var emptyResult []*models.Course

		mockRepo.EXPECT().GetCourses(ctx).Return(emptyResult, nil)

		response, err := service.GetCourses(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, response) // Should return empty slice, not nil
		assert.Len(t, response, 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetCourses(ctx).Return(nil, errors.New("database error"))

		response, err := service.GetCourses(ctx)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestUpdateCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()
	now := time.Now()

	t.Run("Success", func(t *testing.T) {
		courseID := 1

		req := &course.UpdateCourseRequest{
			Code:         "CS101",
			CourseName:   "Introduction to Computer Science",
			CreditCourse: 3,
		}

		courseModel := &models.Course{
			ID:           courseID,
			Code:         "CS100",
			CourseName:   "Old Course Name",
			CreditCourse: 2,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(courseModel, nil)
		mockRepo.EXPECT().UpdateCourse(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, course *models.Course) error {
			assert.Equal(t, req.Code, course.Code)
			assert.Equal(t, req.CourseName, course.CourseName)
			assert.Equal(t, req.CreditCourse, course.CreditCourse)
			assert.True(t, course.UpdatedAt.After(now))
			return nil
		})

		response, err := service.UpdateCourse(ctx, courseID, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, courseID, response.ID)
		assert.Equal(t, req.Code, response.Code)
		assert.Equal(t, req.CourseName, response.CourseName)
		assert.Equal(t, req.CreditCourse, response.CreditCourse)
		assert.Equal(t, now, response.CreatedAt)
		assert.True(t, response.UpdatedAt.After(now))
	})

	t.Run("Course Not Found", func(t *testing.T) {
		courseID := 999
		req := &course.UpdateCourseRequest{}

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(nil, nil)

		response, err := service.UpdateCourse(ctx, courseID, req)

		assert.Error(t, err)
		assert.Equal(t, "course not found", err.Error())
		assert.Nil(t, response)
	})

	t.Run("Repository Error", func(t *testing.T) {
		courseID := 1
		req := &course.UpdateCourseRequest{}

		courseModel := &models.Course{
			ID: courseID,
		}

		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(courseModel, nil)
		mockRepo.EXPECT().UpdateCourse(ctx, gomock.Any()).Return(errors.New("database error"))

		response, err := service.UpdateCourse(ctx, courseID, req)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, response)
	})
}

func TestDeleteCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		courseID := 1
		courseModel := &models.Course{
			ID:           courseID,
			Code:         "CS101",
			CourseName:   "Introduction to Computer Science",
			CreditCourse: 3,
		}

		// Mock untuk GetCourseByID yang mengembalikan course yang valid
		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(courseModel, nil)

		// Mock untuk DeleteCourse
		mockRepo.EXPECT().DeleteCourse(ctx, courseID).Return(nil)

		err := service.DeleteCourse(ctx, courseID)

		assert.NoError(t, err)
	})

	t.Run("Course Not Found", func(t *testing.T) {
		courseID := 1

		// Mock untuk GetCourseByID yang mengembalikan nil (course tidak ditemukan)
		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(nil, nil)

		err := service.DeleteCourse(ctx, courseID)

		assert.Error(t, err)
		assert.Equal(t, "course not found", err.Error())
	})

	t.Run("GetCourseByID Error", func(t *testing.T) {
		courseID := 1

		// Mock untuk GetCourseByID yang mengembalikan error
		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(nil, errors.New("database error"))

		err := service.DeleteCourse(ctx, courseID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("DeleteCourse Error", func(t *testing.T) {
		courseID := 1
		courseModel := &models.Course{
			ID:           courseID,
			Code:         "CS101",
			CourseName:   "Introduction to Computer Science",
			CreditCourse: 3,
		}

		// Mock untuk GetCourseByID yang mengembalikan course yang valid
		mockRepo.EXPECT().GetCourseByID(ctx, courseID).Return(courseModel, nil)

		// Mock untuk DeleteCourse yang mengembalikan error
		mockRepo.EXPECT().DeleteCourse(ctx, courseID).Return(errors.New("database error"))

		err := service.DeleteCourse(ctx, courseID)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestImportCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockCourseRepo.NewMockCourseRepositoryInterface(ctrl)
	service := courseService.NewCourseService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		reqs := []course.CreateCourseRequest{
			{
				Code:         "CS101",
				CourseName:   "Introduction to Computer Science",
				CreditCourse: 3,
			},
			{
				Code:         "CS102",
				CourseName:   "Data Structures",
				CreditCourse: 4,
			},
		}

		mockRepo.EXPECT().CreateCourse(ctx, gomock.Any()).Return(nil).Times(len(reqs))

		err := service.ImportCourses(ctx, reqs)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		reqs := []course.CreateCourseRequest{
			{
				Code:         "CS101",
				CourseName:   "Introduction to Computer Science",
				CreditCourse: 3,
			},
		}

		mockRepo.EXPECT().CreateCourse(ctx, gomock.Any()).Return(errors.New("database error"))

		err := service.ImportCourses(ctx, reqs)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
