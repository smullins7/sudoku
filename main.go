package main

import "github.com/smullins7/sudoku/board"

func main() {
	b := board.NewBoardFromFile("inputs/board-1.txt")
	board.PrintBoard(b)
}
