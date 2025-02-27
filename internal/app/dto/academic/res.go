package academic

import "time"

type AcademicResponse struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Ipk             float64   `json:"ipk"`
	RepeatedCourses int       `json:"repeated_courses"`
	Semester        int       `json:"semester"`
	Year            int       `json:"year"`
	PredicateID     *int      `json:"predicate_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
