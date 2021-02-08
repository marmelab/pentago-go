package game

import (
	"strconv"
	"errors"
)

const GAME_RUNNING = "GAME_RUNNING"
const GAME_DRAW = "GAME_DRAW"
const GAME_PLAYER1_WON = "GAME_PLAYER1_WON"
const GAME_PLAYER2_WON = "GAME_PLAYER2_WON"



func DetectWinner(board[][]int) (string, error) {
	// This function will compare by bitwise "and" operator the played value by the player
	// and all combinations possibles (converted in int64)
	// There are 32 combinations possibles.

	var player1BoardString string = ""
	var player2BoardString string = ""

	// Will be used to detect if a board is full and contain 1 if cell is played, 0 if it's empty
	var bothBoardString string = ""

	// For each cell in board we need to fill a strings with 1 and 0 (binary representation of the board)
	// We ignore other player values.
	for _, row := range board {
		for _, value := range row {
			if value == 1 {
				player1BoardString = player1BoardString + "1"
			} else {
				player1BoardString = player1BoardString + "0"
			}

			if value == 2 {
				player2BoardString = player2BoardString + "1"
			} else {
				player2BoardString = player2BoardString + "0"
			}

			if value != 0 {
				bothBoardString = bothBoardString + "1"
			}


		}
	}

	bothBoardInt64, err := strconv.ParseInt(bothBoardString, 2, 64)
	if err != nil {
		return "", errors.New("Could'nt parse correctly both board")
	}

	// Convert this string into int64.
	player1Int64, err := strconv.ParseInt(player1BoardString, 2, 64)
	if err != nil {
		return "", errors.New("Could'nt parse correctly player 1 board")
	}

	player2Int64, err := strconv.ParseInt(player2BoardString, 2, 64)
	if err != nil {
		return "", errors.New("Could'nt parse correctly player 2 board")
	}

	gameResult := GAME_RUNNING
	// Get all combinations and use binary comparaison.
	for _, combination := range GetAllCombinations() {
		player1win := BinaryCompareInt64(combination, player1Int64)
		player2win := BinaryCompareInt64(combination, player2Int64)
		
		if player1win == true {
			if gameResult == GAME_PLAYER2_WON {
				return GAME_DRAW, nil
			} else {
				gameResult = GAME_PLAYER1_WON
			}
		}

		if player2win == true {
			if gameResult == GAME_PLAYER1_WON {
				return GAME_DRAW, nil
			} else {
				gameResult = GAME_PLAYER2_WON
			}
		}

	}

	if gameResult == GAME_RUNNING && IsBoardFull(bothBoardInt64) {
		return GAME_DRAW, nil
	}

	return gameResult, nil
}

