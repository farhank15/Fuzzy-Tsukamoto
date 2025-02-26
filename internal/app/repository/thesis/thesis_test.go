package thesis_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/thesis"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().CreateThesis(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	thesis := &models.Thesis{}

	err := mockRepo.CreateThesis(ctx, thesis)
	assert.NoError(t, err)
}

func TestGetThesisByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetThesisByID(gomock.Any(), 1).Return(&models.Thesis{ID: 1}, nil)

	ctx := context.Background()
	thesis, err := mockRepo.GetThesisByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, thesis)
	assert.Equal(t, 1, thesis.ID)
}

func TestGetThesesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetThesesByUserID(gomock.Any(), 1).Return([]*models.Thesis{{ID: 1}}, nil)

	ctx := context.Background()
	theses, err := mockRepo.GetThesesByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, theses)
	assert.Len(t, theses, 1)
	assert.Equal(t, 1, theses[0].ID)
}

func TestUpdateThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().UpdateThesis(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()
	thesis := &models.Thesis{ID: 1}

	err := mockRepo.UpdateThesis(ctx, thesis)
	assert.NoError(t, err)
}

func TestDeleteThesis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().DeleteThesis(gomock.Any(), 1).Return(nil)

	ctx := context.Background()

	err := mockRepo.DeleteThesis(ctx, 1)
	assert.NoError(t, err)
}

func TestGetByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByUserID(gomock.Any(), 1).Return(&models.Thesis{ID: 1}, nil)

	ctx := context.Background()
	thesis, err := mockRepo.GetByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, thesis)
	assert.Equal(t, 1, thesis.ID)
}

func TestGetByUserID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := thesis.NewMockThesisRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByUserID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	thesis, err := mockRepo.GetByUserID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, thesis)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
