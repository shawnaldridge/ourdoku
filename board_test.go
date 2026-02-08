package ourdoku_test

import (
	"testing"
	"ourdoku"
)

func TestPrintBoard(t *testing.T) {
	ourdoku.PrintBoard(ourdoku.EasyBoard)
}

func TestIsBoardValid(t *testing.T) {
	if !ourdoku.IsBoardValid(ourdoku.EasyBoard) {
		t.Errorf("Board is not valid")
	}
}

// add an invalid value in a row

func TestInvalidRow(t *testing.T) {
	board := ourdoku.InvalidBoardRow1
	if ourdoku.IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}

// add an invalid value in a column
func TestInvalidColumn(t *testing.T) {
	board := ourdoku.InvalidBoardColumn1
	if ourdoku.IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}

// add an invalid value in a 3x3 box
func TestInvalidBox(t *testing.T) {
	board := ourdoku.InvalidBoardBox1
	if ourdoku.IsBoardValid(board) {
		t.Errorf("Board is valid but should be invalid")
	}
}
