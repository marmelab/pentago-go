package ai

import (
    "testing"
	"fmt"
	game "game"
)


var getPlayerBoardsFromBoardDatasets = []struct {
	in  string
	expectedPlayer1 int64
	expectedPlayer2 int64
}{
	{`
		┌────────────┐
		|1|1|1||1|1|0|
		|0|0|0||1|0|0|
		|0|2|0||2|0|0|
		|────────────|
		|0|2|0||0|0|0|
		|0|2|0||1|0|0|
		|0|0|1||0|0|0|
		└────────────┘`,
		0b111110_000100_000000_000000_000100_001000,
		0b000000_000000_010100_010000_010000_000000,
	},
}

func TestGetPlayerBoardsFromBoard(t *testing.T) {
	for _, data := range getPlayerBoardsFromBoardDatasets {
		board, _ := game.DeserializeBoard(data.in)
		boardStringified := game.ToStringBoard(board)
		player1, player2, err := GetPlayerBoardsFromBoard(boardStringified)

		if err != nil {
			t.Errorf("Error GetPlayerBoardsFromBoard throw an error")

		}

		if player1 != data.expectedPlayer1 {
			t.Errorf("Error GetPlayerBoardsFromBoard player1 returned %d, expected %d", player1, data.expectedPlayer1)
		}

		if player2 != data.expectedPlayer2 {
			t.Errorf("Error GetPlayerBoardsFromBoard player2 returned %d, expected %d", player1, data.expectedPlayer1)
		}
	}
}

var detectWinnerDatasets = []struct {
	in  string
	out int
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
	└────────────┘`, 10000},
	{`
	┌────────────┐
	|1|0|1||1|1|0|
	|0|1|0||0|0|0|
	|0|0|1||0|0|0|
	|────────────|
	|0|0|0||1|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	└────────────┘`, 10000},
	{`
	┌────────────┐
	|1|1|1||1|1|0|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|2|2|2||2|2|0|
	└────────────┘`, 0},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, -10000},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|1|0||1|0|0|
	|────────────|
	|0|1|1||0|0|0|
	|0|1|0||0|0|0|
	|1|1|0||1|1|0|
	└────────────┘`, 10000},
	{`
	┌────────────┐
	|0|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, 0},
	{`
	┌────────────┐
	|1|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, 0},
}

func TestEvaluateGameStatus(t *testing.T) {
	for _, data := range detectWinnerDatasets {
		board, err := game.DeserializeBoard(data.in)
		boardStringified := game.ToStringBoard(board)
		player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStringified)

		if err != nil {
			t.Errorf("Error thrown during deserialization")
		}

		result, score, err := EvaluateGameStatus(player1Int64, player2Int64, "1")

		if err != nil {
			t.Errorf("Error thrown during winner detection")
		}
		if score != data.out {
			fmt.Println(data.in)
			t.Errorf("Error EvaluteGameStatus: got %v, %v, want %v",result, score, data.out)
		}
	}
}


var evaluateAllCombinationsOfWinDatasets = []struct {
	in  string
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
	└────────────┘`, -1107},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, -1107},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|1|0||1|0|0|
	|────────────|
	|0|1|2||0|0|0|
	|0|1|0||0|0|0|
	|1|1|0||1|1|0|
	└────────────┘`, 2248},
	{`
	┌────────────┐
	|0|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, 0},
	{`
	┌────────────┐
	|1|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|2|2|
	|1|1|1||1|2|2|
	└────────────┘`, 0},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||2|1|0|
	|1|0|0||1|1|0|
	└────────────┘`, 235},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|1|0|
	|0|0|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|0|0||2|1|0|
	|1|0|2||1|1|0|
	└────────────┘`, 124},
}

func TestEvaluateAllCombinationsOfWin(t *testing.T) {
	for _, data := range evaluateAllCombinationsOfWinDatasets {
		board, _ := game.DeserializeBoard(data.in)
		boardStr := game.ToStringBoard(board)

		player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStr)

		result := EvaluateAllCombinationsOfWin(player1Int64, player2Int64)
		if (result != data.out) {
			t.Errorf("Error EvaluateAllCombinationsOfWin : returned %d, expected %d", result, data.out)
		}
	}
}

var evaluateCentersDatasets = []struct {
	in  string
	out int
}{
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, -5},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|0|2|0||0|0|0|
	|────────────|
	|0|0|0||0|0|0|
	|0|2|0||0|0|0|
	|1|2|0||1|1|0|
	└────────────┘`, -10},
	{`
	┌────────────┐
	|0|0|0||0|0|0|
	|0|2|0||0|1|0|
	|0|1|0||1|0|0|
	|────────────|
	|0|1|2||0|0|0|
	|0|1|0||0|2|0|
	|1|1|0||1|1|0|
	└────────────┘`, 0},
	{`
	┌────────────┐
	|0|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|2|2||1|1|2|
	|1|1|1||1|2|2|
	└────────────┘`, 10},
	{`
	┌────────────┐
	|1|2|2||1|2|2|
	|1|1|2||2|1|1|
	|1|1|2||1|1|1|
	|────────────|
	|2|1|1||2|2|2|
	|1|1|2||1|1|2|
	|1|1|1||1|2|2|
	└────────────┘`, 20},
}


func TestEvaluateCenters(t *testing.T) {
	for _, data := range evaluateCentersDatasets {
		board, _ := game.DeserializeBoard(data.in)
		boardStr := game.ToStringBoard(board)

		player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStr)

		result := EvaluateCenters(player1Int64, player2Int64)
		if (result != data.out) {
			t.Errorf("Error EvaluateAllCombinationsOfWin : returned %d, expected %d", result, data.out)
		}
	}
}
