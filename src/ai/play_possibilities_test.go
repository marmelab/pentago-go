package ai

import (
    "testing"
	game "game"
)

func TestPlayAllPossibleMoves(t *testing.T) {
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
	
	var results []Result
	var expected_results []Result
	expected_results = []Result{
		Result{
			PlaceMarble: [2]int{0, 0},
			Rotate: [2]string{"1", "clockwise"},
			Score: 122,
		},
		Result{
			PlaceMarble: [2]int{1, 4},
			Rotate: [2]string{"1", "counter clockwise"},
		},
	}
	numberOfResults := 152

	results = PlayAllPossibleMoves(board)
	if len(results) != numberOfResults {
		t.Errorf("PlayAllPossibleMoves returned %v, expected %v, length is different", results, expected_results)
	}
}
