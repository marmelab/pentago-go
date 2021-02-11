package ai

import (
	"fmt"
	"math/rand"
	"time"
	"game"
	"constants"
)


type Move struct {
	PlaceMarble [3]int
	Rotate [2]string
}

func GetAllPossibleQuadrantRotation(board game.Board, moves []Move, placeMarble [3]int) []Move {
	for rotateQuadrantIndex, _ := range(board.Quadrants) {
		// And for each quadrant, rotate it clockwise and counter clockwise
		for _, rotateKey := range([2]string{constants.ROTATE_CLOCKWISE, constants.ROTATE_COUNTER_CLOCKWISE})  {
			move := Move{
				PlaceMarble: placeMarble,
				Rotate: [2]string{fmt.Sprint(rotateQuadrantIndex), rotateKey},
			}
			moves = append(moves, move)
		}
	}

	return moves
}

func GetAllPossibleMoves(board game.Board) []Move {
	var moves []Move
	for quadrantIndex, quadrant := range(board.Quadrants) { // browse all quadrants
		for rowIndex, row := range(quadrant) {
			for columnIndex, value := range(row) {
				// If it's an empty cell
				if value == constants.EMPTY_CELL_VALUE {
					placeMarble := [3]int{quadrantIndex, rowIndex, columnIndex}
					moves = GetAllPossibleQuadrantRotation(board, moves, placeMarble)
				}
			}
		}
	}
	if constants.RANDOMIZE_MOVES {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(moves), func(i, j int) {
			moves[i], moves[j] = moves[j], moves[i]
		})
	}
	return moves
}

func GetAllPlacements(board game.Board) []Move {
	var moves []Move
	for quadrantIndex, quadrant := range(board.Quadrants) { // browse all quadrants
		for rowIndex, row := range(quadrant) {
			for columnIndex, value := range(row) {
				// If it's an empty cell
				if value == constants.EMPTY_CELL_VALUE {
					placeMarble := [3]int{quadrantIndex, rowIndex, columnIndex}
					move := Move{
						PlaceMarble: placeMarble,
					}
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}
