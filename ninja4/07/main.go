package main

import "fmt"

// Hands-on exercise #7
// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.
// solution: https://play.golang.org/p/FMM4c2PhZg
// video: 077

func main() {

	var twoDSlice = make([][]string, 2)
	twoDSlice[0] = []string{"James", "Bond", "Shaken, not stirred"}
	twoDSlice[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	for _, row := range twoDSlice {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
