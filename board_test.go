package main

import (
	"testing"
)

func TestPrintBoard(t *testing.T) {
	PrintBoard(EasyBoard)
}

func TestIsBoardValid(t *testing.T) {
	if !IsBoardValid(EasyBoard) {
		t.Errorf("Board is not valid")
	}
}

// add an invalid value in a row

func TestInvalidRow(t *testing.T) {
	board := InvalidBoardRow1
	if IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}

// add an invalid value in a column
func TestInvalidColumn(t *testing.T) {
	board := InvalidBoardColumn1
	if IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}

// add an invalid value in a 3x3 box
func TestInvalidBox(t *testing.T) {
	board := InvalidBoardBox1
	if IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}
