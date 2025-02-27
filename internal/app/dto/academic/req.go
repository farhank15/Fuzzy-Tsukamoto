package academic

type CreateAcademicRequest struct {
	UserID          int     `json:"user_id" validate:"required"`
	Ipk             float64 `json:"ipk" validate:"required"`
	RepeatedCourses int     `json:"repeated_courses" validate:"required"`
	Semester        int     `json:"semester" validate:"required"`
	Year            int     `json:"year" validate:"required"`
	PredicateID     int     `json:"predicate_id"`
}

type UpdateAcademicRequest struct {
	Ipk             float64 `json:"ipk"`
	RepeatedCourses int     `json:"repeated_courses"`
	Semester        int     `json:"semester"`
	Year            int     `json:"year"`
	PredicateID     int     `json:"predicate_id"`
}
