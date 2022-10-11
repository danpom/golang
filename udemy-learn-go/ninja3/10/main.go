package main

// Write down what these print:
// 1 fmt.Println(true && true)
// 2 fmt.Println(true && false)
// 3 fmt.Println(true || true)
// 4 fmt.Println(true || false)
// 5 fmt.Println(!true)
//Answers:
//1.true
//2.false
//3.true
//4.true
//5.false
//
import (
	"fmt"
)

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
