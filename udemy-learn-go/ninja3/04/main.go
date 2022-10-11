package main

// Create a for loop using this syntax
// ● for { }
// Have it print out the years you have been alive
import (
	"fmt"
	"time"
)

func main() {
	birthdate := 1990
	year := time.Now().Year()
	for {
		if birthdate > year {
			break
		}
		fmt.Println(birthdate)
		birthdate++
	}

}
