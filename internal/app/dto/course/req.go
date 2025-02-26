package course

type CreateCourseRequest struct {
	Code         string `json:"code" validate:"required"`
	CourseName   string `json:"course_name" validate:"required"`
	CreditCourse int    `json:"credit_course" validate:"required"`
}

type UpdateCourseRequest struct {
	Code         string `json:"code"`
	CourseName   string `json:"course_name"`
	CreditCourse int    `json:"credit_course"`
}
