package user

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,max=50"`
	Name      string `json:"name" validate:"required,max=50"`
	Nim       string `json:"nim" validate:"required,max=20"`
	Password  string `json:"password" validate:"required"`
	StartYear int    `json:"start_year" validate:"required"`
}

type UpdateUserRequest struct {
	Username  string `json:"username" validate:"max=50"`
	Name      string `json:"name" validate:"max=50"`
	Nim       string `json:"nim" validate:"max=20"`
	StartYear int    `json:"start_year"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required"`
}
