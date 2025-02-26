package fuzzifikasi

import "go-tsukamoto/internal/modules/utils"

// IPKFuzzification menangani fuzzifikasi nilai IPK
type IPKFuzzification struct {
	Value float64
}

// MembershipSangatRendah - fungsi monoton turun
func (i *IPKFuzzification) MembershipSangatRendah() float64 {
	return utils.LinearMembershipDown(i.Value, 0.00, 2.00)
}

// MembershipRendah - fungsi monoton turun
func (i *IPKFuzzification) MembershipRendah() float64 {
	if i.Value <= 1.50 || i.Value >= 2.75 {
		return 0
	}
	return utils.LinearMembershipDown(i.Value, 1.50, 2.75)
}

// MembershipSedang - fungsi monoton naik lalu turun
func (i *IPKFuzzification) MembershipSedang() float64 {
	if i.Value <= 2.00 || i.Value >= 3.50 {
		return 0
	}
	if i.Value <= 2.75 {
		return utils.LinearMembershipUp(i.Value, 2.00, 2.75)
	}
	return utils.LinearMembershipDown(i.Value, 2.75, 3.50)
}

// MembershipTinggi - fungsi monoton naik
func (i *IPKFuzzification) MembershipTinggi() float64 {
	if i.Value <= 3.00 || i.Value >= 3.75 {
		return 0
	}
	return utils.LinearMembershipUp(i.Value, 3.00, 3.75)
}

// MembershipSangatTinggi - fungsi monoton naik
func (i *IPKFuzzification) MembershipSangatTinggi() float64 {
	return utils.LinearMembershipUp(i.Value, 3.50, 4.00)
}

// FuzzifyIPK melakukan fuzzifikasi nilai IPK
func FuzzifyIPK(ipk float64) map[string]float64 {
	fuzzy := &IPKFuzzification{Value: ipk}

	return map[string]float64{
		"SangatRendah": fuzzy.MembershipSangatRendah(),
		"Rendah":       fuzzy.MembershipRendah(),
		"Sedang":       fuzzy.MembershipSedang(),
		"Tinggi":       fuzzy.MembershipTinggi(),
		"SangatTinggi": fuzzy.MembershipSangatTinggi(),
	}
}
