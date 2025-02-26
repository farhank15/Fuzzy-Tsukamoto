package activity

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type ActivityRepositoryInterface interface {
	CreateActivity(ctx context.Context, activity *models.Activity) error
	GetByStudentID(ctx context.Context, studentID int) ([]models.Activity, error)
	GetActivityByID(ctx context.Context, id int) (*models.Activity, error)
	GetActivitiesByUserID(ctx context.Context, userID int) ([]*models.Activity, error)
	UpdateActivity(ctx context.Context, activity *models.Activity) error
	DeleteActivity(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*models.Users, error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepositoryInterface {
	return &activityRepository{db: db}
}
