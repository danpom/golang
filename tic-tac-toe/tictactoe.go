package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

//Square is a type to track the state of each square on the tic-tac-toe board. Default is Empty with the other options being "X" or "O"
type Square int

const (
	Empty Square = iota
	O
	X
)

//String method for Square values
func (s Square) String() string {
	switch s {
	case O:
		return "O"
	case X:
		return "X"
	case Empty:
		return " "
	}
	return "unknown"
}

//Board is a 2 dimensional 3*3 array of type Square
type Board [3][3]Square

//printBoard prints the current board position
func (b Board) printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			fmt.Print(b[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		// fmt.Println("---")
	}

}

//printBoard prints the current board position
func (b Board) printBoardEmptyCoords() {
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == Empty {
				fmt.Print(count)
			} else {
				fmt.Print(b[i][j])
			}
			if j < 2 {
				fmt.Print("|")
			}
			count++
		}
		fmt.Println()
		// fmt.Println("---")
	}

}

//given an int to indicate the square (1-9) and a bool to indicate the turn. Change the square to X or O
func (b *Board) move(coord int, turn bool) error {

	//place is used to indicate which player piece we're using
	var place Square
	if turn {
		place = O
	} else {
		place = X
	}
	//placing the player piece on the desiganted coordinate
	switch coord {
	case 1:
		if b[0][0] == Empty {
			b[0][0] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 2:
		if b[0][1] == Empty {
			b[0][1] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 3:
		if b[0][2] == Empty {
			b[0][2] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 4:
		if b[1][0] == Empty {
			b[1][0] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 5:
		if b[1][1] == Empty {
			b[1][1] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 6:
		if b[1][2] == Empty {
			b[1][2] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 7:
		if b[2][0] == Empty {
			b[2][0] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 8:
		if b[2][1] == Empty {
			b[2][1] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	case 9:
		if b[2][2] == Empty {
			b[2][2] = place
			return nil
		} else {
			return errors.New("Square already occupied. Try again")
		}
	default:
		return errors.New("invalid square. Try again")
	}
}

//isWin() returns true if the Board is winning position. There are surely more efficient ways to track this but I'd like a working solution first
func (b *Board) isWin() bool {
	switch {
	//rows
	case (b[0][0] == b[0][1]) && (b[0][0] == b[0][2]) && (b[0][0] != Empty):
		return true
	case (b[1][0] == b[1][1]) && (b[1][0] == b[1][2]) && (b[1][0] != Empty):
		return true
	case (b[2][0] == b[2][1]) && (b[2][0] == b[2][2]) && (b[2][0] != Empty):
		return true
		//columns
	case (b[0][0] == b[1][0]) && (b[1][0] == b[2][0]) && (b[0][0] != Empty):
		return true
	case (b[0][1] == b[1][1]) && (b[0][1] == b[2][1]) && (b[0][1] != Empty):
		return true
	case (b[0][2] == b[1][2]) && (b[0][2] == b[2][2]) && (b[0][2] != Empty):
		return true
		//diagonals
	case (b[0][0] == b[1][1]) && (b[0][0] == b[2][2]) && (b[0][0] != Empty):
		return true
	case (b[0][2] == b[1][1]) && (b[0][2] == b[2][0]) && (b[0][2] != Empty):
		return true
	default:
		return false
	}
}

func main() {

	var board1 Board
	var turn bool
	var player string
	//create isWin() which will determine if there is a winning position

	//check if isWin if false continue playing else output the winner
	draw := 0
	fmt.Println("starting game")
	fmt.Println("printing board at start")
	board1.printBoard()

	//fmt.Println("isBoardwin?: ", board1.isWin())
	for !board1.isWin() {
		if turn {
			player = "O"
		} else {
			player = "X"
		}
		//fmt.Println("reached here")
		//play

		fmt.Printf("%v Player's turn. Choose your coordinate from the available squares below (input 1 -9):\n", player)

		board1.printBoardEmptyCoords()

		fmt.Print("Your move:")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		//fmt.Println(input.Text())
		coord, coordError := strconv.Atoi(input.Text())
		if coordError != nil {
			fmt.Println("invalid move input error: ", coordError)
			turn = !turn
		} else {
			//fmt.Println("coord: ", coord)
			moveError := board1.move(coord, turn)
			if moveError != nil {
				fmt.Println(moveError)
			} else {
				draw++
				turn = !turn
				fmt.Println("printing board after move")
				board1.printBoard()
				fmt.Println()
				if draw == 9 {
					fmt.Println("Draw!")
					break
				}
			}
		}

	}

	//output who won
	if draw != 9 {
		fmt.Printf("Congratulations %v player, you won!", player)
	}

}
