package course

import "time"

type CourseResponse struct {
	ID           int       `json:"id"`
	Code         string    `json:"code"`
	CourseName   string    `json:"course_name"`
	CreditCourse int       `json:"credit_course"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
