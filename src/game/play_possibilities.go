package game

import (
	"fmt"
)

type Results struct {
	Win []Result
	Loose []Result
}

type Result struct {
	PlaceMarble [2]int
	Rotate [2]string
}

func GetResultForHit(board Board, quadrantIndex int, rotateKey string, position [3]int) (Result, string, error) {
	// For a given position and a given quadrant rotation
	// Check if player1 win or loose

	var result Result

	board.quadrants[quadrantIndex] = Rotate(board.quadrants[quadrantIndex], rotateKey)

	// convert board into string
	boardStringified := ToStringBoard(board);

	// Detect winner
	status, err := DetectWinner(boardStringified);
	if err != nil {
		return result, "", err
	}

	// If nobody won, we don't need to compute result
	if status == GAME_DRAW || status == GAME_RUNNING {
		return result, status, err
	}

	// Compute position to be more human readable in results
	boardRelativePositions, err := ConvertQuadrantPositionIntoBoardPosition(position)
	if err != nil {
		return result, "", err
	}

	result = Result{
		PlaceMarble: boardRelativePositions,
		Rotate: [2]string{fmt.Sprint(quadrantIndex + 1), rotateKey},
	}
	
	return result, status, err
}

func PlaceMarbleAndMakeAllQuadrantRotations(board Board, position [3]int, results Results) (Results, error) {
	// After placing a marble, we need to check every rotation to finalize the turn
	// And be able to detect an alignment.
	// Loop again on quadrants

	board.quadrants[position[0]][position[1]][position[2]] = "1"

	for quadrantIndex, _ := range(board.quadrants) {
		// And for each quadrant, rotate it clockwise and counter clockwise
		for _, rotateKey := range([2]string{ROTATE_CLOCKWISE, ROTATE_COUNTER_CLOCKWISE})  {

			// Get result for this hit
			result, status, err := GetResultForHit(board, quadrantIndex, rotateKey, position)
			if err != nil {
				return results, err
			}
			// It can be a win or a loose depending on status
			switch status {
			case GAME_PLAYER1_WON:
				results.Win = append(results.Win, result)
			case GAME_PLAYER2_WON:
				results.Loose = append(results.Loose, result)
			}
		}
	}

	return results, nil
}

func PlayAllPossibleMoves(board Board) Results {
	var results Results

	for quadrantIndex, quadrant := range(board.quadrants) {
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
