package thesis

type CreateThesisRequest struct {
	UserID   int    `json:"user_id" validate:"required"`
	Title    string `json:"title" validate:"required,max=255"`
	Year     int    `json:"year" validate:"required"`
	Semester int    `json:"semester" validate:"required"`
	Value    string `json:"value" validate:"required,max=50"`
	Level    string `json:"level" validate:"required,max=50"`
}

type UpdateThesisRequest struct {
	UserID   int    `json:"user_id"`
	Title    string `json:"title" validate:"max=255"`
	Year     int    `json:"year"`
	Semester int    `json:"semester"`
	Value    string `json:"value" validate:"max=50"`
	Level    string `json:"level" validate:"max=50"`
}
