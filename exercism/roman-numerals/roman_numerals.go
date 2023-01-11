package romannumerals

import "errors"

var romDict = map[int]string{
	1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
}

var romIndex = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func ToRomanNumeral(input int) (string, error) {
	var output string
	if input <= 0 || input >= 4000 {
		return output, errors.New("input out of range")
	}
	//3999
	for _, v := range romIndex {
		n := input / v
		input = input % v

		for i := 0; i < n; i++ {
			output += romDict[v]
		}
	}
	return output, nil
}
