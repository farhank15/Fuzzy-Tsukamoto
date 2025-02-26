package achievement

type CreateAchievementRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required,max=50"`
	Certificate bool   `json:"certificate" validate:"required"`
	Rank        int    `json:"rank" validate:"required"`
	Level       string `json:"level" validate:"required"`
	Year        int    `json:"year" validate:"required"`
}

type UpdateAchievementRequest struct {
	Title       string `json:"title" validate:"max=50"`
	Certificate bool   `json:"certificate"`
	Rank        int    `json:"rank"`
	Level       string `json:"level"`
	Year        int    `json:"year"`
}
