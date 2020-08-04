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

	test_data_1_0 := [3][3]int{
		{7, 3, 0},
		{5, 9, 6},
		{0, 0, 0},
	}

	test_data_2_0 := [3][3]int{
		{8, 0, 5},
		{7, 0, 3},
		{0, 9, 0},
	}

	test_data_0_1 := [3][3]int{
		{0, 0, 0},
		{0, 8, 0},
		{9, 3, 0},
	}

	test_data_1_1 := [3][3]int{
		{0, 5, 0},
		{0, 0, 0},
		{1, 7, 0},
	}

	test_data_2_1 := [3][3]int{
		{3, 7, 8},
		{9, 1, 0},
		{0, 0, 6},
	}

	test_data_0_2 := [3][3]int{
		{7, 0, 9},
		{0, 0, 0},
		{2, 5, 0},
	}

	test_data_1_2 := [3][3]int{
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
	if !sub_0_1.Contains(8) {
		t.Error("Could not find 8 in subgrid [0,1]")
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

func TestGridValidityVerification(t *testing.T) {
	valid_board := [9][9]int{
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
	valid_grid := NewGrid(valid_board)

	invalid_board := [9][9]int{
		{0, 9, 6, 7, 3, 0, 9, 0, 5},
		{0, 0, 0, 5, 9, 6, 7, 0, 3},
		{3, 7, 0, 0, 0, 0, 0, 9, 0},
		{0, 0, 0, 0, 5, 0, 3, 7, 8},
		{0, 8, 3, 0, 0, 0, 9, 1, 0},
		{9, 3, 0, 1, 7, 0, 0, 0, 6},
		{7, 0, 9, 0, 4, 5, 0, 8, 2},
		{0, 0, 0, 0, 6, 0, 5, 3, 0},
		{2, 5, 0, 8, 0, 7, 4, 0, 0},
	}

	invalid_grid := NewGrid(invalid_board)

	fmt.Printf("Valid test data:\n%v\n", valid_grid)
	if !valid_grid.Valid() {
		t.Error("Valid grid reported as invalid!")
	}
	for i := 0; i < 9; i++ {
		if !valid_grid.ValidRow(i) {
			t.Error("Invalid row detected in valid board")
		}
		if !valid_grid.ValidColumn(i) {
			t.Error("Invalid column detected in valid board")
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sub := valid_grid.GetSubGrid(i, j)
			if !sub.Valid() {
				t.Error("Invalid subgrid claimed in valid board")
			}
		}
	}

	fmt.Printf("Invalid test data:\n%v\n", invalid_grid)
	if invalid_grid.Valid() {
		t.Error("Invalid grid reported as valid")
	}
	if invalid_grid.ValidRow(0) {
		t.Error("Valid row claimed for invalid board")
	}
	if invalid_grid.ValidColumn(6) {
		t.Error("Valid column claimed for invalid board")
	}
	invalid_subgrid := invalid_grid.GetSubGrid(0, 1)
	fmt.Printf("Checking that this subgrid is invalid:\n%v\n", &invalid_subgrid)
	if invalid_subgrid.Valid() {
		t.Errorf("Valid subgrid claimed for invalid subgrid [0,1]:\n%v\n", &invalid_subgrid)
	}
}

func TestSolvedChecking(t *testing.T) {
	solved_valid_board := [9][9]int{
		{4, 2, 1, 9, 7, 3, 5, 8, 6},
		{9, 8, 3, 2, 5, 6, 4, 1, 7},
		{7, 6, 5, 4, 1, 8, 9, 3, 2},
		{1, 5, 7, 6, 4, 9, 3, 2, 8},
		{3, 9, 8, 7, 2, 5, 6, 4, 1},
		{6, 4, 2, 3, 8, 1, 7, 5, 9},
		{8, 7, 6, 5, 3, 2, 1, 9, 4},
		{5, 1, 9, 8, 6, 4, 2, 7, 3},
		{2, 3, 4, 1, 9, 7, 8, 6, 5},
	}
	solved_invalid_board := [9][9]int{
		{4, 2, 1, 9, 4, 3, 5, 8, 6},
		{9, 8, 3, 2, 5, 6, 4, 1, 7},
		{7, 6, 5, 4, 1, 8, 9, 3, 2},
		{1, 5, 7, 6, 4, 9, 3, 2, 8},
		{3, 1, 8, 7, 2, 5, 6, 4, 1},
		{6, 4, 2, 3, 8, 1, 7, 5, 9},
		{8, 7, 6, 5, 3, 2, 8, 9, 4},
		{5, 1, 9, 8, 6, 4, 2, 7, 3},
		{2, 3, 4, 1, 9, 7, 8, 6, 5},
	}
	unsolved_valid_board := [9][9]int{
		{1, 4, 2, 0, 9, 0, 0, 0, 8},
		{7, 0, 6, 8, 0, 3, 9, 0, 0},
		{3, 9, 0, 7, 1, 0, 0, 6, 0},
		{9, 0, 1, 3, 8, 5, 4, 2, 6},
		{2, 0, 0, 0, 6, 7, 3, 0, 1},
		{0, 0, 3, 0, 0, 0, 8, 7, 0},
		{0, 0, 9, 1, 7, 0, 5, 4, 2},
		{4, 0, 0, 6, 0, 0, 0, 0, 3},
		{8, 0, 5, 4, 0, 2, 0, 9, 0},
	}
	unsolved_invalid_board := [9][9]int{
		{1, 4, 2, 0, 9, 0, 0, 0, 8},
		{7, 0, 6, 8, 0, 3, 9, 0, 0},
		{3, 9, 0, 7, 1, 0, 0, 6, 0},
		{9, 0, 1, 3, 8, 4, 4, 2, 6},
		{2, 0, 0, 0, 6, 7, 3, 0, 1},
		{0, 0, 3, 0, 0, 0, 8, 7, 0},
		{0, 4, 9, 1, 7, 0, 5, 4, 2},
		{4, 0, 0, 6, 0, 0, 0, 0, 3},
		{8, 0, 5, 4, 0, 2, 0, 9, 0},
	}

	solved_valid_grid := NewGrid(solved_valid_board)
	solved_invalid_grid := NewGrid(solved_invalid_board)
	unsolved_valid_grid := NewGrid(unsolved_valid_board)
	unsolved_invalid_grid := NewGrid(unsolved_invalid_board)

	if !solved_valid_grid.Solved() {
		t.Error("Solved & valid board reported as non-solved")
	}
	if solved_invalid_grid.Solved() {
		t.Error("Solved & invalid board reported as solved")
	}
	if unsolved_valid_grid.Solved() {
		t.Error("Unsolved & valid board reported as solved")
	}
	if unsolved_invalid_grid.Solved() {
		t.Error("Unsolved & invalid board reported as solved")
	}
}
