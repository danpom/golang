package main

import "fmt"

//
//define chess board as 2 dimensional array of
//note bitboards seem to be the most efficient way for full chess engines but lets tackle that later

type Board struct {
	board [8][8]int
}

func (b Board) printBoard() {
	for _, v := range b.board {
		fmt.Println(v)
	}

}

func (b *Board) resetBoard() {
	//fmt.Println("printing board at start of function")
	//b.printBoard()
	b.board = [8][8]int{
		{8, 9, 10, 11, 12, 10, 9, 8},
		{7, 7, 7, 7, 7, 7, 7, 7},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{2, 3, 4, 5, 6, 4, 3, 2}}
	//fmt.Println("printing board at end of function")
	//b.printBoard()

}

func main() {

	var board1 Board
	//board1.printBoard()
	board1.resetBoard()

	// board1.board = [8][8]int{
	// 	{8, 9, 10, 11, 12, 10, 9, 8},
	// 	{7, 7, 7, 7, 7, 7, 7, 7},
	// 	{0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0},
	// 	{1, 1, 1, 1, 1, 1, 1, 1},
	// 	{2, 3, 4, 5, 6, 4, 3, 2}}
	//print the board

	fmt.Println("printing board at end of main")
	board1.printBoard()

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

	//resetBoard()

}
