package stringSimilarity

import (
	"math"
)

func Jaro(s1, s2 string) float64 {
	m1 := commonCharacters(s1, s2)
	m2 := commonCharacters(s2, s1)
	lm1 := float64(len(m1))
	lm2 := float64(len(m2))
	// divide transpositions by 2 (as required by the algorithm)
	t := float64(transpositions(m1, m2)) / 2
	l1 := float64(len(s1))
	l2 := float64(len(s2))

	return (1.0 / 3) * (lm1/l1 + lm2/l2 + (lm1-t)/lm1)
}

// commonCharacters returns all characters in s1 which are common (within a distance) in s2
// these characters ([]byte) are in the same order as in s1
func commonCharacters(s1, s2 string) []byte {
	// output slice
	common := make([]byte, 0)
	// calculate the maximum distance within which a character is common
	maxDistance := int(math.Floor(math.Min(float64(len(s1)), float64(len(s2))) / 2))

	// loop through all characters of s1
	for i := 0; i < len(s1); i++ {
		char := s1[i]
		// check if the current character is common in s2
		charCommon := charCommon(char, i, s2, maxDistance)
		// if it is common append it to the output
		if charCommon {
			common = append(common, char)
		}
	}
	return common
}

// charCommon checks if a character (char) at a position in a string (pos)
// is common in another string (otherString) with a maximum distance (maxDistance)
// returns a boolean value
func charCommon(char byte, pos int, otherString string, maxDistance int) bool {
	for i := 0; i < len(otherString); i++ {
		c := otherString[i]
		// not the same character
		if char != c {
			continue
		}
		// calculate the distance between the char and the current character c
		if dist := int(math.Abs(float64(pos - i))); dist <= maxDistance {
			return true
		}
	}
	return false
}

// transpositions returns the number of transpositions between a and b
// a transposition is defined by a[i] != b[i]
// it is not devided by 2 yet
func transpositions(a, b []byte) (t int) {
	la := len(a)
	lb := len(b)

	// internalTranspositions counts the number of transpositions
	// it requires that a is smaller than b
	internalTranspositions := func(a, b []byte) (t int) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t++
			}
		}
		return
	}

	// check for the smaller slice of bytes and call internalTranspositions accordingly
	if la <= lb {
		t = internalTranspositions(a, b)
	} else {
		t = internalTranspositions(b, a)
	}
	// add all the transpositions where a and b are not of equal length
	t += int(math.Abs(float64(la) - float64(lb)))
	return
}
