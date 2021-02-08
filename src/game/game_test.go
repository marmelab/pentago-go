package game

import (
    "testing"
)

func TestDeserialize(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    board := DeserializeBoard(rawBoard)
	numberOfRows := len(board)
	numberOfColumns := len(board[0])
	if len(board) != 6 || len(board[0]) != 6 {
		t.Errorf("The board's size are %d*%d 2 d array (6*6 expected)", numberOfRows, numberOfColumns)
	}
}
