package ai

import (
    "testing"
	game "game"
)

func TestApplyMoveOnBoard(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|0|0||0|0|0|
				|0|0|0||0|1|0|
				|0|0|0||0|0|0|
				|────────────|
				|0|0|0||0|0|0|
				|0|0|0||2|1|0|
				|1|0|2||1|1|0|
				└────────────┘`

	board, _ := game.DeserializeBoard(rawBoard)

	move := Move{
		PlaceMarble: [3]int{2, 0, 0},
		Rotate: [2]string{"3", "counter clockwise"},
	}
	newBoard := ApplyMoveOnBoard(board, move, "2")

	result := game.ToStringBoard(newBoard)

	expected_result := "000000000010000000200000000011102021"

	if (result != expected_result) {
		t.Errorf("Error ApplyMoveOnBoard : returned %v, expected %v", result, expected_result)
	}
}
