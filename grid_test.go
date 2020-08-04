package main

import (
	"fmt"
	"testing"
)

func TestGridFunctions(t *testing.T) {
	// generated at https://sudokuweb.org
	test_data := [9][9]int{
		{0, 0, 0, 3, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 4, 0, 3, 0},
		{0, 3, 0, 0, 8, 0, 4, 6, 1},
		{0, 9, 8, 5, 0, 1, 0, 0, 0},
		{2, 4, 3, 0, 0, 6, 0, 5, 7},
		{0, 0, 5, 4, 0, 0, 0, 0, 0},
		{4, 0, 0, 9, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 5, 0, 1, 0},
		{3, 7, 0, 0, 0, 0, 6, 9, 5},
	}

	grid := NewGrid(test_data)
	fmt.Printf("Test data:\n%v\n", grid)
	if !grid.ContainsOnRow(3, 0) {
		t.Error("Failed to find 3 on row 0")
	}
	if !grid.ContainsOnRow(9, 3) {
		t.Error("Failed to find 9 on row 3")
	}
	if grid.ContainsOnRow(6, 7) {
		t.Error("Found 6 on row 7")
	}
	if !grid.ContainsOnColumn(8, 2) {
		t.Error("Failed to find 8 on column 2")
	}
	if !grid.ContainsOnColumn(5, 5) {
		t.Error("Failed to find 5 on column 5")
	}
	if grid.ContainsOnColumn(2, 7) {
		t.Error("Found 2 on column 7")
	}
	if grid.ValueAt(0, 8) != 3 {
		t.Error("Wrong value at [0,8]")
	}
	if grid.ValueAt(5, 4) != 6 {
		t.Error("Wrong value at [5,4]")
	}
}

func TestSubGridFunctions(t *testing.T) {
	// generated at https://sudokuweb.org
	test_data := [9][9]int{
		{0, 9, 6, 7, 3, 0, 8, 0, 5},
		{0, 0, 0, 5, 9, 6, 7, 0, 3},
		{3, 7, 0, 0, 0, 0, 0, 9, 0},
		{0, 0, 0, 0, 5, 0, 3, 7, 8},
		{0, 8, 0, 0, 0, 0, 9, 1, 0},
		{9, 3, 0, 1, 7, 0, 0, 0, 6},
		{7, 0, 9, 0, 4, 5, 0, 8, 2},
		{0, 0, 0, 0, 6, 0, 5, 3, 0},
		{2, 5, 0, 8, 0, 7, 4, 0, 0},
	}

	test_data_0_0 := [3][3]int{
		{0, 9, 6},
		{0, 0, 0},
		{3, 7, 0},
	}

	test_data_0_1 := [3][3]int{
		{7, 3, 0},
		{5, 9, 6},
		{0, 0, 0},
	}

	test_data_0_2 := [3][3]int{
		{8, 0, 5},
		{7, 0, 3},
		{0, 9, 0},
	}

	test_data_1_0 := [3][3]int{
		{0, 0, 0},
		{0, 8, 0},
		{9, 3, 0},
	}

	test_data_1_1 := [3][3]int{
		{0, 5, 0},
		{0, 0, 0},
		{1, 7, 0},
	}

	test_data_1_2 := [3][3]int{
		{3, 7, 8},
		{9, 1, 0},
		{0, 0, 6},
	}

	test_data_2_0 := [3][3]int{
		{7, 0, 9},
		{0, 0, 0},
		{2, 5, 0},
	}

	test_data_2_1 := [3][3]int{
		{0, 4, 5},
		{0, 6, 0},
		{8, 0, 7},
	}

	test_data_2_2 := [3][3]int{
		{0, 8, 2},
		{5, 3, 0},
		{4, 0, 0},
	}

	grid := NewGrid(test_data)
	fmt.Printf("Test data:\n%v\n", grid)
	sub_0_0 := grid.GetSubGrid(0, 0)
	sub_0_1 := grid.GetSubGrid(0, 1)
	sub_0_2 := grid.GetSubGrid(0, 2)

	sub_1_0 := grid.GetSubGrid(1, 0)
	sub_1_1 := grid.GetSubGrid(1, 1)
	sub_1_2 := grid.GetSubGrid(1, 2)

	sub_2_0 := grid.GetSubGrid(2, 0)
	sub_2_1 := grid.GetSubGrid(2, 1)
	sub_2_2 := grid.GetSubGrid(2, 2)

	if grid.SubGridAt(2, 7) != sub_0_2 {
		t.Error("Wrong subgrid returned for coordinates [2,7]")
	}

	if sub_0_0.contents != test_data_0_0 {
		t.Error("Wrong subgrid contents at [0,0]")
	}
	if sub_0_1.contents != test_data_0_1 {
		t.Error("Wrong subgrid contents at [0,1]")
	}
	if sub_0_2.contents != test_data_0_2 {
		t.Error("Wrong subgrid contents at [0,2]")
	}

	if sub_1_0.contents != test_data_1_0 {
		t.Error("Wrong subgrid contents at [1,0]")
	}
	if sub_1_1.contents != test_data_1_1 {
		t.Error("Wrong subgrid contents at [1,1]")
	}
	if sub_1_2.contents != test_data_1_2 {
		t.Error("Wrong subgrid contents at [1,2]")
	}

	if sub_2_0.contents != test_data_2_0 {
		t.Error("Wrong subgrid contents at [2,0]")
	}
	if sub_2_1.contents != test_data_2_1 {
		t.Error("Wrong subgrid contents at [2,1]")
	}
	if sub_2_2.contents != test_data_2_2 {
		t.Error("Wrong subgrid contents at [2,2]")
	}

	if !sub_0_0.Contains(3) {
		t.Error("Could not find 3 in subgrid [0,0]")
	}
	if !sub_0_0.Contains(9) {
		t.Error("Could not find 9 in subgrid [0,0]")
	}
	if !sub_0_1.Contains(5) {
		t.Error("Could not find 5 in subgrid [0,1]")
	}
	if !sub_2_2.Contains(8) {
		t.Error("Could not find 8 in subgrid [2,2]")
	}
	if !sub_1_1.Contains(7) {
		t.Error("Could not find 7 in subgrid [1,1]")
	}
	if sub_2_1.Contains(2) {
		t.Error("Falsely found 2 in subgrid [2,1]")
	}

	if sub_2_2.ValueAt(1, 1) != 3 {
		t.Error("Wrong value for subgrid [2,2] at [1,1]")
	}
	if sub_2_2.ValueAt(0, 1) != 5 {
		t.Error("Wrong value for subgrid [2,2] at [0,1]")
	}
	if sub_2_2.ValueAt(2, 0) != 2 {
		t.Error("Wrong value for subgrid [2,2] at [2,0]")
	}
}