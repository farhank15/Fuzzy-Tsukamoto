package academic

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type AcademicRepositoryInterface interface {
	CreateAcademic(ctx context.Context, academic *models.Academic) error
	GetByUserID(ctx context.Context, userID int) (*models.Academic, error)
	GetAcademicByID(ctx context.Context, id int) (*models.Academic, error)
	GetAcademicsByUserID(ctx context.Context, userID int) ([]*models.Academic, error)
	GetAllAcademics(ctx context.Context) ([]*models.Academic, error)
	UpdateAcademic(ctx context.Context, academic *models.Academic) error
	DeleteAcademic(ctx context.Context, id int) error
}

type academicRepository struct {
	db *gorm.DB
}

func NewAcademicRepository(db *gorm.DB) AcademicRepositoryInterface {
	return &academicRepository{db: db}
}
