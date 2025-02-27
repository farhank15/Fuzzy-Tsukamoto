package predicate_test

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/internal/app/repository/predicate"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetPredicateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetPredicateByID(gomock.Any(), 1).Return(&models.Predicate{ID: 1}, nil)

	ctx := context.Background()
	predicate, err := mockRepo.GetPredicateByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, predicate)
	assert.Equal(t, 1, predicate.ID)
}

func TestGetPredicateByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetPredicateByID(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	predicate, err := mockRepo.GetPredicateByID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, predicate)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestGetPredicateByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetPredicateByID(gomock.Any(), 1).Return(nil, errors.New("unexpected error"))

	ctx := context.Background()
	predicate, err := mockRepo.GetPredicateByID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, predicate)
	assert.Equal(t, "unexpected error", err.Error())
}

func TestGetByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByName(gomock.Any(), "example").Return(&models.Predicate{ID: 1, Name: "example"}, nil)

	ctx := context.Background()
	predicate, err := mockRepo.GetByName(ctx, "example")
	assert.NoError(t, err)
	assert.NotNil(t, predicate)
	assert.Equal(t, "example", predicate.Name)
}

func TestGetByName_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByName(gomock.Any(), "example").Return(nil, gorm.ErrRecordNotFound)

	ctx := context.Background()
	predicate, err := mockRepo.GetByName(ctx, "example")
	assert.Error(t, err)
	assert.Nil(t, predicate)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestGetByName_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := predicate.NewMockPredicateRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetByName(gomock.Any(), "example").Return(nil, errors.New("unexpected error"))

	ctx := context.Background()
	predicate, err := mockRepo.GetByName(ctx, "example")
	assert.Error(t, err)
	assert.Nil(t, predicate)
	assert.Equal(t, "unexpected error", err.Error())
}
