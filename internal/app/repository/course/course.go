package course

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepositoryInterface {
	return &courseRepository{db: db}
}

func (r *courseRepository) CreateCourse(ctx context.Context, course *models.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r *courseRepository) GetCourseByID(ctx context.Context, id int) (*models.Course, error) {
	var course models.Course
	if err := r.db.WithContext(ctx).First(&course, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) GetCourses(ctx context.Context) ([]*models.Course, error) {
	var courses []*models.Course
	if err := r.db.WithContext(ctx).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) UpdateCourse(ctx context.Context, course *models.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

func (r *courseRepository) DeleteCourse(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Course{}, id).Error
}
