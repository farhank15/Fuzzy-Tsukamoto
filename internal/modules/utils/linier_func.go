package utils

import log "github.com/sirupsen/logrus"

// LinearMembershipDown menghitung nilai keanggotaan untuk fungsi monoton turun
// x adalah nilai input, min adalah batas bawah, max adalah batas atas
func LinearMembershipDown(x, min, max float64) float64 {
	if x <= min {
		log.Infof("LinearMembershipDown: x=%f, min=%f, max=%f, result=1", x, min, max)
		return 1
	}
	if x >= max {
		log.Infof("LinearMembershipDown: x=%f, min=%f, max=%f, result=0", x, min, max)
		return 0
	}
	result := (max - x) / (max - min)
	log.Infof("LinearMembershipDown: x=%f, min=%f, max=%f, result=%f", x, min, max, result)
	return result
}

// LinearMembershipUp menghitung nilai keanggotaan untuk fungsi monoton naik
// x adalah nilai input, min adalah batas bawah, max adalah batas atas
func LinearMembershipUp(x, min, max float64) float64 {
	if x <= min {
		log.Infof("LinearMembershipUp: x=%f, min=%f, max=%f, result=0", x, min, max)
		return 0
	}
	if x >= max {
		log.Infof("LinearMembershipUp: x=%f, min=%f, max=%f, result=1", x, min, max)
		return 1
	}
	result := (x - min) / (max - min)
	log.Infof("LinearMembershipUp: x=%f, min=%f, max=%f, result=%f", x, min, max, result)
	return result
}
