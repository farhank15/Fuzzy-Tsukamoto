package achievement

import "time"

type AchievementResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Certificate bool      `json:"certificate"`
	Rank        int       `json:"rank"`
	Level       string    `json:"level"`
	Year        int       `json:"year"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
