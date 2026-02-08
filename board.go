package ourdoku

import (
	"fmt"
	"strconv"
)

var EasyBoard = [9][9]int{
	{9, 0, 0, 5, 0, 8, 0, 0, 7},
	{0, 8, 0, 3, 0, 2, 9, 0, 5},
	{0, 5, 4, 0, 0, 0, 0, 8, 0},
	{0, 7, 0, 6, 8, 0, 0, 3, 2},
	{1, 0, 0, 0, 0, 4, 0, 0, 8},
	{5, 0, 0, 2, 1, 9, 0, 6, 0},
	{0, 0, 0, 9, 0, 6, 0, 0, 1},
	{7, 2, 6, 0, 0, 1, 0, 4, 0},
	{0, 0, 1, 4, 7, 0, 0, 5, 6},
}

var EmptyBoard = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},	
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var InvalidBoardRow1 = [9][9]int{
	{9, 9, 0, 5, 0, 8, 0, 0, 7},
	{0, 8, 0, 3, 0, 2, 9, 0, 5},
	{0, 5, 4, 0, 0, 0, 0, 8, 0},
	{0, 7, 0, 6, 8, 0, 0, 3, 2},
	{1, 0, 0, 0, 0, 4, 0, 0, 8},
	{5, 0, 0, 2, 1, 9, 0, 6, 0},
	{0, 0, 0, 9, 0, 6, 0, 0, 1},
	{7, 2, 6, 0, 0, 1, 0, 4, 0},
	{0, 0, 1, 4, 7, 0, 0, 5, 6},
}

var InvalidBoardColumn1 = [9][9]int{
	{9, 0, 0, 5, 0, 8, 0, 0, 7},
	{1, 8, 0, 3, 0, 2, 9, 0, 5},
	{0, 5, 4, 0, 0, 0, 0, 8, 0},
	{0, 7, 0, 6, 8, 0, 0, 3, 2},
	{1, 0, 0, 0, 0, 4, 0, 0, 8},
	{5, 0, 0, 2, 1, 9, 0, 6, 0},
	{0, 0, 0, 9, 0, 6, 0, 0, 1},
	{7, 2, 6, 0, 0, 1, 0, 4, 0},
	{0, 0, 1, 4, 7, 0, 0, 5, 6},
}

var InvalidBoardBox1 = [9][9]int{
	{9, 0, 0, 5, 0, 8, 0, 0, 7},
	{0, 4, 0, 3, 0, 2, 9, 0, 5},
	{0, 5, 4, 0, 0, 0, 0, 8, 0},
	{0, 7, 0, 6, 8, 0, 0, 3, 2},
	{1, 0, 0, 0, 0, 4, 0, 0, 8},
	{5, 0, 0, 2, 1, 9, 0, 6, 0},
	{0, 0, 0, 9, 0, 6, 0, 0, 1},
	{7, 2, 6, 0, 0, 1, 0, 4, 0},
	{0, 0, 1, 4, 7, 0, 0, 5, 6},
}

func PrintBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Print(strconv.Itoa(board[i][j]) + " ")
			}
		}
		fmt.Println()
	}
}

func IsBoardValid(board [9][9]int) bool {
	// check rows
	for i := 0; i < 9; i++ {
		var rowSeen = make(map[int]bool)
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == 0 {
				continue
			}
			if rowSeen[v] {
				return false
			}
			rowSeen[v] = true
		}
	}
	// check columns
	for j := 0; j < 9; j++ {
		var columnSeen = make(map[int]bool)
		for i := 0; i < 9; i++ {
			v := board[i][j]
			if v == 0 {
				continue
			}
			if columnSeen[v] {
				return false
			}
			columnSeen[v] = true
		}
	}
	// check 3x3 boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var boxSeen = make(map[int]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					v := board[i+k][j+l]
					if v == 0 {
						continue
					}
					if boxSeen[v] {
						return false
					}
					boxSeen[v] = true
				}
			}
		}
	}
	return true
}