package achievement

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

func (r *achievementRepository) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
	return r.db.WithContext(ctx).Create(achievement).Error
}

func (r *achievementRepository) GetAchievementByID(ctx context.Context, id int) (*models.Achievement, error) {
	var achievement models.Achievement
	if err := r.db.WithContext(ctx).First(&achievement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &achievement, nil
}

func (r *achievementRepository) GetAchievementsByUserID(ctx context.Context, userID int) ([]*models.Achievement, error) {
	var achievements []*models.Achievement
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&achievements).Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (r *achievementRepository) UpdateAchievement(ctx context.Context, achievement *models.Achievement) error {
	return r.db.WithContext(ctx).Save(achievement).Error
}

func (r *achievementRepository) DeleteAchievement(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Achievement{}, id).Error
}

func (r *achievementRepository) GetByStudentID(ctx context.Context, studentID int) ([]models.Achievement, error) {
	var achievements []models.Achievement
	if err := r.db.WithContext(ctx).Where("user_id = ?", studentID).Find(&achievements).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return achievements, nil
}
