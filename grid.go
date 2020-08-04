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

type Numbers map[int]bool

func NewNumbers() Numbers {
	return map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}
}

func (n Numbers) Contains(i int) bool {
	_, ok := n[i]
	return ok
}

func (n Numbers) MissingNumbers() []int {
	var r []int
	for num, _ := range n {
		if n[num] == false {
			r = append(r, num)
		}
	}
	return r
}

func (n Numbers) ContainedNumbers() []int {
	var r []int
	for num, _ := range n {
		if n[num] == true {
			r = append(r, num)
		}
	}
	return r
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

func (g *Grid) Solved() bool {
	if !g.Valid() {
		return false
	}

	for x := 0; x < len(g.contents); x++ {
		for y := 0; y < len(g.contents[x]); y++ {
			if g.contents[x][y] == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Grid) GetColumn(x int) [9]int {
	var column [9]int

	for n, col := range g.contents {
		column[n] = col[x]
	}
	return column
}

func (g *Grid) GetRow(y int) [9]int {
	var row [9]int
	for n, val := range g.contents[y] {
		row[n] = val
	}
	return row
}

func (g *Grid) Valid() bool {
	for i := 0; i < len(g.contents); i++ {
		if !g.ValidRow(i) {
			return false
		}
		if !g.ValidColumn(i) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sub := g.GetSubGrid(i, j)
			if !sub.Valid() {
				return false
			}
		}
	}
	return true
}

func (g *Grid) ValidRow(y int) bool {
	found := NewNumbers()
	for _, value := range g.contents[y] {
		if val, ok := found[value]; ok && !val {
			found[value] = true
		} else if ok {
			return false
		}
	}

	return true
}

func (g *Grid) ValidColumn(x int) bool {
	found := NewNumbers()

	for _, value := range g.contents {
		if val, ok := found[value[x]]; ok && !val {
			found[value[x]] = true
		} else if ok {
			return false
		}
	}
	return true
}

func (g *Grid) SubGridAt(x, y int) SubGrid {
	return g.GetSubGrid(x/3, y/3)
}

// [x,y] is which small square you want,
// [0,0] top-left, [2,2] bottom-right
// [0,1] middle-left [1,0] top-middle
func (g *Grid) GetSubGrid(x, y int) SubGrid {
	if x > 2 || y > 2 || x < 0 || y < 0 {
		panic("Invalid arguments")
	}

	// to jump to the right coordinates
	x *= 3
	y *= 3

	var s SubGrid
	for i := 0; i < len(s.contents); i++ {
		for j := 0; j < len(s.contents[i]); j++ {
			s.contents[j][i] = g.contents[y+j][x+i]
		}
	}

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

func (s *SubGrid) Valid() bool {
	found := NewNumbers()
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if val, ok := found[s.contents[x][y]]; ok && !val {
				found[s.contents[x][y]] = true
			} else if ok {
				return false
			}
		}
	}
	return true
}

func (s *SubGrid) String() string {
	var str string
	for x := 0; x < len(s.contents); x++ {
		for y := 0; y < len(s.contents[x]); y++ {
			if y != 0 {
				str += " "
			}
			if s.contents[x][y] != 0 {
				str += fmt.Sprintf("%v", s.contents[x][y])
			} else {
				str += string('\u25A1')
			}
		}
		str += "\n"
	}

	return str
}
