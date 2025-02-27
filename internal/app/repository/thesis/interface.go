package thesis

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type ThesisRepositoryInterface interface {
	CreateThesis(ctx context.Context, thesis *models.Thesis) error
	GetThesisByID(ctx context.Context, id int) (*models.Thesis, error)
	GetThesesByUserID(ctx context.Context, userID int) ([]*models.Thesis, error)
	GetAllTheses(ctx context.Context) ([]*models.Thesis, error)
	UpdateThesis(ctx context.Context, thesis *models.Thesis) error
	DeleteThesis(ctx context.Context, id int) error
}

type thesisRepository struct {
	db *gorm.DB
}

func NewThesisRepository(db *gorm.DB) ThesisRepositoryInterface {
	return &thesisRepository{db: db}
}
