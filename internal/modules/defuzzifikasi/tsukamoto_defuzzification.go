package defuzzifikasi

import log "github.com/sirupsen/logrus"

// Bobot nilai masing-masing kategori
var categoryWeights = map[string]float64{
	"Summa Cum Laude":  4.0,
	"Magna Cum Laude":  3.5,
	"Cum Laude":        3.0,
	"Sangat Memuaskan": 2.5,
	"Memuaskan":        2.0,
	"Cukup":            1.5,
}

// Defuzzify the results using Weighted Average Method (WAM)
func Defuzzify(ruleResults map[string]float64) string {
	log.Infof("Hasil Aturan Fuzzy: %+v", ruleResults)

	numerator := 0.0
	denominator := 0.0

	for predicate, membership := range ruleResults {
		weight := categoryWeights[predicate]
		numerator += membership * weight
		denominator += membership
	}

	// Jika tidak ada hasil, kembalikan kategori default
	if denominator == 0 {
		log.Info("Denominator 0, returning Cukup")
		return "Cukup"
	}

	// Hitung hasil akhir (WAM)
	finalScore := numerator / denominator
	log.Infof("Final Score: %f", finalScore)

	result := determineCategory(finalScore)
	log.Infof("Hasil Defuzzifikasi: %s", result)

	return result
}

// Tentukan kategori berdasarkan skor akhir
func determineCategory(score float64) string {
	switch {
	case score >= 3.75:
		return "Summa Cum Laude"
	case score >= 3.25:
		return "Magna Cum Laude"
	case score >= 2.75:
		return "Cum Laude"
	case score >= 2.25:
		return "Sangat Memuaskan"
	case score >= 1.75:
		return "Memuaskan"
	default:
		return "Cukup"
	}
}
