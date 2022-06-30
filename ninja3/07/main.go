package main

//Create a program that shows the “if statement” in action
import "fmt"

func main() {
	x := "the question"
	//x := "the answer"
	//x := "something else"
	if x == "the question" {
		fmt.Println("Then what's the answer?")
	} else if x == "the answer" {
		fmt.Println("Then what's the question?")
	} else {
		fmt.Println("Try again")
	}

}
