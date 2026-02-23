package main

import (
	"fmt"
)

func main() {
	boards := []struct {
		name  string
		board [9][9]int
	}{
		{"Easy", EasyBoard},
		{"Medium", MediumBoard},
		{"Hard", HardBoard},
		{"Expert", ExpertBoard},
		{"Master", MasterBoard},
		{"Extreme", ExtremeBoard},
	}

	for _, b := range boards {
		solved, _, executions := SolveViaBacktracking(b.board)
		if solved {
			fmt.Printf("%s: solved in %d executions\n", b.name, executions)
		} else {
			fmt.Printf("%s: no solution found (%d executions)\n", b.name, executions)
		}
	}
}

