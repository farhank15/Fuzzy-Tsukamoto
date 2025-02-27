package academic_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/academic"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateAcademic(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	academic := &models.Academic{}

	err := mockRepo.CreateAcademic(ctx, academic)
	assert.NoError(t, err)
}

func TestGetAcademicByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetAcademicByID(gomock.Any(), 1).Return(&models.Academic{ID: 1}, nil)

	ctx := context.Background()
	academic, err := mockRepo.GetAcademicByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, academic)
	assert.Equal(t, 1, academic.ID)
}

func TestGetAcademicsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetAcademicsByUserID(gomock.Any(), 1).Return([]*models.Academic{{ID: 1}}, nil)

	ctx := context.Background()
	academics, err := mockRepo.GetAcademicsByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, academics)
	assert.Len(t, academics, 1)
	assert.Equal(t, 1, academics[0].ID)
}

func TestGetAllAcademics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetAllAcademics(gomock.Any()).Return([]*models.Academic{{ID: 1}, {ID: 2}}, nil)

	ctx := context.Background()
	academics, err := mockRepo.GetAllAcademics(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, academics)
	assert.Len(t, academics, 2)
	assert.Equal(t, 1, academics[0].ID)
	assert.Equal(t, 2, academics[1].ID)
}

func TestUpdateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateAcademic(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	academic := &models.Academic{ID: 1}

	err := mockRepo.UpdateAcademic(ctx, academic)
	assert.NoError(t, err)
}

func TestDeleteAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteAcademic(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteAcademic(ctx, 1)
	assert.NoError(t, err)
}

func TestGetByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByUserID(gomock.Any(), 1).Return(&models.Academic{ID: 1}, nil)

	ctx := context.Background()
	academic, err := mockRepo.GetByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, academic)
	assert.Equal(t, 1, academic.ID)
}

func TestGetByUserID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := academic.NewMockAcademicRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByUserID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	academic, err := mockRepo.GetByUserID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, academic)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
