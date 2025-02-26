package activity

import "time"

type ActivityResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Organization string    `json:"organization"`
	Year         int       `json:"year"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
