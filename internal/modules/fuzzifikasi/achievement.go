package fuzzifikasi

import "go-tsukamoto/internal/modules/utils"

// AchievementFuzzification handles the fuzzification of achievements
type AchievementFuzzification struct {
	Rank  int
	Level string
}

// MembershipSangatTinggi - fungsi monoton turun untuk internasional
func (a *AchievementFuzzification) MembershipSangatTinggi() float64 {
	if a.Level != "internasional" {
		return 0
	}
	return utils.LinearMembershipDown(float64(a.Rank), 1, 3)
}

// MembershipTinggi - fungsi monoton turun untuk nasional atau internasional
func (a *AchievementFuzzification) MembershipTinggi() float64 {
	if a.Level != "internasional" && a.Level != "nasional" {
		return 0
	}
	return utils.LinearMembershipDown(float64(a.Rank), 1, 5)
}

// MembershipSedang - fungsi monoton turun untuk nasional
func (a *AchievementFuzzification) MembershipSedang() float64 {
	if a.Level != "nasional" {
		return 0
	}
	return utils.LinearMembershipDown(float64(a.Rank), 1, 10)
}

// MembershipRendah - fungsi monoton turun untuk internal
func (a *AchievementFuzzification) MembershipRendah() float64 {
	if a.Level != "internal" {
		return 0
	}
	return utils.LinearMembershipDown(float64(a.Rank), 1, 10)
}

// MembershipSangatRendah - nilai tetap untuk level internal atau rank > 10
func (a *AchievementFuzzification) MembershipSangatRendah() float64 {
	if a.Level == "internal" {
		return 1
	}
	if a.Rank > 10 {
		return 1
	}
	return 0
}

// FuzzifyAchievement performs fuzzification of the achievement
func FuzzifyAchievement(rank int, level string) map[string]float64 {
	fuzzy := &AchievementFuzzification{Rank: rank, Level: level}

	return map[string]float64{
		"SangatTinggi": fuzzy.MembershipSangatTinggi(),
		"Tinggi":       fuzzy.MembershipTinggi(),
		"Sedang":       fuzzy.MembershipSedang(),
		"Rendah":       fuzzy.MembershipRendah(),
		"SangatRendah": fuzzy.MembershipSangatRendah(),
	}
}
