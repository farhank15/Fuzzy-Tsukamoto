package activity

type CreateActivityRequest struct {
	UserID       int    `json:"user_id" validate:"required"`
	Organization string `json:"organization" validate:"required"`
	Year         int    `json:"year" validate:"required"`
}

type UpdateActivityRequest struct {
	Organization string `json:"organization"`
	Year         int    `json:"year"`
}
