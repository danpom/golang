package luhn

import (
	"strings"
)

func Valid(id string) bool {

	//strip spaces
	id = strings.Replace(id, " ", "", -1)

	//storing the length of id number
	idLen := len(id)

	//len <= 1 not valid
	if idLen <= 1 {
		return false
	}

	//all other non-digit chars disallowed - converting to int will implicitly check for this
	// doesn't work for the case of very large ints. Will move to checking each digit in the loop below instead
	// _, err := strconv.Atoi(id)
	// if err != nil {
	// 	return false
	// }

	//variable to track every second digit working from the right
	second := false
	sum := 0
	//iterating through id starting from the right
	for i := idLen - 1; i >= 0; i-- {
		if 48 <= id[i] && id[i] <= 57 {
			idDig := int(id[i] - '0')
			if second {
				idDig = idDig * 2
				if idDig > 9 {
					idDig = idDig - 9
				}
				sum += idDig
			} else {
				sum += idDig
			}
			second = !second
		} else {
			return false
		}
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}

}
