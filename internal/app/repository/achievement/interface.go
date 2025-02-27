package achievement

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type AchievementRepositoryInterface interface {
	CreateAchievement(ctx context.Context, achievement *models.Achievement) error
	GetAchievementByID(ctx context.Context, id int) (*models.Achievement, error)
	GetAchievementsByUserID(ctx context.Context, userID int) ([]*models.Achievement, error)
	GetAllAchievements(ctx context.Context) ([]*models.Achievement, error)
	UpdateAchievement(ctx context.Context, achievement *models.Achievement) error
	DeleteAchievement(ctx context.Context, id int) error
}

type achievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository(db *gorm.DB) AchievementRepositoryInterface {
	return &achievementRepository{db: db}
}
