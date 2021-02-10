package ai

import (
	"game"
	"strconv"
)

func switchPlayer(currentPlayer string) string {
	if currentPlayer == "1" {
		return "2"
	}

	return "1"
}

func ApplyMoveOnBoard(board game.Board, move Move, currentPlayer string) game.Board {
	placeMarble := move.PlaceMarble
	board.Quadrants[placeMarble[0]][placeMarble[1]][placeMarble[2]] = currentPlayer
	
	rotateQuadrant := move.Rotate
	quadrantIndex, _ := strconv.Atoi(rotateQuadrant[0])

	board.Quadrants[quadrantIndex] = game.Rotate(board.Quadrants[quadrantIndex], rotateQuadrant[1])
	return board;
}

func Minimax(depth int, board game.Board, currentPlayer string, move Move) (int, Move) {
	// Get player1 and player2 binaries representation of their marbles
	boardStringified := game.ToStringBoard(board)
	
	player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStringified)

	gameStatus, score, _ := EvaluateGameStatus(player1Int64, player2Int64, currentPlayer)

	if gameStatus != GAME_RUNNING {
		return score, move
	}

	if depth == 0 {
		score, _ := EvaluateScore(player1Int64, player2Int64, currentPlayer)
		return score, move
	}

	moves := GetAllPossibleMoves(board);

	var bestMove Move
	var bestScore int
	if currentPlayer == "1" {
		bestScore = -SCORE_ALIGNED[4]
	} else {
		bestScore = SCORE_ALIGNED[4]
	}

	for _, move := range moves {
		newBoard := ApplyMoveOnBoard(board, move, currentPlayer)
		opponent := switchPlayer(currentPlayer)
		childScore, childMove := Minimax(depth - 1, newBoard, opponent, move)

		if currentPlayer == "1" && bestScore < childScore {
			bestScore = childScore
			bestMove = childMove

			// Alpha beta pruning should be here
		} else if currentPlayer == "2" && bestScore > childScore {
			bestScore = childScore
			bestMove = childMove
			// Alpha beta pruning should be here

		}	
	}

	return bestScore, bestMove 
}
