package thesis

import (
	"context"
	"go-tsukamoto/internal/app/dto/thesis"
	repo "go-tsukamoto/internal/app/repository/thesis"
	userRepo "go-tsukamoto/internal/app/repository/user"

	"gorm.io/gorm"
)

type thesisService struct {
	repo     repo.ThesisRepositoryInterface
	userRepo userRepo.UserRepositoryInterface
}

func NewThesisService(repo repo.ThesisRepositoryInterface, userRepo userRepo.UserRepositoryInterface) ThesisService {
	return &thesisService{repo: repo, userRepo: userRepo}
}

func NewService(db *gorm.DB) ThesisService {
	repository := repo.NewThesisRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	return &thesisService{repo: repository, userRepo: userRepository}
}

type ThesisService interface {
	CreateThesis(ctx context.Context, req *thesis.CreateThesisRequest) (*thesis.ThesisResponse, error)
	GetThesisByID(ctx context.Context, id int) (*thesis.ThesisResponse, error)
	GetThesesByUserID(ctx context.Context, userID int) ([]*thesis.ThesisResponse, error)
	GetAllTheses(ctx context.Context) ([]*thesis.ThesisResponse, error)
	UpdateThesis(ctx context.Context, id int, req *thesis.UpdateThesisRequest) (*thesis.ThesisResponse, error)
	DeleteThesis(ctx context.Context, id int) error
}
