package fuzzifikasi

import "go-tsukamoto/internal/modules/utils"

// ActivityFuzzification handles the fuzzification of activities
type ActivityFuzzification struct {
	Count int
}

// MembershipSangatRendah - fungsi monoton turun untuk 0-1 aktivitas
func (a *ActivityFuzzification) MembershipSangatRendah() float64 {
	return utils.LinearMembershipDown(float64(a.Count), 0, 1)
}

// MembershipRendah - fungsi monoton turun untuk 1-3 aktivitas
func (a *ActivityFuzzification) MembershipRendah() float64 {
	return utils.LinearMembershipDown(float64(a.Count), 1, 3)
}

// MembershipSedang - fungsi monoton naik lalu turun untuk 3-5 aktivitas
func (a *ActivityFuzzification) MembershipSedang() float64 {
	count := float64(a.Count)
	if count <= 1 || count >= 5 {
		return 0
	}
	if count <= 3 {
		return utils.LinearMembershipUp(count, 1, 3)
	}
	return utils.LinearMembershipDown(count, 3, 5)
}

// MembershipTinggi - fungsi monoton naik untuk 5-7 aktivitas
func (a *ActivityFuzzification) MembershipTinggi() float64 {
	return utils.LinearMembershipUp(float64(a.Count), 5, 7)
}

// MembershipSangatTinggi - fungsi monoton naik untuk >6 aktivitas
func (a *ActivityFuzzification) MembershipSangatTinggi() float64 {
	return utils.LinearMembershipUp(float64(a.Count), 6, 10)
}

// FuzzifyActivity performs fuzzification of the activities
func FuzzifyActivity(count int) map[string]float64 {
	fuzzy := &ActivityFuzzification{Count: count}

	return map[string]float64{
		"SangatRendah": fuzzy.MembershipSangatRendah(),
		"Rendah":       fuzzy.MembershipRendah(),
		"Sedang":       fuzzy.MembershipSedang(),
		"Tinggi":       fuzzy.MembershipTinggi(),
		"SangatTinggi": fuzzy.MembershipSangatTinggi(),
	}
}
