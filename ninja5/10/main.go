package main

import "fmt"

// Hands-on exercise #10
// Using the code from the previous example, delete a record from your map. Now print the map
// out using the “range” loop
// solution: https://play.golang.org/p/TYl5EbjoeC
// video: 080

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["Dan"] = []string{`Chess`, `Movies`, `Audiobooks`}

	delete(m, "bond_james")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}
}
