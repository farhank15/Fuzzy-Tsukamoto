package achievement

import (
	"context"
	"go-tsukamoto/internal/app/dto/achievement"
	repo "go-tsukamoto/internal/app/repository/achievement"

	"gorm.io/gorm"
)

type achievementService struct {
	repo repo.AchievementRepositoryInterface
}

func NewAchievementService(repo repo.AchievementRepositoryInterface) AchievementService {
	return &achievementService{repo: repo}
}

func NewService(db *gorm.DB) AchievementService {
	repository := repo.NewAchievementRepository(db)
	return &achievementService{repo: repository}
}

type AchievementService interface {
	CreateAchievement(ctx context.Context, req *achievement.CreateAchievementRequest) (*achievement.AchievementResponse, error)
	GetAchievementByID(ctx context.Context, id int) (*achievement.AchievementResponse, error)
	GetAchievementsByUserID(ctx context.Context, userID int) ([]*achievement.AchievementResponse, error)
	GetAllAchievements(ctx context.Context) ([]*achievement.AchievementResponse, error)
	UpdateAchievement(ctx context.Context, id int, req *achievement.UpdateAchievementRequest) (*achievement.AchievementResponse, error)
	DeleteAchievement(ctx context.Context, id int) error
}
