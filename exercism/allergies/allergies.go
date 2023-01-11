package allergies

//dictionary of allergens and their numeric score
var ad = map[uint]string{
	128: "cats", 64: "pollen", 32: "chocolate", 16: "tomatoes", 8: "strawberries", 4: "shellfish", 2: "peanuts", 1: "eggs",
}

//index for ad so we can use it in order of highest numeric allergen and descending (map is unordered)
var adi = []uint{
	128, 64, 32, 16, 8, 4, 2, 1,
}

//Allergies outputs a slice of allergens, given a person's allergy score. It ignores allergens not listed in our map above
func Allergies(allergies uint) []string {
	var output []string

	//ignore allergies above 128
	allergies = allergies % 256

	for _, v := range adi {
		if allergies/v == 1 {
			output = append(output, ad[v])
			allergies -= v
		}

	}

	return output
}

//AllergicTo determines if a numeric allergen score "allergies" contains the given allergen
func AllergicTo(allergies uint, allergen string) bool {

	return contains(Allergies(allergies), allergen)
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
