package thesis

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/thesis"
	"go-tsukamoto/internal/app/models"
	"time"
)

func (s *thesisService) CreateThesis(ctx context.Context, req *thesis.CreateThesisRequest) (*thesis.ThesisResponse, error) {
	// Validate if UserID exists
	user, err := s.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	thesisModel := &models.Thesis{
		UserID:    req.UserID,
		Title:     req.Title,
		Year:      req.Year,
		Semester:  req.Semester,
		Value:     req.Value,
		Level:     req.Level,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.CreateThesis(ctx, thesisModel); err != nil {
		return nil, err
	}
	return &thesis.ThesisResponse{
		ID:        thesisModel.ID,
		UserID:    thesisModel.UserID,
		Title:     thesisModel.Title,
		Year:      thesisModel.Year,
		Semester:  thesisModel.Semester,
		Value:     thesisModel.Value,
		Level:     thesisModel.Level,
		CreatedAt: thesisModel.CreatedAt,
		UpdatedAt: thesisModel.UpdatedAt,
	}, nil
}

func (s *thesisService) GetThesisByID(ctx context.Context, id int) (*thesis.ThesisResponse, error) {
	thesisModel, err := s.repo.GetThesisByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if thesisModel == nil {
		return nil, errors.New("thesis not found")
	}
	return &thesis.ThesisResponse{
		ID:        thesisModel.ID,
		UserID:    thesisModel.UserID,
		Title:     thesisModel.Title,
		Year:      thesisModel.Year,
		Semester:  thesisModel.Semester,
		Value:     thesisModel.Value,
		Level:     thesisModel.Level,
		CreatedAt: thesisModel.CreatedAt,
		UpdatedAt: thesisModel.UpdatedAt,
	}, nil
}

func (s *thesisService) GetThesesByUserID(ctx context.Context, userID int) ([]*thesis.ThesisResponse, error) {
	thesisModels, err := s.repo.GetThesesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var theses []*thesis.ThesisResponse
	for _, thesisModel := range thesisModels {
		theses = append(theses, &thesis.ThesisResponse{
			ID:        thesisModel.ID,
			UserID:    thesisModel.UserID,
			Title:     thesisModel.Title,
			Year:      thesisModel.Year,
			Semester:  thesisModel.Semester,
			Value:     thesisModel.Value,
			Level:     thesisModel.Level,
			CreatedAt: thesisModel.CreatedAt,
			UpdatedAt: thesisModel.UpdatedAt,
		})
	}
	return theses, nil
}

func (s *thesisService) UpdateThesis(ctx context.Context, id int, req *thesis.UpdateThesisRequest) (*thesis.ThesisResponse, error) {
	thesisModel, err := s.repo.GetThesisByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if thesisModel == nil {
		return nil, errors.New("thesis not found")
	}

	// Validate if UserID exists
	user, err := s.userRepo.GetUserByID(ctx, thesisModel.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	thesisModel.Title = req.Title
	thesisModel.Year = req.Year
	thesisModel.Semester = req.Semester
	thesisModel.Value = req.Value
	thesisModel.Level = req.Level
	thesisModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateThesis(ctx, thesisModel); err != nil {
		return nil, err
	}
	return &thesis.ThesisResponse{
		ID:        thesisModel.ID,
		UserID:    thesisModel.UserID,
		Title:     thesisModel.Title,
		Year:      thesisModel.Year,
		Semester:  thesisModel.Semester,
		Value:     thesisModel.Value,
		Level:     thesisModel.Level,
		CreatedAt: thesisModel.CreatedAt,
		UpdatedAt: thesisModel.UpdatedAt,
	}, nil
}

func (s *thesisService) DeleteThesis(ctx context.Context, id int) error {
	return s.repo.DeleteThesis(ctx, id)
}
