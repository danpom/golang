// Package proverb is used to create a proverb given a list of words
package proverb

import "fmt"

// Proverb outputs a proverb using a given a list of words
func Proverb(rhyme []string) []string {
	var output []string

	if len(rhyme) < 1 {
		return output
	}

	first := rhyme[0]

	for i := 0; i < len(rhyme)-1; i++ {
		output = append(output, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
		//fmt.Printf("For want of a %v the %v was lost.\n", rhyme[i], rhyme[i+1])
	}

	output = append(output, fmt.Sprintf("And all for the want of a %v.", first))
	//fmt.Printf("And all for the want of a %v.", first)
	return output
}

/*
["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]
For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.
*/
