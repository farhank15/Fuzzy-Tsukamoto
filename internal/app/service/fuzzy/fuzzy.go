package fuzzy

import (
	"context"
	"fmt"
	dto "go-tsukamoto/internal/app/dto/fuzzy"
	"go-tsukamoto/internal/app/models"
	academicRepo "go-tsukamoto/internal/app/repository/academic"
	achievementRepo "go-tsukamoto/internal/app/repository/achievement"
	activityRepo "go-tsukamoto/internal/app/repository/activity"
	predicateRepo "go-tsukamoto/internal/app/repository/predicate"
	thesisRepo "go-tsukamoto/internal/app/repository/thesis"
	"go-tsukamoto/internal/modules/inferensia"

	log "github.com/sirupsen/logrus"
)

type FuzzyService struct {
	academicRepo    academicRepo.AcademicRepositoryInterface
	thesisRepo      thesisRepo.ThesisRepositoryInterface
	achievementRepo achievementRepo.AchievementRepositoryInterface
	activityRepo    activityRepo.ActivityRepositoryInterface
	predicateRepo   predicateRepo.PredicateRepositoryInterface
}

func (s *FuzzyService) CalculateFuzzy(ctx context.Context, studentID int) (*dto.FuzzyResponseDTO, error) {
	academics, err := s.academicRepo.GetAcademicsByUserID(ctx, studentID)
	if err != nil {
		return nil, fmt.Errorf("error getting academic data: %v", err)
	}
	if len(academics) == 0 {
		return nil, fmt.Errorf("academic data not found for student ID: %d", studentID)
	}
	academic := academics[0]

	theses, err := s.thesisRepo.GetThesesByUserID(ctx, studentID)
	if err != nil {
		log.Warnf("error getting thesis data: %v", err)
	}
	var thesis *models.Thesis
	if len(theses) > 0 {
		thesis = theses[0]
	} else {
		log.Warnf("thesis data not found for student ID: %d", studentID)
		thesis = &models.Thesis{}
	}

	achievements, err := s.achievementRepo.GetAchievementsByUserID(ctx, studentID)
	if err != nil {
		log.Warnf("error getting achievement data: %v", err)
	}

	activities, err := s.activityRepo.GetActivitiesByUserID(ctx, studentID)
	if err != nil {
		log.Warnf("error getting activity data: %v", err)
	}

	// 2. Persiapkan data untuk fuzzy
	bestAchievement := getBestAchievement(achievements)
	activityCount := len(activities)

	// Default values if bestAchievement is nil
	bestAchievementLevel := ""
	bestAchievementRank := 0
	if bestAchievement != nil {
		bestAchievementLevel = string(bestAchievement.Level)
		bestAchievementRank = bestAchievement.Rank
	}

	thesisImpactFactor := calculateThesisImpact(*thesis)

	// 3. Jalankan proses fuzzy menggunakan package yang sudah ada
	hasilPredicate := inferensia.TsukamotoInference(
		academic.Ipk,             // IPK mahasiswa
		academic.Semester,        // Semester yang telah ditempuh
		academic.RepeatedCourses, // Jumlah mata kuliah mengulang
		bestAchievementRank,      // Ranking prestasi terbaik
		bestAchievementLevel,     // Level prestasi (internasional/nasional/internal)
		thesisImpactFactor,       // Impact factor skripsi
		thesis.Level,             // Level publikasi skripsi
		activityCount,            // Jumlah aktivitas organisasi
	)
	// 4. Update predicateID di tabel academic
	predicate, err := s.predicateRepo.GetByName(ctx, hasilPredicate)
	if err != nil {
		return nil, fmt.Errorf("error getting predicate: %v", err)
	}

	// Update academic dengan predicate baru
	academic.PredicateID = predicate.ID
	if err := s.academicRepo.UpdateAcademic(ctx, academic); err != nil {
		return nil, fmt.Errorf("error updating academic predicate: %v", err)
	}

	// 5. Buat response
	response := &dto.FuzzyResponseDTO{
		StudentID:       studentID,
		IPK:             academic.Ipk,
		Semester:        academic.Semester,
		MataKuliahUlang: academic.RepeatedCourses,
		PrestasiLevel:   bestAchievementLevel,
		PrestasiRank:    bestAchievementRank,
		SkripsiLevel:    thesis.Level,
		SkripsiImpact:   thesisImpactFactor,
		JumlahAktivitas: activityCount,
		HasilPredicate:  hasilPredicate,
	}

	return response, nil
}

func getBestAchievement(achievements []*models.Achievement) *models.Achievement {
	if len(achievements) == 0 {
		return nil
	}

	best := achievements[0]
	for _, achievement := range achievements[1:] {
		// Prioritaskan level yang lebih tinggi
		if getLevelPriority(achievement.Level) > getLevelPriority(best.Level) {
			best = achievement
			continue
		}

		// Jika level sama, pilih ranking yang lebih kecil (lebih baik)
		if achievement.Level == best.Level && achievement.Rank < best.Rank {
			best = achievement
		}
	}
	return best
}

// Fungsi helper untuk menentukan prioritas level
func getLevelPriority(level models.Level) int {
	switch level {
	case models.LevelInternasional:
		return 3
	case models.LevelNasional:
		return 2
	case models.LevelInternal:
		return 1
	default:
		return 0
	}
}

// Fungsi helper untuk menghitung impact factor skripsi
func calculateThesisImpact(thesis models.Thesis) float64 {
	// Implementasi sesuai dengan kriteria penilaian skripsi
	// Contoh sederhana:
	switch thesis.Level {
	case "internasional":
		return 5.0
	case "nasional":
		return 3.0
	default: // internal
		return 1.0
	}
}
