package inferensia

import (
	"go-tsukamoto/internal/modules/defuzzifikasi"
	"go-tsukamoto/internal/modules/rules"

	log "github.com/sirupsen/logrus"
)

// TsukamotoInference menjalankan proses inferensi menggunakan metode Fuzzy Tsukamoto
func TsukamotoInference(ipk float64, completedSemester int, repeatedCourses int, achievementRank int, achievementLevel string, thesisImpactFactor float64, thesisLevel string, activityCount int) string {
	// Mengambil hasil aturan Fuzzy Tsukamoto
	ruleResults := rules.TsukamotoRules(ipk, completedSemester, repeatedCourses, achievementRank, achievementLevel, thesisImpactFactor, thesisLevel, activityCount)

	// Log hasil aturan fuzzy
	log.Infof("Hasil Aturan Fuzzy: %+v", ruleResults)

	// Defuzzifikasi hasil untuk mendapatkan output final
	finalResult := defuzzifikasi.Defuzzify(ruleResults)

	// Log hasil defuzzifikasi
	log.Infof("Hasil Defuzzifikasi: %s", finalResult)

	return finalResult
}
