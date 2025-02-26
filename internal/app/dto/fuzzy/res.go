package dto

type FuzzyResponseDTO struct {
	StudentID       int     `json:"student_id"`
	IPK             float64 `json:"ipk"`
	Semester        int     `json:"semester"`
	MataKuliahUlang int     `json:"mata_kuliah_ulang"`
	PrestasiLevel   string  `json:"prestasi_level"`
	PrestasiRank    int     `json:"prestasi_rank"`
	SkripsiLevel    string  `json:"skripsi_level"`
	SkripsiImpact   float64 `json:"skripsi_impact"`
	JumlahAktivitas int     `json:"jumlah_aktivitas"`
	HasilPredicate  string  `json:"hasil_predicate"`
}
