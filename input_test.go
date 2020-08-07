package main

import (
	"reflect"
	"testing"
)

func TestInputToGrid(t *testing.T) {
	input := `{"row0":[0,8,4,0,7,1,2,5,0],"row1":[0,0,0,0,0,0,9,3,7],"row2":[3,0,0,0,6,2,1,8,0],"row3":[5,0,0,0,4,0,0,0,1],"row4":[8,4,0,1,9,0,3,2,5],"row5":[0,9,7,0,5,3,4,6,0],"row6":[6,0,2,7,0,0,0,4,9],"row7":[0,0,0,0,3,0,0,0,2],"row8":[0,0,0,0,0,9,0,0,3]}`

	controlBoard := [9][9]int{
		{0, 8, 4, 0, 7, 1, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 9, 3, 7},
		{3, 0, 0, 0, 6, 2, 1, 8, 0},
		{5, 0, 0, 0, 4, 0, 0, 0, 1},
		{8, 4, 0, 1, 9, 0, 3, 2, 5},
		{0, 9, 7, 0, 5, 3, 4, 6, 0},
		{6, 0, 2, 7, 0, 0, 0, 4, 9},
		{0, 0, 0, 0, 3, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 9, 0, 0, 3},
	}

	controlGrid := NewGrid(controlBoard)

	inputGrid := stringToGrid(input)

	if !reflect.DeepEqual(controlGrid, inputGrid) {
		t.Error("Parsed board != control board")
	}

	input_with_newlines := `{
    "row0":[0,8,4,0,7,1,2,5,0],
    "row1":[0,0,0,0,0,0,9,3,7],
    "row2":[3,0,0,0,6,2,1,8,0],
    "row3":[5,0,0,0,4,0,0,0,1],
    "row4":[8,4,0,1,9,0,3,2,5],
    "row5":[0,9,7,0,5,3,4,6,0],
    "row6":[6,0,2,7,0,0,0,4,9],
    "row7":[0,0,0,0,3,0,0,0,2],
    "row8":[0,0,0,0,0,9,0,0,3]
}`
	inputGrid_nl := stringToGrid(input_with_newlines)

	if !reflect.DeepEqual(controlGrid, inputGrid_nl) {
		t.Error("Parsed board (with newlines) != control board")
	}
}
