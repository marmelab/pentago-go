package ai

import (
	"fmt"
	game "game"
)

type Results struct {
	Win []Result
	Loose []Result
}

type Result struct {
	PlaceMarble [2]int
	Rotate [2]string
	Score int
}

func GetResultForHit(board game.Board, quadrantIndex int, rotateKey string, position [3]int) (Result, error) {
	// For a given position and a given quadrant rotation
	// Check if player1 win or loose

	var result Result

	board.Quadrants[quadrantIndex] = game.Rotate(board.Quadrants[quadrantIndex], rotateKey)

	// convert board into string
	boardStringified := game.ToStringBoard(board);

	// Detect winner
	score, err := Score(boardStringified);
	if err != nil {
		return result, err
	}

	// Compute position to be more human readable in results
	boardRelativePositions, err := game.ConvertQuadrantPositionIntoBoardPosition(position)
	if err != nil {
		return result, err
	}

	result = Result{
		PlaceMarble: boardRelativePositions,
		Rotate: [2]string{fmt.Sprint(quadrantIndex + 1), rotateKey},
		Score: score,
	}
	
	return result, err
}

func PlaceMarbleAndMakeAllQuadrantRotations(board game.Board, position [3]int, results []Result) ([]Result, error) {
	// After placing a marble, we need to check every rotation to finalize the turn
	// And be able to detect an alignment.
	// Loop again on quadrants

	board.Quadrants[position[0]][position[1]][position[2]] = "1"

	for quadrantIndex, _ := range(board.Quadrants) {
		// And for each quadrant, rotate it clockwise and counter clockwise
		for _, rotateKey := range([2]string{game.ROTATE_CLOCKWISE, game.ROTATE_COUNTER_CLOCKWISE})  {

			// Get result for this hit
			result, err := GetResultForHit(board, quadrantIndex, rotateKey, position)
			if err != nil {
				return results, err
			}
			
			results = append(results, result)
		}
	}

	return results, nil
}

func PlayAllPossibleMoves(board game.Board) []Result {
	var results []Result

	for quadrantIndex, quadrant := range(board.Quadrants) {
		for rowIndex, row := range(quadrant) {
			for columnIndex, value := range(row) {
				// If it's an empty cell
				if value == "0" {
					results, _ = PlaceMarbleAndMakeAllQuadrantRotations(
						board,
						[3]int{quadrantIndex, rowIndex, columnIndex},
						results,
					)
				}
			}
		}
	}


	return results
}
