package fuzzy

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-tsukamoto/internal/app/models"
	mockAcademicRepo "go-tsukamoto/internal/app/repository/academic"
	mockAchievementRepo "go-tsukamoto/internal/app/repository/achievement"
	mockActivityRepo "go-tsukamoto/internal/app/repository/activity"
	mockPredicateRepo "go-tsukamoto/internal/app/repository/predicate"
	mockThesisRepo "go-tsukamoto/internal/app/repository/thesis"
)

func TestCalculateFuzzy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAcademicRepo := mockAcademicRepo.NewMockAcademicRepositoryInterface(ctrl)
	mockThesisRepo := mockThesisRepo.NewMockThesisRepositoryInterface(ctrl)
	mockAchievementRepo := mockAchievementRepo.NewMockAchievementRepositoryInterface(ctrl)
	mockActivityRepo := mockActivityRepo.NewMockActivityRepositoryInterface(ctrl)
	mockPredicateRepo := mockPredicateRepo.NewMockPredicateRepositoryInterface(ctrl)

	fuzzyService := &FuzzyService{
		academicRepo:    mockAcademicRepo,
		thesisRepo:      mockThesisRepo,
		achievementRepo: mockAchievementRepo,
		activityRepo:    mockActivityRepo,
		predicateRepo:   mockPredicateRepo,
	}

	ctx := context.Background()
	studentID := 1

	t.Run("Success", func(t *testing.T) {
		// Mock data
		academics := []*models.Academic{
			{
				ID:              1,
				UserID:          studentID,
				Ipk:             3.75,
				Semester:        8,
				RepeatedCourses: 1,
				PredicateID:     0,
			},
		}

		theses := []*models.Thesis{
			{
				ID:     1,
				UserID: studentID,
				Level:  "nasional",
			},
		}

		achievements := []*models.Achievement{
			{
				ID:     1,
				UserID: studentID,
				Level:  models.LevelNasional,
				Rank:   2,
			},
		}

		activities := []*models.Activity{
			{
				ID:     1,
				UserID: studentID,
			},
			{
				ID:     2,
				UserID: studentID,
			},
		}

		predicate := &models.Predicate{
			ID:   1,
			Name: "Sangat Memuaskan",
		}

		// Set expectations
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(academics, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, studentID).Return(theses, nil)
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, studentID).Return(achievements, nil)
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, studentID).Return(activities, nil)
		mockPredicateRepo.EXPECT().GetByName(ctx, gomock.Any()).Return(predicate, nil)
		mockAcademicRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).Return(nil)

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, studentID, result.StudentID)
		assert.Equal(t, 3.75, result.IPK)
		assert.Equal(t, 8, result.Semester)
		assert.Equal(t, 1, result.MataKuliahUlang)
		assert.Equal(t, "nasional", result.PrestasiLevel)
		assert.Equal(t, 2, result.PrestasiRank)
		assert.Equal(t, "nasional", result.SkripsiLevel)
		assert.Equal(t, 3.0, result.SkripsiImpact)
		assert.Equal(t, 2, result.JumlahAktivitas)
		assert.NotEmpty(t, result.HasilPredicate)
	})

	t.Run("No Academic Data", func(t *testing.T) {
		// Empty academics array
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return([]*models.Academic{}, nil)

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "academic data not found")
	})

	t.Run("Academic Repository Error", func(t *testing.T) {
		// Error from academic repository
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(nil, errors.New("database error"))

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error getting academic data")
	})

	t.Run("No Thesis Data", func(t *testing.T) {
		// Mock data
		academics := []*models.Academic{
			{
				ID:              1,
				UserID:          studentID,
				Ipk:             3.75,
				Semester:        8,
				RepeatedCourses: 1,
				PredicateID:     0,
			},
		}

		achievements := []*models.Achievement{
			{
				ID:     1,
				UserID: studentID,
				Level:  models.LevelNasional,
				Rank:   2,
			},
		}

		activities := []*models.Activity{
			{
				ID:     1,
				UserID: studentID,
			},
		}

		predicate := &models.Predicate{
			ID:   1,
			Name: "Sangat Memuaskan",
		}

		// Set expectations
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(academics, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, studentID).Return([]*models.Thesis{}, nil) // Empty thesis
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, studentID).Return(achievements, nil)
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, studentID).Return(activities, nil)
		mockPredicateRepo.EXPECT().GetByName(ctx, gomock.Any()).Return(predicate, nil)
		mockAcademicRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).Return(nil)

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "", result.SkripsiLevel)   // Empty thesis level
		assert.Equal(t, 1.0, result.SkripsiImpact) // Default impact factor
	})

	t.Run("No Achievement Data", func(t *testing.T) {
		// Mock data
		academics := []*models.Academic{
			{
				ID:              1,
				UserID:          studentID,
				Ipk:             3.75,
				Semester:        8,
				RepeatedCourses: 1,
				PredicateID:     0,
			},
		}

		theses := []*models.Thesis{
			{
				ID:     1,
				UserID: studentID,
				Level:  "nasional",
			},
		}

		activities := []*models.Activity{
			{
				ID:     1,
				UserID: studentID,
			},
		}

		predicate := &models.Predicate{
			ID:   1,
			Name: "Sangat Memuaskan",
		}

		// Set expectations - Urutan sangat penting!
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(academics, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, studentID).Return(theses, nil)
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, studentID).Return([]*models.Achievement{}, nil) // Empty achievements
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, studentID).Return(activities, nil)
		mockPredicateRepo.EXPECT().GetByName(ctx, gomock.Any()).Return(predicate, nil)
		mockAcademicRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).Return(nil)

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "", result.PrestasiLevel) // Empty achievement level
		assert.Equal(t, 0, result.PrestasiRank)   // Default achievement rank
	})

	t.Run("Predicate Repository Error", func(t *testing.T) {
		// Mock data
		academics := []*models.Academic{
			{
				ID:              1,
				UserID:          studentID,
				Ipk:             3.75,
				Semester:        8,
				RepeatedCourses: 1,
				PredicateID:     0,
			},
		}

		theses := []*models.Thesis{
			{
				ID:     1,
				UserID: studentID,
				Level:  "nasional",
			},
		}

		achievements := []*models.Achievement{
			{
				ID:     1,
				UserID: studentID,
				Level:  models.LevelNasional,
				Rank:   2,
			},
		}

		activities := []*models.Activity{
			{
				ID:     1,
				UserID: studentID,
			},
		}

		// Set expectations
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(academics, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, studentID).Return(theses, nil)
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, studentID).Return(achievements, nil)
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, studentID).Return(activities, nil)
		mockPredicateRepo.EXPECT().GetByName(ctx, gomock.Any()).Return(nil, errors.New("predicate not found"))

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error getting predicate")
	})

	t.Run("Update Academic Error", func(t *testing.T) {
		// Mock data
		academics := []*models.Academic{
			{
				ID:              1,
				UserID:          studentID,
				Ipk:             3.75,
				Semester:        8,
				RepeatedCourses: 1,
				PredicateID:     0,
			},
		}

		theses := []*models.Thesis{
			{
				ID:     1,
				UserID: studentID,
				Level:  "nasional",
			},
		}

		achievements := []*models.Achievement{
			{
				ID:     1,
				UserID: studentID,
				Level:  models.LevelNasional,
				Rank:   2,
			},
		}

		activities := []*models.Activity{
			{
				ID:     1,
				UserID: studentID,
			},
		}

		predicate := &models.Predicate{
			ID:   1,
			Name: "Sangat Memuaskan",
		}

		// Set expectations
		mockAcademicRepo.EXPECT().GetAcademicsByUserID(ctx, studentID).Return(academics, nil)
		mockThesisRepo.EXPECT().GetThesesByUserID(ctx, studentID).Return(theses, nil)
		mockAchievementRepo.EXPECT().GetAchievementsByUserID(ctx, studentID).Return(achievements, nil)
		mockActivityRepo.EXPECT().GetActivitiesByUserID(ctx, studentID).Return(activities, nil)
		mockPredicateRepo.EXPECT().GetByName(ctx, gomock.Any()).Return(predicate, nil)
		mockAcademicRepo.EXPECT().UpdateAcademic(ctx, gomock.Any()).Return(errors.New("update error"))

		// Call the service
		result, err := fuzzyService.CalculateFuzzy(ctx, studentID)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error updating academic predicate")
	})
}

