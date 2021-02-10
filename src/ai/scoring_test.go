package ai

import (
    "testing"
	"fmt"
	game "game"
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

func TestEvaluateGameStatus(t *testing.T) {
	for _, data := range detectWinnerDatasets {
		board, err := game.DeserializeBoard(data.in)
		boardStringified := game.ToStringBoard(board)
		player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStringified)

		if err != nil {
			t.Errorf("Error thrown during deserialization")
		}

		result, _, err := EvaluateGameStatus(player1Int64, player2Int64, "1")

		if err != nil {
			t.Errorf("Error thrown during winner detection")
		}
		if result != data.out {
			fmt.Println(data.in)
			t.Errorf("Error EvaluteGameStatus: got %v, want %v", result, data.out)
		}
	}
}


var scoreDatasets = []struct {
	in  string
	currentPlayer string
	out int
}{
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, "1", -1107},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, "2", -1107},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|1|0||1|0|0|
	|────────────|
	|0|1|2||0|0|0|
	|0|1|0||0|0|0|
	|1|1|0||1|1|0|
	└────────────┘`, "1", 2250},
	{`
	┌────────────┐
	|0|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, "1", 0},
	{`
	┌────────────┐
	|1|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, "1", 0},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||2|1|0|
	|1|0|0||1|1|0|
	└────────────┘`, "1", 236},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||2|1|0|
	|1|0|2||1|1|0|
	└────────────┘`, "1", 126},
}
func TestEvaluateScore(t *testing.T) {
	for _, data := range scoreDatasets {
		board, _ := game.DeserializeBoard(data.in)
		boardStr := game.ToStringBoard(board)

		player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStr)

		result, _ := EvaluateScore(player1Int64, player2Int64, data.currentPlayer)
		if (result != data.out) {
			t.Errorf("Error EvaluateScore : returned %d, expected %d", result, data.out)
		}
	}
}
