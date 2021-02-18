package ai

import (
	"strconv"
	"strings"
	"errors"
	"constants"
)




func GetPlayerBoardsFromBoard(board string) (int64, int64, error) {

	var player1BoardString string = ""
	var player2BoardString string = ""

	// For each cell in board we need to fill a strings with 1 and 0 (binary representation of the board)
	// We ignore other player values.

	player1BoardString = strings.ReplaceAll(board, "2", "0");

	player2BoardString = strings.ReplaceAll(board, "1", "0");
	player2BoardString = strings.ReplaceAll(player2BoardString, "2", "1");

	// Convert this string into int64.
	player1Int64, err := strconv.ParseInt(player1BoardString, 2, 64)
	if err != nil {
		return 0, 0, errors.New("Cannot parse correctly player 1 board")
	}

	player2Int64, err := strconv.ParseInt(player2BoardString, 2, 64)
	if err != nil {
		return 0, 0, errors.New("Cannot parse correctly player 2 board")
	}

	return player1Int64, player2Int64, nil
}

func DetectWinner(player1Int64 int64, player2Int64 int64) (string, error) {
	// This function will compare by bitwise "and" operator the played value by the player
	// and all combinations possibles (converted in int64)
	// There are 32 combinations possibles.

	gameResult := constants.GAME_RUNNING

	// Get all combinations and use binary comparaison.
	for _, combination := range ALL_COMBINATIONS_OF_WIN {
		
		// We don't need to check other combinations if player 1 has already an alignment
		if gameResult != constants.GAME_PLAYER1_WON {
			player1win := BinaryCompareInt64(combination, player1Int64)
			if player1win == true {
				// If we previously detect that player2 won, it's a draw
				if gameResult == constants.GAME_PLAYER2_WON {
					return constants.GAME_DRAW, nil
				} else {
					gameResult = constants.GAME_PLAYER1_WON
				}
			}
		}

		// We don't need to check other combinations if player 2 has already an alignment
		if gameResult != constants.GAME_PLAYER2_WON {
			player2win := BinaryCompareInt64(combination, player2Int64)
			if player2win == true {
				// If we previously detect that player1 won, it's a draw
				if gameResult == constants.GAME_PLAYER1_WON {
					return constants.GAME_DRAW, nil
				} else {
					gameResult = constants.GAME_PLAYER2_WON
				}
			}
		}	

	}

	if gameResult == constants.GAME_RUNNING && IsBoardFull(player1Int64, player2Int64) {
		return constants.GAME_DRAW, nil
	}

	return gameResult, nil
}

func EvaluateGameStatus(player1Int64 int64, player2Int64 int64, firstPlayer string) (string, int, error) {

	winStatus, err := DetectWinner(player1Int64, player2Int64)
	score := 0
	if err != nil {
		return "", score, err
	}

	switch true {
	case winStatus == constants.GAME_PLAYER1_WON && firstPlayer == constants.PLAYER1_VALUE:
		score = constants.SCORE_ALIGNED[4]
	case winStatus == constants.GAME_PLAYER2_WON && firstPlayer == constants.PLAYER2_VALUE:
		score = constants.SCORE_ALIGNED[4]
	case winStatus == constants.GAME_PLAYER1_WON && firstPlayer != constants.PLAYER1_VALUE:
		score = -constants.SCORE_ALIGNED[4]
	case winStatus == constants.GAME_PLAYER2_WON && firstPlayer != constants.PLAYER2_VALUE:
		score = -constants.SCORE_ALIGNED[4]
	}

	return winStatus, score, nil
}

func EvaluateAllCombinationsOfWin(playerInt64 int64, opponentInt64 int64) int {
	score := 0
	// Get all combinations and use binary comparaison.
	for _, combination := range ALL_COMBINATIONS_OF_WIN {
		marblesAligned := CountBitsForCombinationIfStillPossible(combination, playerInt64, opponentInt64)

		// 0 means no marbles are in this combinaton or opponent already countered this combination.
		if marblesAligned != 0 {
			score = score +  constants.SCORE_ALIGNED[marblesAligned - 1]
		}

		marblesAligned = CountBitsForCombinationIfStillPossible(combination, opponentInt64, playerInt64)
		if marblesAligned != 0 {
			score = score -  constants.SCORE_ALIGNED[marblesAligned - 1]
		}
	}

	return score
}

func EvaluateCenters(playerInt64 int64, opponentInt64 int64) int {
	score := 0

	for _, combination := range MIDDLE_QUADRANTS {
		marblesAligned := CountBitsForCombinationIfStillPossible(combination, playerInt64, opponentInt64)
		
		// 0 means no marbles are in this combinaton or opponent already countered this combination.
		if marblesAligned != 0 {
			score = score +  constants.SCORE_CENTER
		}

		marblesAligned = CountBitsForCombinationIfStillPossible(combination, opponentInt64, playerInt64)
		if marblesAligned != 0 {
			score = score -  constants.SCORE_CENTER
		}
	}
	return score

}

func EvaluateScore(player1Int64 int64, player2Int64 int64, firstPlayer string) int {

	score := 0

	var playerInt64, opponentInt64 int64

	if firstPlayer == constants.PLAYER1_VALUE {
		playerInt64 = player1Int64
		opponentInt64 = player2Int64
	} else {
		playerInt64 = player2Int64
		opponentInt64 = player1Int64
	}

	score = score + EvaluateCenters(playerInt64, opponentInt64)
	score = score + EvaluateAllCombinationsOfWin(playerInt64, opponentInt64)
	

	return score
}
