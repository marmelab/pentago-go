package game

import (
    "testing"
	"fmt"
)

var detectWinnerDatasets = []struct {
	in  string
	out string
}{
	{`
	┌────────────┐
	|1|1|1||1|1|0|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	└────────────┘`, "GAME_PLAYER1_WON"},
	{`
	┌────────────┐
	|1|0|1||1|1|0|
	|0|1|0||0|0|0|
	|0|0|1||0|0|0|
	|────────────|
	|0|0|0||1|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	└────────────┘`, "GAME_PLAYER1_WON"},
	{`
	┌────────────┐
	|1|1|1||1|1|0|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|2|2|2||2|2|0|
	└────────────┘`, "GAME_DRAW"},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, "GAME_PLAYER2_WON"},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|1|0||1|0|0|
	|────────────|
	|0|1|1||0|0|0|
	|0|1|0||0|0|0|
	|1|1|0||1|1|0|
	└────────────┘`, "GAME_PLAYER1_WON"},
	{`
	┌────────────┐
	|0|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, "GAME_RUNNING"},
	{`
	┌────────────┐
	|1|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, "GAME_DRAW"},
}

func TestDetectWinner(t *testing.T) {
	for _, data := range detectWinnerDatasets {
		board, err := DeserializeBoard(data.in)
		boardStringified := ToStringBoard(board)
		if err != nil {
			t.Errorf("Error thrown during deserialization")
		}

		result, err := DetectWinner(boardStringified)

		if err != nil {
			t.Errorf("Error thrown during winner detection")
		}
		if result != data.out {
			fmt.Println(data.in)
			t.Errorf("Error : got %v, want %v", result, data.out)
		}
	}
}
