package main

//Create a program that uses a switch statement with no switch expression specified
import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("shouldn't print")

	case true:
		fmt.Println("that's the one")
	}

}
