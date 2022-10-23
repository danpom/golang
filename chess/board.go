package main

import "fmt"

func main() {

	//
	//define chess board as 2 dimensional array of
	//note bitboards seem to be the most efficient way for full chess engines but lets tackle that later
	board := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board = [8][8]int{
		{8, 9, 10, 11, 12, 10, 9, 8},
		{7, 7, 7, 7, 7, 7, 7, 7},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{2, 3, 4, 5, 6, 4, 3, 2},
	}

	/*pieces := map[int]string{
		0: "empty"
		1:  "pawn(w)",
		2:  "rook(w)",
		3:  "knight(w)",
		4:  "bishop(w)",
		5:  "queen(w)",
		6:  "king(w)",
		7:  "pawn(b)",
		8:  "rook(b)",
		9:  "knight(b)",
		10:  "bishop(b)",
		11: "queen(b)",
		12: "king(b)",
	}*/

	//this should be defined in a struct with methods but again lets do some experimentation first
	// type board struct {
	//}

	//print the board
	for i, v := range board {
		fmt.Println(i, v)
	}

	//fmt.Print(board)

}
