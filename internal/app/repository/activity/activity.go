package activity

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

func (r *activityRepository) CreateActivity(ctx context.Context, activity *models.Activity) error {
	return r.db.WithContext(ctx).Create(activity).Error
}

func (r *activityRepository) GetActivityByID(ctx context.Context, id int) (*models.Activity, error) {
	var activity models.Activity
	if err := r.db.WithContext(ctx).First(&activity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepository) GetActivitiesByUserID(ctx context.Context, userID int) ([]*models.Activity, error) {
	var activities []*models.Activity
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) GetAllActivities(ctx context.Context) ([]*models.Activity, error) {
	var activities []*models.Activity
	if err := r.db.WithContext(ctx).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) UpdateActivity(ctx context.Context, activity *models.Activity) error {
	return r.db.WithContext(ctx).Save(activity).Error
}

func (r *activityRepository) DeleteActivity(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Activity{}, id).Error
}
