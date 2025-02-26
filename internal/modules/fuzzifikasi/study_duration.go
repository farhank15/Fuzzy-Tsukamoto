package fuzzifikasi

import "go-tsukamoto/internal/modules/utils"

// StudyDurationFuzzification menangani fuzzifikasi durasi studi dengan metode Tsukamoto
type StudyDurationFuzzification struct {
	CompletedSemester int
}

// MembershipSangatCepat - fungsi monoton turun
func (s *StudyDurationFuzzification) MembershipSangatCepat() float64 {
	semester := float64(s.CompletedSemester)
	return utils.LinearMembershipDown(semester, 6, 8)
}

// MembershipCepat - fungsi monoton turun
func (s *StudyDurationFuzzification) MembershipCepat() float64 {
	semester := float64(s.CompletedSemester)
	return utils.LinearMembershipDown(semester, 7, 9)
}

// MembershipSedang - menggunakan 2 fungsi monoton untuk representasi
func (s *StudyDurationFuzzification) MembershipSedang() float64 {
	semester := float64(s.CompletedSemester)
	if semester < 8 || semester > 10 {
		return 0
	}

	if semester <= 9 {
		return utils.LinearMembershipUp(semester, 8, 9)
	}
	return utils.LinearMembershipDown(semester, 9, 10)
}

// MembershipLama - fungsi monoton naik
func (s *StudyDurationFuzzification) MembershipLama() float64 {
	semester := float64(s.CompletedSemester)
	return utils.LinearMembershipUp(semester, 9, 11)
}

// MembershipSangatLama - fungsi monoton naik
func (s *StudyDurationFuzzification) MembershipSangatLama() float64 {
	semester := float64(s.CompletedSemester)
	return utils.LinearMembershipUp(semester, 11, 14)
}

// FuzzifyStudyDuration melakukan fuzzifikasi durasi studi
func FuzzifyStudyDuration(completedSemester int) map[string]float64 {
	fuzzy := &StudyDurationFuzzification{CompletedSemester: completedSemester}

	return map[string]float64{
		"SangatCepat": fuzzy.MembershipSangatCepat(),
		"Cepat":       fuzzy.MembershipCepat(),
		"Sedang":      fuzzy.MembershipSedang(),
		"Lama":        fuzzy.MembershipLama(),
		"SangatLama":  fuzzy.MembershipSangatLama(),
	}
}
