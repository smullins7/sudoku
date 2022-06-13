package board

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var zero int8 = int8(0)
var one int8 = int8(1)
var two int8 = int8(2)
var three int8 = int8(3)

func TestCell(t *testing.T) {
	cases := []struct {
		in   rune
		want SudokuCell
	}{
		{'1', SudokuCell{&one}},
		{'.', SudokuCell{&zero}},
	}
	for _, c := range cases {
		got := NewCell(c.in)
		if !cmp.Equal(c.want, got) {
			t.Errorf("NewCell(%q) == %v, want %v", c.in, *got.Value, *c.want.Value)
		}
	}
}

func TestParseRow(t *testing.T) {
	cases := []struct {
		in   string
		want []SudokuCell
	}{
		{"123", []SudokuCell{SudokuCell{&one}, SudokuCell{&two}, SudokuCell{&three}}},
	}
	for _, c := range cases {
		got := ParseRow(c.in)
		if !cmp.Equal(c.want, got) {
			t.Errorf("ParseRow(%s) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestNewBoardFromReader(t *testing.T) {
	m := map[Point]SudokuCell{
		Point{0, 0}: SudokuCell{&zero},
		Point{1, 0}: SudokuCell{&zero},
		Point{2, 0}: SudokuCell{&one},
		Point{0, 1}: SudokuCell{&two},
		Point{1, 1}: SudokuCell{&zero},
		Point{2, 1}: SudokuCell{&three},
	}
	cases := []struct {
		in   string
		want map[Point]SudokuCell
	}{
		{"..1\n2.3", m},
	}
	for _, c := range cases {
		got := NewBoardFromReader(strings.NewReader(c.in))
		if !cmp.Equal(c.want, got.Cells) {
			t.Errorf("Wrong object received, got=%s", cmp.Diff(c.want, got.Cells))
		}
	}
}
