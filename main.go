package main

import (
	"github.com/smullins7/sudoku/board"
	"os"
)

func main() {
	boardFilename := os.Args[1]
	b := board.NewBoardFromFile(boardFilename)
	board.PrintBoard(b)
}
