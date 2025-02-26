package predicate

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type PredicateRepositoryInterface interface {
	GetPredicateByID(ctx context.Context, id int) (*models.Predicate, error)
	GetByName(ctx context.Context, name string) (*models.Predicate, error)
}

type predicateRepository struct {
	db *gorm.DB
}

func NewPredicateRepository(db *gorm.DB) PredicateRepositoryInterface {
	return &predicateRepository{db: db}
}
