package main

import ()

type Invalid struct {
	error
}

func (e Invalid) Error() string {
	return "Invalid grid"
}

type NoSolution struct {
	error
}

func (e NoSolution) Error() string {
	return "No solution"
}

func SolveWithBacktracking(g *Grid) error {
	if g.Valid() && g.Solved() {
		return nil
	}
	if !g.Valid() {
		return Invalid{}
	}
	for x := 0; x < len(g.contents); x++ {
		for y := 0; y < len(g.contents[x]); y++ {
			if g.contents[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if g.Possible(x, y, n) {
						g.contents[y][x] = n
						err := SolveWithBacktracking(g)
						if err == nil {
							return nil
						}

						g.contents[y][x] = 0
					}
				}
				return NoSolution{}
			}
		}
	}
	return nil
}
