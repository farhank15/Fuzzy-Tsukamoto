package predicate

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

func (r *predicateRepository) GetPredicateByID(ctx context.Context, id int) (*models.Predicate, error) {
	var predicate models.Predicate
	if err := r.db.WithContext(ctx).First(&predicate, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &predicate, nil
}

func (r *predicateRepository) GetByName(ctx context.Context, name string) (*models.Predicate, error) {
	var predicate models.Predicate
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&predicate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &predicate, nil
}