func TestGetBestAchievement(t *testing.T) {
	t.Run("Empty Achievements", func(t *testing.T) {
		achievements := []*models.Achievement{}
		best := getBestAchievement(achievements)
		assert.Nil(t, best)
	})

	t.Run("Single Achievement", func(t *testing.T) {
		achievements := []*models.Achievement{
			{
				ID:    1,
				Level: models.LevelNasional,
				Rank:  2,
			},
		}
		best := getBestAchievement(achievements)
		assert.Equal(t, 1, best.ID)
		assert.Equal(t, models.LevelNasional, best.Level)
		assert.Equal(t, 2, best.Rank)
	})

	t.Run("Multiple Achievements - Different Levels", func(t *testing.T) {
		achievements := []*models.Achievement{
			{
				ID:    1,
				Level: models.LevelNasional,
				Rank:  2,
			},
			{
				ID:    2,
				Level: models.LevelInternasional,
				Rank:  5,
			},
			{
				ID:    3,
				Level: models.LevelInternal,
				Rank:  1,
			},
		}
		best := getBestAchievement(achievements)
		assert.Equal(t, 2, best.ID) // International level has highest priority
		assert.Equal(t, models.LevelInternasional, best.Level)
	})

	t.Run("Multiple Achievements - Same Level", func(t *testing.T) {
		achievements := []*models.Achievement{
			{
				ID:    1,
				Level: models.LevelNasional,
				Rank:  2,
			},
			{
				ID:    2,
				Level: models.LevelNasional,
				Rank:  1, // Better rank
			},
			{
				ID:    3,
				Level: models.LevelNasional,
				Rank:  3,
			},
		}
		best := getBestAchievement(achievements)
		assert.Equal(t, 2, best.ID) // Rank 1 is better
		assert.Equal(t, models.LevelNasional, best.Level)
		assert.Equal(t, 1, best.Rank)
	})
}

func TestGetLevelPriority(t *testing.T) {
	tests := []struct {
		name     string
		level    models.Level
		expected int
	}{
		{"International", models.LevelInternasional, 3},
		{"National", models.LevelNasional, 2},
		{"Internal", models.LevelInternal, 1},
		{"Unknown", "unknown", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priority := getLevelPriority(tt.level)
			assert.Equal(t, tt.expected, priority)
		})
	}
}

func TestCalculateThesisImpact(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected float64
	}{
		{"International", "internasional", 5.0},
		{"National", "nasional", 3.0},
		{"Internal", "internal", 1.0},
		{"Default", "", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thesis := models.Thesis{Level: tt.level}
			impact := calculateThesisImpact(thesis)
			assert.Equal(t, tt.expected, impact)
		})
	}
}
