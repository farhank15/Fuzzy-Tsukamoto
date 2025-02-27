package thesis

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

func (r *thesisRepository) CreateThesis(ctx context.Context, thesis *models.Thesis) error {
	return r.db.WithContext(ctx).Create(thesis).Error
}

func (r *thesisRepository) GetThesisByID(ctx context.Context, id int) (*models.Thesis, error) {
	var thesis models.Thesis
	if err := r.db.WithContext(ctx).First(&thesis, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &thesis, nil
}

func (r *thesisRepository) GetThesesByUserID(ctx context.Context, userID int) ([]*models.Thesis, error) {
	var theses []*models.Thesis
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&theses).Error; err != nil {
		return nil, err
	}
	return theses, nil
}

func (r *thesisRepository) GetAllTheses(ctx context.Context) ([]*models.Thesis, error) {
	var theses []*models.Thesis
	if err := r.db.WithContext(ctx).Find(&theses).Error; err != nil {
		return nil, err
	}
	return theses, nil
}

func (r *thesisRepository) UpdateThesis(ctx context.Context, thesis *models.Thesis) error {
	return r.db.WithContext(ctx).Save(thesis).Error
}

func (r *thesisRepository) DeleteThesis(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Thesis{}, id).Error
}
