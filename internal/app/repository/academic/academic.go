package academic

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

func (r *academicRepository) CreateAcademic(ctx context.Context, academic *models.Academic) error {
	return r.db.WithContext(ctx).Create(academic).Error
}

func (r *academicRepository) GetAcademicByID(ctx context.Context, id int) (*models.Academic, error) {
	var academic models.Academic
	if err := r.db.WithContext(ctx).First(&academic, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &academic, nil
}

func (r *academicRepository) GetAcademicsByUserID(ctx context.Context, userID int) ([]*models.Academic, error) {
	var academics []*models.Academic
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&academics).Error; err != nil {
		return nil, err
	}
	return academics, nil
}

func (r *academicRepository) GetAllAcademics(ctx context.Context) ([]*models.Academic, error) {
	var academics []*models.Academic
	if err := r.db.WithContext(ctx).Find(&academics).Error; err != nil {
		return nil, err
	}
	return academics, nil
}

func (r *academicRepository) UpdateAcademic(ctx context.Context, academic *models.Academic) error {
	return r.db.WithContext(ctx).Save(academic).Error
}

func (r *academicRepository) DeleteAcademic(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Academic{}, id).Error
}

func (r *academicRepository) Update(ctx context.Context, academic *models.Academic) error {
	return r.db.WithContext(ctx).Save(academic).Error
}
