package thesis

import "time"

type ThesisResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Semester  int       `json:"semester"`
	Value     string    `json:"value"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
