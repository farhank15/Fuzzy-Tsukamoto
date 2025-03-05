package user

import (
	"context"
	"errors"
	"go-tsukamoto/internal/app/dto/user"
	"go-tsukamoto/internal/app/models"
	"go-tsukamoto/utils"
	"log"
	"time"
)

func (s *userService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	// Check if NIM already exists
	existingUser, err := s.repo.GetUserByNim(ctx, req.Nim)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("NIM already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	log.Printf("Hashed Password: %s", hashedPassword) // Log the hashed password
	userModel := &models.Users{
		Username:  req.Username,
		Name:      req.Name,
		Nim:       req.Nim,
		Password:  hashedPassword,
		StartYear: req.StartYear,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.CreateUser(ctx, userModel); err != nil {
		return nil, err
	}
	return &user.UserResponse{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Name:      userModel.Name,
		Nim:       userModel.Nim,
		Password:  hashedPassword,
		StartYear: userModel.StartYear,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*user.UserResponse, error) {
	userModel, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}
	return &user.UserResponse{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Name:      userModel.Name,
		Nim:       userModel.Nim,
		Password:  userModel.Password, // Include hashed password in response
		StartYear: userModel.StartYear,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}

func (s *userService) GetUserWithRelatedData(ctx context.Context, id int) (*user.UserWithRelatedDataResponse, error) {
	userModel, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	academics, err := s.academicRepo.GetAcademicsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	achievements, err := s.achievementRepo.GetAchievementsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	activities, err := s.activityRepo.GetActivitiesByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	theses, err := s.thesisRepo.GetThesesByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user.UserWithRelatedDataResponse{
		ID:           userModel.ID,
		Username:     userModel.Username,
		Name:         userModel.Name,
		Nim:          userModel.Nim,
		Password:     userModel.Password,
		StartYear:    userModel.StartYear,
		Academics:    convertToInterfaceSlice(academics),
		Achievements: convertToInterfaceSlice(achievements),
		Activities:   convertToInterfaceSlice(activities),
		Theses:       convertToInterfaceSlice(theses),
		CreatedAt:    userModel.CreatedAt,
		UpdatedAt:    userModel.UpdatedAt,
	}, nil
}

func convertToInterfaceSlice[T any](input []*T) []interface{} {
	output := make([]interface{}, len(input))
	for i, v := range input {
		output[i] = v
	}
	return output
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	userModel, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	userModel, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}
	userModel.Username = req.Username
	userModel.Name = req.Name
	userModel.Nim = req.Nim
	userModel.StartYear = req.StartYear
	userModel.UpdatedAt = time.Now()

	if err := s.repo.UpdateUser(ctx, userModel); err != nil {
		return nil, err
	}
	return &user.UserResponse{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Name:      userModel.Name,
		Nim:       userModel.Nim,
		StartYear: userModel.StartYear,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
