package fuzzy_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/fuzzy"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetStudentData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := fuzzy.NewMockFuzzyRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetStudentData(gomock.Any(), 1).Return(&models.Users{ID: 1}, nil)

	ctx := context.Background()
	student, err := mockRepo.GetStudentData(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, student)
	assert.Equal(t, 1, student.ID)
}

func TestGetStudentData_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := fuzzy.NewMockFuzzyRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetStudentData(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	student, err := mockRepo.GetStudentData(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, student)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
