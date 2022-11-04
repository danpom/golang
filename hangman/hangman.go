package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type gameState struct {
	//users guesses
	correctGuesses   []rune
	incorrectGuesses []rune
	//variable to store the bot's selected hidden phrase
	hf string

	//store the letters of the hidden phrase
	letters []rune
	//count stores the number of guesses allowed before the game is lost
	count int
}

func (gs *gameState) whatsLeft() (string, bool) {
	var output string
	gameWon := true
	for i := 0; i < len(gs.hf); i++ {
		switch {
		case string(gs.hf[i]) == " ":
			output += "/"
		case contains(gs.correctGuesses, rune(gs.hf[i])):
			output += string(gs.hf[i])
		default:
			output += "_"
			gameWon = false
		}
	}
	return output, gameWon
}

var pics = [7]string{"  +---+\n  |   |\n      |\n      |\n      |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n      |\n      |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n  |   |\n      |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n /|   |\n      |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n /|\\  |\n      |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n /|\\  |\n /    |\n      |\n=========",

	"  +---+\n  |   |\n  O   |\n /|\\  |\n / \\  |\n      |\n=========",
}

var movies = []string{
	"the godfather",
	"the shawshank redemption",
	"raging bull",
	"casablanca",
	"citizen Kane",
	"gone with the wind",
	"the wizard of oz",
	"lawrence of arabia",
	"vertigo",
	"psycho",
	"on the waterfront",
	"sunset boulevard",
	"forrest gump",
	"the sound of music",
	"west side story",
	"star wars",
}

/*
going with a different implementation so  letterIndex method not needed anymore and need a new method to find letters in phrase
//map to store the indexes of letters in a phrase
type letterIndex map[rune][]int

//given a string, letterIndexes returns a letterIndex map
func letterIndexes(p string) letterIndex {
	m := letterIndex{}
	for i, r := range p {
		m[r] = append(m[r], i)
	}
	return m
}
*/

//given a string, letters returns a slice of rune consosting of the letters from the string
func letters(s string) []rune {
	r := []rune(s)
	if len(s) < 1 {
		return r
	}
	//sort slice of rune
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			r[prev] = r[curr]
			prev++
		}
	}
	return r[:prev]
}

//contains returns true if an array of runes contains a given rune r.
func contains(runes []rune, r rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}

//convert an array of ruins to a string of

func main() {
	//initialising a gamestate for a new game
	gs := gameState{}

	//assigning random movie from movies array as the hidden phrase
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	gs.hf = movies[r.Intn(len(movies))]
	gs.letters = letters(gs.hf)
	gs.count = 0

	gameover := "Answer was: " + gs.hf + "\nGame over, you lose!"
	//allGuesses := append(gs.incorrectGuesses, gs.correctGuesses...)

	for gs.count <= 6 {
		revealed, gameWon := gs.whatsLeft()
		fmt.Println(revealed)
		if gameWon {
			gameover = "Congratulations, you won!"
			break
		}
		fmt.Println()
		fmt.Println("Your guesses so far: ", string((append(gs.incorrectGuesses, gs.correctGuesses...))))
		fmt.Print("Your move:")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		guess := strings.ToLower(input.Text())

		//!!todo check if input is one just one character and a valid letter e.g. a-z
		switch {
		case len(guess) != 1:
			fmt.Println("Sorry, your guess must be a single character. Try again.")
		default:
			//converting the first character of the users input string to a rune (should only be one character)
			guessRune := []rune(guess)[0]
			//fmt.Println("your guesses: ", guessRune)

			//have they guessed this already?

			if contains(append(gs.incorrectGuesses, gs.correctGuesses...), guessRune) {
				fmt.Printf("Already guessed %q, please try again\n", guess)
				//fmt.Println("Your guesses: ", append(gs.incorrectGuesses, gs.correctGuesses...))
			} else {
				//is it correct?
				if contains(gs.letters, guessRune) {
					gs.correctGuesses = append(gs.correctGuesses, guessRune)
				} else {
					gs.incorrectGuesses = append(gs.incorrectGuesses, guessRune)
					fmt.Println(pics[gs.count])
					gs.count++
				}

			}
		}
	}

	fmt.Println(gameover)
}
