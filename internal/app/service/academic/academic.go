package academic

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/academic"
	"go-tsukamoto/internal/app/models"
	"time"
)

func (s *academicService) CreateAcademic(ctx context.Context, req *academic.CreateAcademicRequest) (*academic.AcademicResponse, error) {
	// Validate if UserID exists
	user, err := s.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	academicModel := &models.Academic{
		UserID:          req.UserID,
		Ipk:             req.Ipk,
		Ips:             req.Ips,
		RepeatedCourses: req.RepeatedCourses,
		Semester:        req.Semester,
		Year:            req.Year,
		PredicateID:     req.PredicateID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := s.repo.CreateAcademic(ctx, academicModel); err != nil {
		return nil, err
	}
	return &academic.AcademicResponse{
		ID:              academicModel.ID,
		UserID:          academicModel.UserID,
		Ipk:             academicModel.Ipk,
		Ips:             academicModel.Ips,
		RepeatedCourses: academicModel.RepeatedCourses,
		Semester:        academicModel.Semester,
		Year:            academicModel.Year,
		PredicateID:     academicModel.PredicateID,
		CreatedAt:       academicModel.CreatedAt,
		UpdatedAt:       academicModel.UpdatedAt,
	}, nil
}

func (s *academicService) GetAcademicByID(ctx context.Context, id int) (*academic.AcademicResponse, error) {
	academicModel, err := s.repo.GetAcademicByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if academicModel == nil {
		return nil, errors.New("academic record not found")
	}
	return &academic.AcademicResponse{
		ID:              academicModel.ID,
		UserID:          academicModel.UserID,
		Ipk:             academicModel.Ipk,
		Ips:             academicModel.Ips,
		RepeatedCourses: academicModel.RepeatedCourses,
		Semester:        academicModel.Semester,
		Year:            academicModel.Year,
		PredicateID:     academicModel.PredicateID,
		CreatedAt:       academicModel.CreatedAt,
		UpdatedAt:       academicModel.UpdatedAt,
	}, nil
}

func (s *academicService) GetAcademicsByUserID(ctx context.Context, userID int) ([]*academic.AcademicResponse, error) {
	academicModels, err := s.repo.GetAcademicsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var academics []*academic.AcademicResponse
	for _, academicModel := range academicModels {
		academics = append(academics, &academic.AcademicResponse{
			ID:              academicModel.ID,
			UserID:          academicModel.UserID,
			Ipk:             academicModel.Ipk,
			Ips:             academicModel.Ips,
			RepeatedCourses: academicModel.RepeatedCourses,
			Semester:        academicModel.Semester,
			Year:            academicModel.Year,
			PredicateID:     academicModel.PredicateID,
			CreatedAt:       academicModel.CreatedAt,
			UpdatedAt:       academicModel.UpdatedAt,
		})
	}
	return academics, nil
}

func (s *academicService) UpdateAcademic(ctx context.Context, id int, req *academic.UpdateAcademicRequest) (*academic.AcademicResponse, error) {
	academicModel, err := s.repo.GetAcademicByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if academicModel == nil {
		return nil, errors.New("academic record not found")
	}

	// Validate if UserID exists
	user, err := s.userRepo.GetUserByID(ctx, academicModel.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	academicModel.Ipk = req.Ipk
	academicModel.Ips = req.Ips
	academicModel.RepeatedCourses = req.RepeatedCourses
	academicModel.Semester = req.Semester
	academicModel.Year = req.Year
	academicModel.PredicateID = req.PredicateID
	academicModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateAcademic(ctx, academicModel); err != nil {
		return nil, err
	}
	return &academic.AcademicResponse{
		ID:              academicModel.ID,
		UserID:          academicModel.UserID,
		Ipk:             academicModel.Ipk,
		Ips:             academicModel.Ips,
		RepeatedCourses: academicModel.RepeatedCourses,
		Semester:        academicModel.Semester,
		Year:            academicModel.Year,
		PredicateID:     academicModel.PredicateID,
		CreatedAt:       academicModel.CreatedAt,
		UpdatedAt:       academicModel.UpdatedAt,
	}, nil
}

func (s *academicService) DeleteAcademic(ctx context.Context, id int) error {
	return s.repo.DeleteAcademic(ctx, id)
}
