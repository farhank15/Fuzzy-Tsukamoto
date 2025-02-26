package fuzzifikasi

import (
	"go-tsukamoto/internal/modules/utils"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ThesisFuzzification handles the fuzzification of thesis or scientific publications
type ThesisFuzzification struct {
	ImpactFactor float64
	Level        string
}

// MembershipSangatTinggi - fungsi monoton naik
// Requirement: Impact Factor > 5 dan Level Internasional
func (t *ThesisFuzzification) MembershipSangatTinggi() float64 {
	if strings.ToLower(t.Level) != "internasional" {
		return 0
	}
	return utils.LinearMembershipUp(t.ImpactFactor, 5, 7)
}

// MembershipTinggi - fungsi monoton naik untuk internasional
// Requirement: Impact Factor 3-5 dan Level Internasional
func (t *ThesisFuzzification) MembershipTinggi() float64 {
	if strings.ToLower(t.Level) != "internasional" {
		return 0
	}
	if t.ImpactFactor <= 3 {
		return 0
	}
	if t.ImpactFactor >= 5 {
		return 1
	}
	return utils.LinearMembershipUp(t.ImpactFactor, 3, 5)
}

// MembershipSedang - fungsi monoton naik untuk nasional
// Requirement: Impact Factor 1-3 dan Level Nasional
func (t *ThesisFuzzification) MembershipSedang() float64 {
	if strings.ToLower(t.Level) != "nasional" {
		return 0
	}
	if t.ImpactFactor <= 1 {
		return 0
	}
	if t.ImpactFactor >= 3 {
		return 1
	}
	return utils.LinearMembershipUp(t.ImpactFactor, 1, 3)
}

// MembershipRendah - fungsi monoton turun untuk nasional
// Requirement: Impact Factor < 1 dan Level Nasional
func (t *ThesisFuzzification) MembershipRendah() float64 {
	if strings.ToLower(t.Level) != "nasional" {
		return 0
	}
	if t.ImpactFactor >= 1 {
		return 0
	}
	return utils.LinearMembershipDown(t.ImpactFactor, 0, 1)
}

// MembershipSangatRendah - nilai tetap untuk level internal
// Requirement: Level Internal (tanpa melihat Impact Factor)
func (t *ThesisFuzzification) MembershipSangatRendah() float64 {
	if strings.ToLower(t.Level) == "internal" {
		return 1
	}
	return 0
}

// FuzzifyThesis performs fuzzification of the thesis or scientific publication
func FuzzifyThesis(impactFactor float64, level string) map[string]float64 {
	fuzzy := &ThesisFuzzification{
		ImpactFactor: impactFactor,
		Level:        level,
	}

	result := map[string]float64{
		"SangatTinggi": fuzzy.MembershipSangatTinggi(),
		"Tinggi":       fuzzy.MembershipTinggi(),
		"Sedang":       fuzzy.MembershipSedang(),
		"Rendah":       fuzzy.MembershipRendah(),
		"SangatRendah": fuzzy.MembershipSangatRendah(),
	}

	log.Infof("Fuzzifikasi Skripsi: %+v", result)
	return result
}
