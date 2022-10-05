package main

import "fmt"

// Hands-on exercise #9
// Using the code from the previous example, add a record to your map. Now print the map out
// using the “range” loop
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 45solution: https://play.golang.org/p/_CkxAhRrDJ
// video: 079

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println("This is the record for", k)
	// 	for i, v2 := range v {
	// 		fmt.Println("\t", i, v2)
	// 	}
	// }

	m["Dan"] = []string{`Chess`, `Movies`, `Audiobooks`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
