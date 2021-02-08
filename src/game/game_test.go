package game

import (
    "testing"
)

func TestDeserializeBoardShouldReturn2dArray(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    board, err := DeserializeBoard(rawBoard)
	numberOfRows := len(board)
	numberOfColumns := len(board[0])
	if len(board) != 6 || len(board[0]) != 6 || err != nil {
		t.Errorf("The board's size are %d*%d 2 d array (6*6 expected)", numberOfRows, numberOfColumns)
	}
}

func TestDeserializeBoardWithMissingLineShouldThrowAnError(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    _, err := DeserializeBoard(rawBoard)
	if err == nil {
		t.Errorf("The given board should throw an error because second line are missing.")
	}
}

func TestDeserializeBoardWithMissingValueOnLineShouldThrowAnError(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0||
				|0|0|0||0|0|0|
				|0|0|0||0|0|0|
				└────────────┘`

    _, err := DeserializeBoard(rawBoard)
	if err == nil {
		t.Errorf("The given board should throw an error because at line 6, there are only 5 digits")
	}
}
