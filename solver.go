package main

import (
	"fmt"
)

func Solve(g *Grid) *Grid {
	valid := true
	notSolved := true
	workingGrid := NewGrid(g.contents)
	for valid && notSolved {
		if !workingGrid.Valid() {
			valid = false
			continue
		}
		if workingGrid.Solved() {
			notSolved = false
			continue
		}
		foundNext := FigureOutNextNumber(workingGrid)
		if foundNext {
			fmt.Println("Found a number")
		}
	}
	return workingGrid
}

func OnlyOnePossible(g *Grid, x, y int) (bool, int) {
	column := g.GetColumn(x)
	row := g.GetRow(y)
	subgrid := g.GetSubGrid(x, y)

	notPossibleNums := NewNumbers()
	for _, v := range column {
		notPossibleNums[v] = true
	}
	for _, v := range row {
		notPossibleNums[v] = true
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := subgrid.ValueAt(i, j)
			notPossibleNums[val] = true
		}
	}

	notPresentCount := 0
	lastNotPresent := 0
	for n, present := range notPossibleNums {
		if !present {
			notPresentCount++
			lastNotPresent = n
		}
	}

	if notPresentCount == 1 {
		return true, lastNotPresent
	}
	return false, 0
}

func FigureOutNextNumber(g *Grid) bool {
	for x := 0; x < len(g.contents); x++ {
		for y := 0; y < len(g.contents[x]); y++ {
			if onlyOnePossible, val := OnlyOnePossible(g, x, y); onlyOnePossible {
				g.contents[x][y] = val
				return true
			}
		}
	}
	return false
}
