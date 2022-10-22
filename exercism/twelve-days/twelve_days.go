package twelve

//package main

// func main() {
// 	// fmt.Println(Verse(5))
// 	// fmt.Println(Verse(12))
// 	// fmt.Println(Verse(1))
// 	fmt.Println(Song())
// }

func Verse(i int) string {
	//decrement i as arrays start from position 0
	i--

	var lyrics [12]string = [12]string{
		"twelve Drummers Drumming",
		"eleven Pipers Piping",
		"ten Lords-a-Leaping",
		"nine Ladies Dancing",
		"eight Maids-a-Milking",
		"seven Swans-a-Swimming",
		"six Geese-a-Laying",
		"five Gold Rings",
		"four Calling Birds",
		"three French Hens",
		"two Turtle Doves",
		"a Partridge in a Pear Tree.",
	}

	var day [12]string = [12]string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	var output string
	if i == 0 {
		output = "On the " + day[i] + " day of Christmas my true love gave to me: " + lyrics[11]
	} else {
		output += "On the " + day[i] + " day of Christmas my true love gave to me: "
		for j := 11 - i; j <= 11; j++ {
			if j == 11 {
				output += "and " + lyrics[j]
			} else {
				output += lyrics[j] + ", "
			}
		}
	}
	return output
}

func Song() string {
	var output string
	for i := 1; i <= 12; i++ {
		if i == 12 {
			output += Verse(i)
		} else {
			output += Verse(i) + "\n"
		}
	}
	return output
}
