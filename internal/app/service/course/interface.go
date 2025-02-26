package course

import (
	"context"
	"go-tsukamoto/internal/app/dto/course"
	repo "go-tsukamoto/internal/app/repository/course"

	"gorm.io/gorm"
)

type courseService struct {
	repo repo.CourseRepositoryInterface
}

func NewCourseService(repo repo.CourseRepositoryInterface) CourseServiceInterface {
	return &courseService{repo: repo}
}

func NewService(db *gorm.DB) CourseServiceInterface {
	repository := repo.NewCourseRepository(db)
	return &courseService{repo: repository}
}

type CourseServiceInterface interface {
	CreateCourse(ctx context.Context, req *course.CreateCourseRequest) (*course.CourseResponse, error)
	GetCourseByID(ctx context.Context, id int) (*course.CourseResponse, error)
	GetCourses(ctx context.Context) ([]*course.CourseResponse, error)
	UpdateCourse(ctx context.Context, id int, req *course.UpdateCourseRequest) (*course.CourseResponse, error)
	DeleteCourse(ctx context.Context, id int) error
	ImportCourses(ctx context.Context, reqs []course.CreateCourseRequest) error
}
