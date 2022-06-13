package board

import (
	"bufio"
	"io"
	"os"
)

type Point struct {
	x int8
	y int8
}

type SudokuCell struct {
	// TODO this should be an enum....
	Value *int8
}

type SudokuBoard struct {
	Cells map[Point]SudokuCell
	Size int
}

func NewPoint(x, y int) Point {
	return Point{int8(x), int8(y)}
}

func NewEmptyBoard() SudokuBoard {
	return NewBoard(make(map[Point]SudokuCell))
}

func NewBoard(m map[Point]SudokuCell) SudokuBoard {
	return SudokuBoard{m, 0}
}

func NewCell(r rune) SudokuCell {
	if r == '.' {
		var v int8
		return SudokuCell{&v}
	} else {
		v := int8(r) - 48
		return SudokuCell{&v}
	}
}

func ParseRow(row string) []SudokuCell {
	var cells []SudokuCell
	for _, c := range row {
		cells = append(cells, NewCell(c))
	}
	return cells
}

func NewBoardFromReader(r io.Reader) SudokuBoard {
	scanner := bufio.NewScanner(r)
	b := NewEmptyBoard()
	y := 0
	for scanner.Scan() {
		cells := ParseRow(scanner.Text())
		for x, cell := range cells {
			b.Cells[NewPoint(x, y)] = cell
		}
		y = y + 1
	}
	b.Size = y
	return b

}
func NewBoardFromFile(filename string) SudokuBoard {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	return NewBoardFromReader(input)
}
