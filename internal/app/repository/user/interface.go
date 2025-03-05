package user

import (
	"context"
	"go-tsukamoto/internal/app/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.Users) error
	GetUserByID(ctx context.Context, id int) (*models.Users, error)
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	GetUserByNim(ctx context.Context, nim string) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) error
	DeleteUser(ctx context.Context, id int) error
}
