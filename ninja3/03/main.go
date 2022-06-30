package main

// Create a for loop using this syntax
// ● for condition { }
// Have it print out the years you have been alive
import (
	"fmt"
	"time"
)

func main() {
	birthdate := 1990
	year := time.Now().Year()
	for birthdate <= year {
		fmt.Println(birthdate)
		birthdate++
	}

}
