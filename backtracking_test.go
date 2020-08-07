package main

import (
	"testing"
	"time"
)

func TestBacktrackingSolution(t *testing.T) {
	valid_board := [9][9]int{
		{0, 0, 0, 0, 0, 3, 6, 0, 0},
		{0, 0, 5, 8, 6, 2, 7, 4, 0},
		{8, 6, 0, 5, 9, 0, 0, 0, 0},
		{0, 0, 8, 0, 0, 6, 0, 1, 3},
		{6, 0, 0, 9, 3, 5, 0, 0, 4},
		{0, 3, 9, 1, 4, 8, 5, 0, 7},
		{0, 5, 3, 7, 2, 0, 0, 0, 0},
		{0, 2, 0, 0, 8, 0, 0, 7, 5},
		{0, 8, 7, 6, 0, 4, 3, 9, 2},
	}
	grid := NewGrid(valid_board)
	t.Logf("Before solving:\n%v\n", grid)

	beginning := time.Now()
	err := SolveWithBacktracking(grid)
	dur := time.Since(beginning)
	if err != nil {
		t.Error("Failed to solve sudoku with backtracking algorithm:", err)
	}
	t.Log("Computation took", dur.Milliseconds(), "ms")
	t.Logf("Solution:\n%v\n", grid)
}
