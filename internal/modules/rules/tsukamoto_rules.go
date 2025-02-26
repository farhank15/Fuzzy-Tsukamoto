package rules

import (
	"go-tsukamoto/internal/modules/fuzzifikasi"

	log "github.com/sirupsen/logrus"
)

// TsukamotoRules menerapkan aturan Fuzzy Tsukamoto berdasarkan input
func TsukamotoRules(ipk float64, completedSemester int, repeatedCourses int, achievementRank int, achievementLevel string, thesisImpactFactor float64, thesisLevel string, activityCount int) map[string]float64 {
	// Fuzzifikasi input
	ipkFuzzy := fuzzifikasi.FuzzifyIPK(ipk)
	studyDurationFuzzy := fuzzifikasi.FuzzifyStudyDuration(completedSemester)
	repeatedCoursesFuzzy := fuzzifikasi.FuzzifyRepeatedCourses(repeatedCourses)
	achievementFuzzy := fuzzifikasi.FuzzifyAchievement(achievementRank, achievementLevel)
	thesisFuzzy := fuzzifikasi.FuzzifyThesis(thesisImpactFactor, thesisLevel)
	activityFuzzy := fuzzifikasi.FuzzifyActivity(activityCount)

	// Log hasil fuzzifikasi
	log.Infof("Fuzzifikasi IPK: %+v", ipkFuzzy)
	log.Infof("Fuzzifikasi Durasi Studi: %+v", studyDurationFuzzy)
	log.Infof("Fuzzifikasi Mata Kuliah Ulang: %+v", repeatedCoursesFuzzy)
	log.Infof("Fuzzifikasi Prestasi: %+v", achievementFuzzy)
	log.Infof("Fuzzifikasi Skripsi: %+v", thesisFuzzy)
	log.Infof("Fuzzifikasi Aktivitas: %+v", activityFuzzy)

	// Bobot faktor
	weights := map[string]float64{
		"ipk":             0.4,
		"studyDuration":   0.15,
		"repeatedCourses": 0.15,
		"achievement":     0.15,
		"thesis":          0.1,
		"activity":        0.05,
	}

	rules := map[string]float64{}

	// Summa Cum Laude
	summaCumLaudeScore := weightedMin(
		weights["ipk"], ipkFuzzy["SangatTinggi"],
		weights["studyDuration"], studyDurationFuzzy["SangatCepat"],
		weights["repeatedCourses"], repeatedCoursesFuzzy["SangatRendah"],
		weights["achievement"], achievementFuzzy["SangatTinggi"],
		weights["thesis"], thesisFuzzy["SangatTinggi"],
		weights["activity"], activityFuzzy["Tinggi"],
	)

	// Additional hard constraints for Summa Cum Laude
	if ipk >= 3.90 && repeatedCourses == 0 {
		rules["Summa Cum Laude"] = summaCumLaudeScore
	}

	// Magna Cum Laude
	magnaCumLaudeScore := weightedMin(
		weights["ipk"], ipkFuzzy["Tinggi"],
		weights["studyDuration"], max(studyDurationFuzzy["Cepat"], studyDurationFuzzy["Sedang"]),
		weights["repeatedCourses"], repeatedCoursesFuzzy["Rendah"],
		weights["achievement"], max(achievementFuzzy["Tinggi"], achievementFuzzy["Sedang"]),
		weights["thesis"], max(thesisFuzzy["Tinggi"], thesisFuzzy["Sedang"]),
		weights["activity"], activityFuzzy["Sedang"],
	)

	// Additional constraints for Magna Cum Laude
	if ipk >= 3.75 && repeatedCourses <= 1 {
		rules["Magna Cum Laude"] = magnaCumLaudeScore
	}

	// Cum Laude
	cumLaudeScore := weightedMin(
		weights["ipk"], ipkFuzzy["Sedang"],
		weights["studyDuration"], studyDurationFuzzy["Sedang"],
		weights["repeatedCourses"], repeatedCoursesFuzzy["Sedang"],
		weights["achievement"], achievementFuzzy["Sedang"],
		weights["thesis"], thesisFuzzy["Sedang"],
		weights["activity"], activityFuzzy["Sedang"],
	)

	if ipk >= 3.50 {
		rules["Cum Laude"] = cumLaudeScore
	}

	// Sangat Memuaskan
	rules["Sangat Memuaskan"] = weightedMin(
		weights["ipk"], max(ipkFuzzy["Rendah"], ipkFuzzy["Sedang"]),
		weights["studyDuration"], studyDurationFuzzy["Lama"],
		weights["repeatedCourses"], repeatedCoursesFuzzy["Tinggi"],
		weights["achievement"], achievementFuzzy["Rendah"],
		weights["thesis"], thesisFuzzy["Rendah"],
		weights["activity"], activityFuzzy["Rendah"],
	)

	// Memuaskan
	rules["Memuaskan"] = weightedMin(
		weights["ipk"], ipkFuzzy["SangatRendah"],
		weights["studyDuration"], studyDurationFuzzy["SangatLama"],
		weights["repeatedCourses"], repeatedCoursesFuzzy["SangatTinggi"],
		weights["achievement"], achievementFuzzy["SangatRendah"],
		weights["thesis"], thesisFuzzy["SangatRendah"],
		weights["activity"], activityFuzzy["SangatRendah"],
	)

	// Cukup (default jika tidak memenuhi kriteria lain)
	rules["Cukup"] = 0.1

	// Normalisasi nilai aturan
	totalWeight := 0.0
	for _, weight := range rules {
		totalWeight += weight
	}

	if totalWeight > 0 {
		for category := range rules {
			rules[category] = rules[category] / totalWeight
		}
	}

	// Log hasil aturan fuzzy
	log.Infof("Hasil Aturan Fuzzy: %+v", rules)

	return rules
}

// Fungsi perhitungan minimum berbobot
func weightedMin(weightsAndValues ...float64) float64 {
	sum := 0.0
	totalWeight := 0.0
	for i := 0; i < len(weightsAndValues); i += 2 {
		weight := weightsAndValues[i]
		value := weightsAndValues[i+1]
		sum += weight * value
		totalWeight += weight
	}
	return sum / totalWeight
}

// Fungsi untuk mengambil nilai maksimum
func max(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	maxValue := values[0]
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
