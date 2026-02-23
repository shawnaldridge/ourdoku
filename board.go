package main

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

var MediumBoard = [9][9]int{
	{0, 4, 0, 8, 1, 7, 6, 0, 2},
	{0, 0, 6, 0, 0, 4, 1, 0, 5},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{4, 7, 0, 0, 2, 0, 0, 6, 8},
	{0, 0, 0, 5, 8, 9, 0, 0, 3},
	{0, 0, 0, 7, 0, 6, 9, 0, 1},
	{0, 0, 0, 6, 3, 8, 7, 1, 9},
	{0, 0, 7, 9, 0, 0, 2, 0, 4},
	{0, 1, 0, 4, 7, 0, 8, 5, 0},
}

var HardBoard = [9][9]int{
	{1, 0, 0, 0, 3, 4, 0, 0, 8},
	{0, 7, 0, 6, 8, 0, 0, 3, 0},
	{0, 0, 8, 2, 1, 0, 7, 0, 4},
	{0, 5, 4, 0, 9, 0, 6, 8, 0},
	{9, 1, 0, 5, 0, 8, 0, 2, 0},
	{0, 8, 0, 3, 0, 0, 0, 0, 5},
	{3, 0, 5, 9, 0, 6, 8, 7, 1},
	{0, 0, 6, 0, 0, 0, 0, 4, 0},
	{0, 0, 1, 0, 7, 0, 2, 0, 0},
}

var ExpertBoard = [9][9]int{
	{1, 5, 0, 0, 8, 2, 0, 0, 0},
	{3, 0, 0, 0, 7, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 7, 5, 3},
	{0, 0, 0, 5, 2, 7, 6, 0, 9},
	{0, 0, 0, 0, 0, 0, 5, 0, 0},
	{0, 4, 0, 0, 6, 3, 8, 0, 7},
	{4, 0, 0, 0, 0, 8, 0, 0, 0},
	{7, 0, 3, 0, 4, 0, 1, 0, 0},
	{0, 0, 8, 6, 0, 0, 3, 0, 0},
}

var MasterBoard = [9][9]int{
	{0, 0, 9, 5, 8, 6, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	{4, 0, 0, 0, 0, 0, 6, 8, 3},
	{9, 0, 0, 6, 5, 0, 0, 3, 2},
	{0, 6, 0, 7, 0, 0, 0, 9, 8},
	{0, 3, 0, 2, 0, 0, 7, 0, 4},
	{0, 0, 3, 0, 0, 0, 0, 0, 0},
	{6, 2, 0, 0, 1, 5, 0, 4, 0},
	{0, 0, 0, 4, 0, 0, 0, 5, 0},
}

var ExtremeBoard = [9][9]int{
	{3, 0, 0, 0, 4, 9, 0, 0, 0},
	{0, 0, 0, 6, 0, 0, 5, 0, 1},
	{7, 5, 2, 0, 0, 1, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 7, 0, 0},
	{5, 0, 0, 3, 9, 6, 0, 0, 0},
	{0, 0, 8, 1, 5, 0, 0, 9, 6},
	{0, 0, 3, 0, 1, 0, 0, 6, 0},
	{0, 0, 4, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 2, 8, 0, 0, 0},
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

func FindFirstEmptyCell(board [9][9]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func solveBacktrack(board [9][9]int) (bool, [9][9]int, int) {
	executions := 1
	emptyCellRow, emptyCellColumn := FindFirstEmptyCell(board)

	if emptyCellRow == -1 || emptyCellColumn == -1 {
		return true, board, executions
	}

	for i := 1; i <= 9; i++ {
		board[emptyCellRow][emptyCellColumn] = i
		if IsBoardValid(board) {
			solved, solvedBoard, subExecs := solveBacktrack(board)
			executions += subExecs
			if solved {
				return true, solvedBoard, executions
			}
			board[emptyCellRow][emptyCellColumn] = 0
		}
	}

	return false, board, executions
}

func SolveViaBacktracking(board [9][9]int) (bool, [9][9]int, int) {
	solved, solvedBoard, executions := solveBacktrack(board)
	if !solved {
		return false, solvedBoard, -1
	}
	return true, solvedBoard, executions
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