package ai

import (
    "testing"
	game "game"
)

func TestGetAllPossibleMoves(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|1|0||1|0|2|
				|0|1|0||2|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|2|0||0|0|2|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := game.DeserializeBoard(rawBoard)
	
	var results []Move
	
	expectedNumberResults := 152
	results = GetAllPossibleMoves(board)
	numberOfResults := len(results)

	if len(results) != numberOfResults {
		t.Errorf("PlayAllPossibleMoves returned %d, expected %d, length is different", numberOfResults, expectedNumberResults)
	}
}
