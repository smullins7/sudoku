package board

import "fmt"

func PrintBoard(b SudokuBoard) {
	size := b.Size
	fmt.Printf("Printing board of dimensions %vx%v\n", size, size)
	for y := 0; y < size; y++ {
		fmt.Println()
		for x := 0; x < size; x++ {
			cell := b.Cells[NewPoint(x, y)]
			value := *cell.Value
			if value == 0 {
				fmt.Print("_")
			} else {
				fmt.Printf("%v", *cell.Value)
			}
		}
	}
}
