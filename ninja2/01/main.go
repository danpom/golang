package main

import "fmt"

func main() {

	printemall(42)

}

func printemall(x int) {
	fmt.Printf("%d\n", x)
	fmt.Printf("%#b\n", x)
	fmt.Printf("%#x\n", x)
}
