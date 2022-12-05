package pangram

import "unicode"

type FreqMap map[rune]int

func IsPangram(input string) bool {
	m := FreqMap{}

	for _, v := range input {
		v = unicode.ToLower(v)
		if IsLetter(v) {
			m[v]++
		}
		if len(m) == 26 {
			return true
		}

	}
	return false
}

func IsLetter(r rune) bool {
	if r < 'a' || r > 'z' {
		return false
	}

	return true
}
