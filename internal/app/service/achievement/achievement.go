package achievement

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/achievement"
	"go-tsukamoto/internal/app/models"
	"time"
)

func (s *achievementService) CreateAchievement(ctx context.Context, req *achievement.CreateAchievementRequest) (*achievement.AchievementResponse, error) {
	achievementModel := &models.Achievement{
		UserID:      req.UserID,
		Title:       req.Title,
		Certificate: req.Certificate,
		Rank:        req.Rank,
		Level:       models.Level(req.Level),
		Year:        req.Year,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.CreateAchievement(ctx, achievementModel); err != nil {
		return nil, err
	}
	return &achievement.AchievementResponse{
		ID:          achievementModel.ID,
		UserID:      achievementModel.UserID,
		Title:       achievementModel.Title,
		Certificate: achievementModel.Certificate,
		Rank:        achievementModel.Rank,
		Level:       string(achievementModel.Level),
		Year:        achievementModel.Year,
		CreatedAt:   achievementModel.CreatedAt,
		UpdatedAt:   achievementModel.UpdatedAt,
	}, nil
}

func (s *achievementService) GetAchievementByID(ctx context.Context, id int) (*achievement.AchievementResponse, error) {
	achievementModel, err := s.repo.GetAchievementByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if achievementModel == nil {
		return nil, errors.New("achievement not found")
	}
	return &achievement.AchievementResponse{
		ID:          achievementModel.ID,
		UserID:      achievementModel.UserID,
		Title:       achievementModel.Title,
		Certificate: achievementModel.Certificate,
		Rank:        achievementModel.Rank,
		Level:       string(achievementModel.Level),
		Year:        achievementModel.Year,
		CreatedAt:   achievementModel.CreatedAt,
		UpdatedAt:   achievementModel.UpdatedAt,
	}, nil
}

func (s *achievementService) GetAchievementsByUserID(ctx context.Context, userID int) ([]*achievement.AchievementResponse, error) {
	achievementModels, err := s.repo.GetAchievementsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var achievements []*achievement.AchievementResponse
	for _, achievementModel := range achievementModels {
		achievements = append(achievements, &achievement.AchievementResponse{
			ID:          achievementModel.ID,
			UserID:      achievementModel.UserID,
			Title:       achievementModel.Title,
			Certificate: achievementModel.Certificate,
			Rank:        achievementModel.Rank,
			Level:       string(achievementModel.Level),
			Year:        achievementModel.Year,
			CreatedAt:   achievementModel.CreatedAt,
			UpdatedAt:   achievementModel.UpdatedAt,
		})
	}
	return achievements, nil
}

func (s *achievementService) GetAllAchievements(ctx context.Context) ([]*achievement.AchievementResponse, error) {
	achievementModels, err := s.repo.GetAllAchievements(ctx)
	if err != nil {
		return nil, err
	}
	var achievements []*achievement.AchievementResponse
	for _, achievementModel := range achievementModels {
		achievements = append(achievements, &achievement.AchievementResponse{
			ID:          achievementModel.ID,
			UserID:      achievementModel.UserID,
			Title:       achievementModel.Title,
			Certificate: achievementModel.Certificate,
			Rank:        achievementModel.Rank,
			Level:       string(achievementModel.Level),
			Year:        achievementModel.Year,
			CreatedAt:   achievementModel.CreatedAt,
			UpdatedAt:   achievementModel.UpdatedAt,
		})
	}
	return achievements, nil
}

func (s *achievementService) UpdateAchievement(ctx context.Context, id int, req *achievement.UpdateAchievementRequest) (*achievement.AchievementResponse, error) {
	achievementModel, err := s.repo.GetAchievementByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if achievementModel == nil {
		return nil, errors.New("achievement not found")
	}
	achievementModel.Title = req.Title
	achievementModel.Certificate = req.Certificate
	achievementModel.Rank = req.Rank
	achievementModel.Level = models.Level(req.Level)
	achievementModel.Year = req.Year
	achievementModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateAchievement(ctx, achievementModel); err != nil {
		return nil, err
	}
	return &achievement.AchievementResponse{
		ID:          achievementModel.ID,
		UserID:      achievementModel.UserID,
		Title:       achievementModel.Title,
		Certificate: achievementModel.Certificate,
		Rank:        achievementModel.Rank,
		Level:       string(achievementModel.Level),
		Year:        achievementModel.Year,
		CreatedAt:   achievementModel.CreatedAt,
		UpdatedAt:   achievementModel.UpdatedAt,
	}, nil
}

func (s *achievementService) DeleteAchievement(ctx context.Context, id int) error {
	return s.repo.DeleteAchievement(ctx, id)
}
