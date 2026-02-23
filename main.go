package main

import (
	"fmt"
)

func main() {
	IsBoardValid(EasyBoard)

	solved, solvedBoard := SolveViaBacktracking(EasyBoard)
	if solved {
		PrintBoard(solvedBoard)
	} else {
		fmt.Println("No solution found")
	}
}

