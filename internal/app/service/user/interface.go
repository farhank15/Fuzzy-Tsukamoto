package user

import (
	"context"
	"go-tsukamoto/internal/app/dto/user"
	"go-tsukamoto/internal/app/models"
	academicRepo "go-tsukamoto/internal/app/repository/academic"
	achievementRepo "go-tsukamoto/internal/app/repository/achievement"
	activityRepo "go-tsukamoto/internal/app/repository/activity"
	thesisRepo "go-tsukamoto/internal/app/repository/thesis"
	repo "go-tsukamoto/internal/app/repository/user"

	"gorm.io/gorm"
)

type userService struct {
	repo            repo.UserRepositoryInterface
	academicRepo    academicRepo.AcademicRepositoryInterface
	achievementRepo achievementRepo.AchievementRepositoryInterface
	activityRepo    activityRepo.ActivityRepositoryInterface
	thesisRepo      thesisRepo.ThesisRepositoryInterface
}

func NewUserService(repo repo.UserRepositoryInterface, academicRepo academicRepo.AcademicRepositoryInterface, achievementRepo achievementRepo.AchievementRepositoryInterface, activityRepo activityRepo.ActivityRepositoryInterface, thesisRepo thesisRepo.ThesisRepositoryInterface) UserService {
	return &userService{repo: repo, academicRepo: academicRepo, achievementRepo: achievementRepo, activityRepo: activityRepo, thesisRepo: thesisRepo}
}

func NewService(db *gorm.DB) UserService {
	repository := repo.NewUserRepository(db)
	academicRepository := academicRepo.NewAcademicRepository(db)
	achievementRepository := achievementRepo.NewAchievementRepository(db)
	activityRepository := activityRepo.NewActivityRepository(db)
	thesisRepository := thesisRepo.NewThesisRepository(db)
	return &userService{repo: repository, academicRepo: academicRepository, achievementRepo: achievementRepository, activityRepo: activityRepository, thesisRepo: thesisRepository}
}

type UserService interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error)
	UpdateUser(ctx context.Context, id int, req *user.UpdateUserRequest) (*user.UserResponse, error)
	GetUserByID(ctx context.Context, id int) (*user.UserResponse, error)
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	GetUserWithRelatedData(ctx context.Context, id int) (*user.UserWithRelatedDataResponse, error)
	DeleteUser(ctx context.Context, id int) error
}
