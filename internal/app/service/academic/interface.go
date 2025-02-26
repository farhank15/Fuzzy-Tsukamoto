package academic

import (
	"context"
	"go-tsukamoto/internal/app/dto/academic"
	repo "go-tsukamoto/internal/app/repository/academic"
	predicateRepo "go-tsukamoto/internal/app/repository/predicate"
	userRepo "go-tsukamoto/internal/app/repository/user"

	"gorm.io/gorm"
)

type academicService struct {
	repo          repo.AcademicRepositoryInterface
	userRepo      userRepo.UserRepositoryInterface
	predicateRepo predicateRepo.PredicateRepositoryInterface
}

func NewAcademicService(repo repo.AcademicRepositoryInterface, userRepo userRepo.UserRepositoryInterface, predicateRepo predicateRepo.PredicateRepositoryInterface) AcademicService {
	return &academicService{repo: repo, userRepo: userRepo, predicateRepo: predicateRepo}
}

func NewService(db *gorm.DB) AcademicService {
	repository := repo.NewAcademicRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	predicateRepository := predicateRepo.NewPredicateRepository(db)
	return &academicService{repo: repository, userRepo: userRepository, predicateRepo: predicateRepository}
}

type AcademicService interface {
	CreateAcademic(ctx context.Context, req *academic.CreateAcademicRequest) (*academic.AcademicResponse, error)
	GetAcademicByID(ctx context.Context, id int) (*academic.AcademicResponse, error)
	GetAcademicsByUserID(ctx context.Context, userID int) ([]*academic.AcademicResponse, error)
	UpdateAcademic(ctx context.Context, id int, req *academic.UpdateAcademicRequest) (*academic.AcademicResponse, error)
	DeleteAcademic(ctx context.Context, id int) error
}
