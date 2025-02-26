package service

import (
	"context"
	dto "go-tsukamoto/internal/app/dto/fuzzy"
	academicRepo "go-tsukamoto/internal/app/repository/academic"
	achievementRepo "go-tsukamoto/internal/app/repository/achievement"
	activityRepo "go-tsukamoto/internal/app/repository/activity"
	predicateRepo "go-tsukamoto/internal/app/repository/predicate"
	thesisRepo "go-tsukamoto/internal/app/repository/thesis"

	"gorm.io/gorm"
)

func NewService(db *gorm.DB) FuzzyServiceInterface {
	return &FuzzyService{
		academicRepo:    academicRepo.NewAcademicRepository(db),
		thesisRepo:      thesisRepo.NewThesisRepository(db),
		achievementRepo: achievementRepo.NewAchievementRepository(db),
		activityRepo:    activityRepo.NewActivityRepository(db),
		predicateRepo:   predicateRepo.NewPredicateRepository(db),
	}
}

type FuzzyServiceInterface interface {
	CalculateFuzzy(ctx context.Context, studentID int) (*dto.FuzzyResponseDTO, error)
}
