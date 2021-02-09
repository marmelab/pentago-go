package game

import (
    "testing"
)

func TestPlayAllPossibleMovesWin(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|1|0||1|0|2|
				|0|1|0||2|2|0|
				|────────────|
				|1|0|2||1|2|2|
				|0|2|0||0|0|2|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	
	var results Results
	var expected_results Results
	expected_results.Win = []Result{
		Result{
			PlaceMarble: [2]int{1, 4},
			Rotate: [2]string{"1", "clockwise"},
		},
		Result{
			PlaceMarble: [2]int{1, 4},
			Rotate: [2]string{"1", "counter clockwise"},
		},
	}
	results = PlayAllPossibleMoves(board)

	if len(results.Win) != len(expected_results.Win) {
		t.Errorf("PlayAllPossibleMoves returned %v, expected %v, length is different", results, expected_results)
	}

	if results.Win[0] != expected_results.Win[0] {
		t.Errorf("PlayAllPossibleMoves returned %v, expected %v, first item is different", results, expected_results)
	}

	if results.Win[1] != expected_results.Win[1] {
		t.Errorf("PlayAllPossibleMoves returned %v, expected %v second item is different", results, expected_results)
	}
}

func TestPlayAllPossibleMovesLoose(t *testing.T) {
	rawBoard := `┌────────────┐
				|0|1|1||0|2|1|
				|0|0|0||1|0|2|
				|0|1|0||2|2|2|
				|────────────|
				|1|0|2||1|2|2|
				|0|2|0||0|0|2|
				|0|0|0||0|0|0|
				└────────────┘`

	board, _ := DeserializeBoard(rawBoard)
	
	var results Results
	expected_result := Result{
		PlaceMarble: [2]int{0, 0},
		Rotate: [2]string{"2", "counter clockwise"},
	}

	results = PlayAllPossibleMoves(board)

	numberOfWins := len(results.Win)
	numberOfLoose := len(results.Loose)

	firstLoose := results.Loose[0]

	if numberOfWins != 0 {
		t.Errorf("PlayAllPossibleMoves returned %v win, expected 0 win, length is different", numberOfWins)
	}

	if numberOfLoose != 19 {
		t.Errorf("PlayAllPossibleMoves returned %v loose, expected 18 win, length is different", numberOfLoose)
	}

	if firstLoose != expected_result {
		t.Errorf("PlayAllPossibleMoves returned %v as first item, expected %v, first item is different", firstLoose, expected_result)
	}
}
