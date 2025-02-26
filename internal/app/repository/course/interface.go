package course

import (
	"context"
	"go-tsukamoto/internal/app/models"
)

type CourseRepositoryInterface interface {
	CreateCourse(ctx context.Context, course *models.Course) error
	GetCourseByID(ctx context.Context, id int) (*models.Course, error)
	GetCourses(ctx context.Context) ([]*models.Course, error)
	UpdateCourse(ctx context.Context, course *models.Course) error
	DeleteCourse(ctx context.Context, id int) error
}
