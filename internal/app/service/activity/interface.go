package activity

import (
	"context"
	"go-tsukamoto/internal/app/dto/activity"
	repo "go-tsukamoto/internal/app/repository/activity"

	"gorm.io/gorm"
)

type activityService struct {
	repo repo.ActivityRepositoryInterface
}

func NewActivityService(repo repo.ActivityRepositoryInterface) ActivityService {
	return &activityService{repo: repo}
}

func NewService(db *gorm.DB) ActivityService {
	repository := repo.NewActivityRepository(db)
	return &activityService{repo: repository}
}

type ActivityService interface {
	CreateActivity(ctx context.Context, req *activity.CreateActivityRequest) (*activity.ActivityResponse, error)
	GetActivityByID(ctx context.Context, id int) (*activity.ActivityResponse, error)
	GetActivitiesByUserID(ctx context.Context, userID int) ([]*activity.ActivityResponse, error)
	UpdateActivity(ctx context.Context, id int, req *activity.UpdateActivityRequest) (*activity.ActivityResponse, error)
	DeleteActivity(ctx context.Context, id int) error
}
