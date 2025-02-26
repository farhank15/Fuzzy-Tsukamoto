package user

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Nim       string    `json:"nim"`
	Password  string    `json:"password"`
	StartYear int       `json:"start_year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}

type UserWithRelatedDataResponse struct {
	ID           int           `json:"id"`
	Username     string        `json:"username"`
	Name         string        `json:"name"`
	Nim          string        `json:"nim"`
	Password     string        `json:"password"`
	StartYear    int           `json:"start_year"`
	Academics    []interface{} `json:"academic"`
	Achievements []interface{} `json:"achievements"`
	Activities   []interface{} `json:"activity"`
	Theses       []interface{} `json:"thesis"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
