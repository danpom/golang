package scrabble

import "strings"

//Score calculates the score of a word in scrabble
func Score(word string) int {
	//note map data structure should be an adequete solution in this case however considering efficiency
	// it may be worthwhile to consider an ordered map. Sorted in order of most to least frequently searched in scrabble

	score := 0

	scoreMap := map[rune]int{
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
		'L': 1,
		'N': 1,
		'R': 1,
		'S': 1,
		'T': 1,
		'D': 2,
		'G': 2,
		'B': 3,
		'C': 3,
		'M': 3,
		'P': 3,
		'F': 4,
		'H': 4,
		'V': 4,
		'W': 4,
		'Y': 4,
		'K': 5,
		'J': 8,
		'X': 8,
		'Q': 10,
		'Z': 10,
	}

	for _, letter := range strings.ToUpper(word) {
		for key, value := range scoreMap {
			if letter == key {
				score += value
			}
		}
	}
	return score
}
