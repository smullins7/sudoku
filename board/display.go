package board

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func printRowDivider(size int, c *color.Color) {
	_, err := c.Print(strings.Repeat("+---", size))
	if err != nil {
		return
	}
	_, err = c.Println("+")
	if err != nil {
		return
	}
}

func PrintBoard(b SudokuBoard) {
	size := b.Size
	bold := color.New(color.FgHiCyan).Add(color.Bold)
	pale := color.New(color.FgWhite)
	printRowDivider(size, bold)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x % 3 == 0 {
				_, _ = bold.Print("|")
			} else {
				_, _ = pale.Print("|")
			}

			cell := b.Cells[NewPoint(x, y)]
			value := *cell.Value
			if value == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %v ", *cell.Value)
			}
		}
		_, _ = bold.Print("|")
		fmt.Println()
		if (y + 1) % 3 == 0 {
			printRowDivider(size, bold)
		} else {
			printRowDivider(size, pale)
		}

	}
}
