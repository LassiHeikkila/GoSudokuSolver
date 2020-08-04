package main

import (
	"fmt"
)

// [0,0] is top left, [8,8] is bottom right
type Grid struct {
	contents [9][9]int
}

// [0,0] is top left, [2,2] is bottom right
type SubGrid struct {
	contents [3][3]int
}

func NewGrid(contents [9][9]int) *Grid {
	return &Grid{contents: contents}
}

func (g *Grid) ValueAt(x, y int) int {
	return g.contents[y][x]
}

func (g *Grid) ContainsOnRow(n, y int) bool {
	for _, val := range g.contents[y] {
		if val == n {
			return true
		}
	}
	return false
}

func (g *Grid) ContainsOnColumn(n, x int) bool {
	for _, column := range g.contents {
		if column[x] == n {
			return true
		}
	}
	return false

}

func (g *Grid) SubGridAt(x, y int) SubGrid {
	return g.GetSubGrid(x/3, y/3)
}

// [x,y] is which small square you want, [0,0] top-left, [2,2] bottom-right
func (g *Grid) GetSubGrid(x, y int) SubGrid {
	if x > 2 || y > 2 || x < 0 || y < 0 {
		panic("Invalid arguments")
	}

	// to jump to the right coordinates
	x *= 3
	y *= 3

	var s SubGrid

	copy(s.contents[0][:], g.contents[x][y:y+3])
	copy(s.contents[1][:], g.contents[x+1][y:y+3])
	copy(s.contents[2][:], g.contents[x+2][y:y+3])

	return s
}

func (g *Grid) Print() {
	fmt.Println(g)
}

func (g *Grid) String() string {
	var s string
	for x := 0; x < len(g.contents); x++ {
		for y := 0; y < len(g.contents[x]); y++ {
			if y != 0 {
				s += " "
			}
			if y%3 == 0 {
				s += " "
			}
			if g.contents[x][y] != 0 {
				s += fmt.Sprintf("%v", g.contents[x][y])
			} else {
				s += string('\u25A1')
			}
		}
		if (x+1)%3 == 0 && x != (len(g.contents)-1) {
			s += "\n"
		}
		s += "\n"
	}

	return s
}

func (s *SubGrid) Contains(n int) bool {
	for x := range s.contents {
		for y := range s.contents[x] {
			if s.contents[x][y] == n {
				return true
			}
		}
	}
	return false
}

func (s *SubGrid) ValueAt(x, y int) int {
	return s.contents[y][x]
}
