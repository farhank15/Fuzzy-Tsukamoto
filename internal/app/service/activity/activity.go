package activity

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/activity"
	"go-tsukamoto/internal/app/models"
	"time"
)

func (s *activityService) CreateActivity(ctx context.Context, req *activity.CreateActivityRequest) (*activity.ActivityResponse, error) {
	// Validate if UserID exists
	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	activityModel := &models.Activity{
		UserID:       req.UserID,
		Organization: req.Organization,
		Year:         req.Year,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.repo.CreateActivity(ctx, activityModel); err != nil {
		return nil, err
	}
	return &activity.ActivityResponse{
		ID:           activityModel.ID,
		UserID:       activityModel.UserID,
		Organization: activityModel.Organization,
		Year:         activityModel.Year,
		CreatedAt:    activityModel.CreatedAt,
		UpdatedAt:    activityModel.UpdatedAt,
	}, nil
}

func (s *activityService) GetActivityByID(ctx context.Context, id int) (*activity.ActivityResponse, error) {
	activityModel, err := s.repo.GetActivityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if activityModel == nil {
		return nil, errors.New("activity not found")
	}
	return &activity.ActivityResponse{
		ID:           activityModel.ID,
		UserID:       activityModel.UserID,
		Organization: activityModel.Organization,
		Year:         activityModel.Year,
		CreatedAt:    activityModel.CreatedAt,
		UpdatedAt:    activityModel.UpdatedAt,
	}, nil
}

func (s *activityService) GetActivitiesByUserID(ctx context.Context, userID int) ([]*activity.ActivityResponse, error) {
	activityModels, err := s.repo.GetActivitiesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var activities []*activity.ActivityResponse
	for _, activityModel := range activityModels {
		activities = append(activities, &activity.ActivityResponse{
			ID:           activityModel.ID,
			UserID:       activityModel.UserID,
			Organization: activityModel.Organization,
			Year:         activityModel.Year,
			CreatedAt:    activityModel.CreatedAt,
			UpdatedAt:    activityModel.UpdatedAt,
		})
	}
	return activities, nil
}

func (s *activityService) UpdateActivity(ctx context.Context, id int, req *activity.UpdateActivityRequest) (*activity.ActivityResponse, error) {
	activityModel, err := s.repo.GetActivityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if activityModel == nil {
		return nil, errors.New("activity not found")
	}

	// Validate if UserID exists
	user, err := s.repo.GetUserByID(ctx, activityModel.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	activityModel.Organization = req.Organization
	activityModel.Year = req.Year
	activityModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateActivity(ctx, activityModel); err != nil {
		return nil, err
	}
	return &activity.ActivityResponse{
		ID:           activityModel.ID,
		UserID:       activityModel.UserID,
		Organization: activityModel.Organization,
		Year:         activityModel.Year,
		CreatedAt:    activityModel.CreatedAt,
		UpdatedAt:    activityModel.UpdatedAt,
	}, nil
}

func (s *activityService) DeleteActivity(ctx context.Context, id int) error {
	return s.repo.DeleteActivity(ctx, id)
}
