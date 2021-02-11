package ai

import (
	"strconv"
	"math"
	"game"
	"constants"
)

func switchPlayer(currentPlayer string) string {
	if currentPlayer == constants.PLAYER1_VALUE {
		return constants.PLAYER2_VALUE
	}

	return constants.PLAYER1_VALUE
}

func ApplyMoveOnBoard(board game.Board, move Move, currentPlayer string) game.Board {
	placeMarble := move.PlaceMarble
	board.Quadrants[placeMarble[0]][placeMarble[1]][placeMarble[2]] = currentPlayer
	
	rotateQuadrant := move.Rotate
	quadrantIndex, _ := strconv.Atoi(rotateQuadrant[0])

	board.Quadrants[quadrantIndex] = game.Rotate(board.Quadrants[quadrantIndex], rotateQuadrant[1])
	return board;
}

func Minimax(depth int, board game.Board, currentPlayer string, move Move, alpha int, beta int) (int, Move) {
	// Get player1 and player2 binaries representation of their marbles
	boardStringified := game.ToStringBoard(board)
	
	player1Int64, player2Int64, _ := GetPlayerBoardsFromBoard(boardStringified)

	gameStatus, score, _ := EvaluateGameStatus(player1Int64, player2Int64)

	if gameStatus != constants.GAME_RUNNING {
		// fmt.Println(boardStringified, gameStatus, score)
		return score, move
	}

	if depth == 0 {
		score, _ := EvaluateScore(player1Int64, player2Int64)
		return score, move
	}

	moves := GetAllPossibleMoves(board);

	var bestMove Move
	var bestScore int
	if currentPlayer == constants.PLAYER1_VALUE {
		bestScore = - constants.SCORE_ALIGNED[4]
	} else {
		bestScore =  constants.SCORE_ALIGNED[4]
	}


	for _, move := range moves {
		newBoard := ApplyMoveOnBoard(board, move, currentPlayer)
		opponent := switchPlayer(currentPlayer)
		childScore, _ := Minimax(depth - 1, newBoard, opponent, move, alpha, beta)

		if currentPlayer == constants.PLAYER1_VALUE && bestScore < childScore {
			bestScore = childScore
			bestMove = move
			alpha = int(math.Max(float64(alpha), float64(bestScore)))
			if beta <= alpha  {
				break
			}
		} else if currentPlayer == constants.PLAYER2_VALUE && bestScore > childScore {
			bestScore = childScore
			bestMove = move
			beta = int(math.Min(float64(beta), float64(bestScore)))
			if beta <= alpha {
				break
			}
		}
	}
	return bestScore, bestMove 
}
