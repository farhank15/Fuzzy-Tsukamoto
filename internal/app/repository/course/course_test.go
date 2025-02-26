package course_test

import (
	"context"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/course"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := course.NewMockCourseRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateCourse(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	course := &models.Course{}

	err := mockRepo.CreateCourse(ctx, course)
	assert.NoError(t, err)
}

func TestGetCourseByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := course.NewMockCourseRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetCourseByID(gomock.Any(), 1).Return(&models.Course{ID: 1}, nil)

	ctx := context.Background()
	course, err := mockRepo.GetCourseByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, course)
	assert.Equal(t, 1, course.ID)
}

func TestGetCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := course.NewMockCourseRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetCourses(gomock.Any()).Return([]*models.Course{{ID: 1}}, nil)

	ctx := context.Background()
	courses, err := mockRepo.GetCourses(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, courses)
	assert.Len(t, courses, 1)
	assert.Equal(t, 1, courses[0].ID)
}

func TestUpdateCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := course.NewMockCourseRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateCourse(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	course := &models.Course{ID: 1}

	err := mockRepo.UpdateCourse(ctx, course)
	assert.NoError(t, err)
}

func TestDeleteCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := course.NewMockCourseRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteCourse(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteCourse(ctx, 1)
	assert.NoError(t, err)
}
