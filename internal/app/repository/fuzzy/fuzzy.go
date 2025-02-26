package repository

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type FuzzyRepositoryInterface interface {
	GetStudentData(ctx context.Context, studentID int) (*models.Users, error)
}

type fuzzyRepository struct {
	db *gorm.DB
}

func NewFuzzyRepository(db *gorm.DB) FuzzyRepositoryInterface {
	return &fuzzyRepository{db: db}
}

func (r *fuzzyRepository) GetStudentData(ctx context.Context, studentID int) (*models.Users, error) {
	var student models.Users
	if err := r.db.WithContext(ctx).Where("id = ?", studentID).First(&student).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}
