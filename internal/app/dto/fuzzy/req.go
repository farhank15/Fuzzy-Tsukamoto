package dto

type FuzzyRequestDTO struct {
	UserID int `json:"user_id" validate:"required"`
}
