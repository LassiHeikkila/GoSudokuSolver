package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

// Solution "enum"
const (
	BACKTRACK = "backtrack"
)

var (
	help       = flag.Bool("h", false, "Print help message")
	strategy   = flag.String("strategy", BACKTRACK, "Which strategy to use [backtrack, ...], default backtrack")
	user_input = flag.String("board", "", "Sudoku to solve as a 9x9 json array, 0s representing empty squares")
)

func stringToGrid(s string) *Grid {
	type BoardJson struct {
		Row0 [9]int `json:"row0"`
		Row1 [9]int `json:"row1"`
		Row2 [9]int `json:"row2"`
		Row3 [9]int `json:"row3"`
		Row4 [9]int `json:"row4"`
		Row5 [9]int `json:"row5"`
		Row6 [9]int `json:"row6"`
		Row7 [9]int `json:"row7"`
		Row8 [9]int `json:"row8"`
	}

	var d BoardJson
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return nil
	}

	var contents [9][9]int

	// just copy each row by hand, not a big deal
	copy(contents[0][:], d.Row0[:])
	copy(contents[1][:], d.Row1[:])
	copy(contents[2][:], d.Row2[:])
	copy(contents[3][:], d.Row3[:])
	copy(contents[4][:], d.Row4[:])
	copy(contents[5][:], d.Row5[:])
	copy(contents[6][:], d.Row6[:])
	copy(contents[7][:], d.Row7[:])
	copy(contents[8][:], d.Row8[:])

	return NewGrid(contents)
}

func main() {
	flag.Parse()
	if *help {
		printUsage()
		os.Exit(0)
	}

	if *user_input == "" {
		fmt.Println("Please provide a board to solve!")
		printUsage()
		os.Exit(1)
	}

	g := stringToGrid(*user_input)
	if g == nil {
		fmt.Println("Could not parse input!")
		os.Exit(2)
	}

	fmt.Printf("Input:\n%v\n", g)
	if *strategy == BACKTRACK {
		b := time.Now()
		err := SolveWithBacktracking(g)
		c := time.Since(b)
		if err != nil {
			fmt.Println("Error solving board:", err)
		} else {
			fmt.Printf("Solved:\n%v\n", g)
			fmt.Printf("Solution took %v ms\n", c.Milliseconds())
		}
	} else {
		fmt.Println("Unknown solution strategy requested")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage for:", os.Args[0])
	flag.PrintDefaults()
}
