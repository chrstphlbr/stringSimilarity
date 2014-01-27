package stringSimilarity

import (
	"errors"
)

const maxPrefixLength = 4

// default boost threshold
var boostThreshold float64 = 0.7

func SetBoostThreshold(bt float64) error {
	if bt < 0 || bt > 1 {
		return errors.New("Bt not in correct range [0.0,1.0]")
	}
	boostThreshold = bt
	return nil
}

// default prefix scale
var prefixScale float64 = 0.1

func SetPrefixScale(p float64) error {
	if p <= 0 || p >= 0.25 {
		return errors.New("p not in correct range ]0.0,1.0]")
	}
	prefixScale = p
	return nil
}

func JaroWinkler(s1, s2 string) (distance float64) {
	// set distance to JaroScore
	distance = Jaro(s1, s2)
	// check if the distance is above the boost threshold and therefor the boost is applied
	if distance >= boostThreshold {
		// add to the distance a factor which includes the prefix length a scale for this length and a factor that depends on the jaro distance
		distance = distance + (float64(prefixLength(s1, s2)) * prefixScale * (1 - distance))
	}
	return
}

// prefixLength calculates the length of the matching characters from the start until the first mismatch
func prefixLength(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	var size int

	// calculate size for loop
	// take the length of the smaller string as the prefix can have a maximum of the length of the smaller string
	if l1 >= l2 {
		size = l2
	} else {
		size = l1
	}

	l := 0
	for i := 0; i < size; i++ {
		e1 := s1[i]
		e2 := s2[i]
		// break condition when either the maximum prefix length is reached or the two characters do not match
		if l == maxPrefixLength || e1 != e2 {
			break
		}
		l++
	}
	return l
}
