package course

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/course"
	"go-tsukamoto/internal/app/models"
	"time"
)

func (s *courseService) CreateCourse(ctx context.Context, req *course.CreateCourseRequest) (*course.CourseResponse, error) {
	courseModel := &models.Course{
		Code:         req.Code,
		CourseName:   req.CourseName,
		CreditCourse: req.CreditCourse,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.repo.CreateCourse(ctx, courseModel); err != nil {
		return nil, err
	}
	return &course.CourseResponse{
		ID:           courseModel.ID,
		Code:         courseModel.Code,
		CourseName:   courseModel.CourseName,
		CreditCourse: courseModel.CreditCourse,
		CreatedAt:    courseModel.CreatedAt,
		UpdatedAt:    courseModel.UpdatedAt,
	}, nil
}

func (s *courseService) GetCourseByID(ctx context.Context, id int) (*course.CourseResponse, error) {
	courseModel, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if courseModel == nil {
		return nil, errors.New("course not found")
	}
	return &course.CourseResponse{
		ID:           courseModel.ID,
		Code:         courseModel.Code,
		CourseName:   courseModel.CourseName,
		CreditCourse: courseModel.CreditCourse,
		CreatedAt:    courseModel.CreatedAt,
		UpdatedAt:    courseModel.UpdatedAt,
	}, nil
}

func (s *courseService) GetCourses(ctx context.Context) ([]*course.CourseResponse, error) {
	courseModels, err := s.repo.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	courses := []*course.CourseResponse{}

	for _, courseModel := range courseModels {
		courses = append(courses, &course.CourseResponse{
			ID:           courseModel.ID,
			Code:         courseModel.Code,
			CourseName:   courseModel.CourseName,
			CreditCourse: courseModel.CreditCourse,
			CreatedAt:    courseModel.CreatedAt,
			UpdatedAt:    courseModel.UpdatedAt,
		})
	}
	return courses, nil
}

func (s *courseService) UpdateCourse(ctx context.Context, id int, req *course.UpdateCourseRequest) (*course.CourseResponse, error) {
	courseModel, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if courseModel == nil {
		return nil, errors.New("course not found")
	}

	courseModel.Code = req.Code
	courseModel.CourseName = req.CourseName
	courseModel.CreditCourse = req.CreditCourse
	courseModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateCourse(ctx, courseModel); err != nil {
		return nil, err
	}
	return &course.CourseResponse{
		ID:           courseModel.ID,
		Code:         courseModel.Code,
		CourseName:   courseModel.CourseName,
		CreditCourse: courseModel.CreditCourse,
		CreatedAt:    courseModel.CreatedAt,
		UpdatedAt:    courseModel.UpdatedAt,
	}, nil
}

func (s *courseService) DeleteCourse(ctx context.Context, id int) error {
	courseModel, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return err
	}
	if courseModel == nil {
		return errors.New("course not found")
	}
	return s.repo.DeleteCourse(ctx, id)
}

func (s *courseService) ImportCourses(ctx context.Context, reqs []course.CreateCourseRequest) error {
	for _, req := range reqs {
		courseModel := &models.Course{
			Code:         req.Code,
			CourseName:   req.CourseName,
			CreditCourse: req.CreditCourse,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if err := s.repo.CreateCourse(ctx, courseModel); err != nil {
			return err
		}
	}
	return nil
}
