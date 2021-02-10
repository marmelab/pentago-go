package ai

import (
	"fmt"
	game "game"
)


type Move struct {
	PlaceMarble [3]int
	Rotate [2]string
}

func GetAllPossibleQuadrantRotation(board game.Board, moves []Move, placeMarble [3]int) []Move {
	for rotateQuadrantIndex, _ := range(board.Quadrants) {
		// And for each quadrant, rotate it clockwise and counter clockwise
		for _, rotateKey := range([2]string{game.ROTATE_CLOCKWISE, game.ROTATE_COUNTER_CLOCKWISE})  {
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
				if value == "0" {
					placeMarble := [3]int{quadrantIndex, rowIndex, columnIndex}
					moves = GetAllPossibleQuadrantRotation(board, moves, placeMarble)
				}
			}
		}
	}

	return moves
}
