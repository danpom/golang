package raindrops

import "strconv"

//converts a number in to a string of raindrop sounds depending on whether they are a factoer of 3,5, and/or 7
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(number)
	} else {
		return result
	}

}
