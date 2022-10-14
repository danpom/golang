package isogram

import "strings"

//IsIsogram determines if a word or a phrase is an isogram
func IsIsogram(word string) bool {
	//remove all occurances of whitespace and hyphens
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)
	word = strings.ToLower(word)
	// initialise a slice of type string to keep track of the letters we come across
	letters := []rune{}

	// iterate through input word and check if letter exists in our letters slice
	for _, letter := range word {
		for i := 0; i < len(letters); i++ {
			if letters[i] == letter {
				return false //second occurance of the same letter
			}
		}
		letters = append(letters, letter)
	}
	return true // no letters repeated in input phrase
}
