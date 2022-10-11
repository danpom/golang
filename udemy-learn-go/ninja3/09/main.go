package main

//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
import (
	"fmt"
	"strings"
)

func main() {
	//favSport := "Rugby"
	//favSport := "SOCCER"
	favSport := "Dressage"

	favSport = strings.ToLower(favSport)
	switch favSport {
	case "rugby":
		fmt.Println("Mine too!")

	case "soccer":
		fmt.Println("Not quite")

	default:
		fmt.Println("Never heard of it")
	}

}
