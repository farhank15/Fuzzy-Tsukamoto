package fuzzifikasi

import "go-tsukamoto/internal/modules/utils"

// RepeatedCoursesFuzzification handles the fuzzification of repeated courses
type RepeatedCoursesFuzzification struct {
	Count int
}

// MembershipSangatRendah - fungsi monoton turun
func (r *RepeatedCoursesFuzzification) MembershipSangatRendah() float64 {
	return utils.LinearMembershipDown(float64(r.Count), 0, 1)
}

// MembershipRendah - fungsi monoton turun
func (r *RepeatedCoursesFuzzification) MembershipRendah() float64 {
	if r.Count <= 1 || r.Count >= 3 {
		return 0
	}
	return utils.LinearMembershipDown(float64(r.Count), 1, 2)
}

// MembershipSedang - fungsi monoton naik lalu turun
func (r *RepeatedCoursesFuzzification) MembershipSedang() float64 {
	count := float64(r.Count)
	if count <= 1 || count >= 3 {
		return 0
	}
	if count <= 2 {
		return utils.LinearMembershipUp(count, 1, 2)
	}
	return utils.LinearMembershipDown(count, 2, 3)
}

// MembershipTinggi - fungsi monoton naik
func (r *RepeatedCoursesFuzzification) MembershipTinggi() float64 {
	return utils.LinearMembershipUp(float64(r.Count), 3, 4)
}

// MembershipSangatTinggi - fungsi monoton naik
func (r *RepeatedCoursesFuzzification) MembershipSangatTinggi() float64 {
	return utils.LinearMembershipUp(float64(r.Count), 4, 5)
}

// FuzzifyRepeatedCourses performs fuzzification of the repeated courses
func FuzzifyRepeatedCourses(count int) map[string]float64 {
	fuzzy := &RepeatedCoursesFuzzification{Count: count}

	return map[string]float64{
		"SangatRendah": fuzzy.MembershipSangatRendah(),
		"Rendah":       fuzzy.MembershipRendah(),
		"Sedang":       fuzzy.MembershipSedang(),
		"Tinggi":       fuzzy.MembershipTinggi(),
		"SangatTinggi": fuzzy.MembershipSangatTinggi(),
	}
}
