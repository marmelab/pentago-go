package game

import (
	"strings"
	"regexp"
	"errors"
)

type Board struct {
	Quadrants [4][3][3]string
}

const ROTATE_CLOCKWISE = "clockwise"
const ROTATE_COUNTER_CLOCKWISE = "counter clockwise"

func DeserializeBoard(str string) (Board, error) {
	var board Board

	// Split string into array based on \n
	strArr := strings.Fields(str)
	// Remove first line, last line & middle line which are board's delimiters.
	if len(strArr) != 9 {
		return board, errors.New("Given string is not correctly formatted and have to contain 9 lines.")
	}
	

	arrayWithoutBorders := append(make([]string, 0, 6), strArr[1:4]...)
	arrayWithoutBorders = append(arrayWithoutBorders, strArr[5:8]...)

	re := regexp.MustCompile("[0-9]+")

	for x, line := range arrayWithoutBorders {
		// Remove all non-digits chars
		lineStrings := re.FindAllString(line, -1)

		if len(lineStrings) != 6 {
			return board, errors.New("All line should contains 6 values.")
		}

		// This loop allow us to cast all digits into integers

		for y, value := range lineStrings {
			switch true {
			case x < 3 && y < 3:
				board.Quadrants[0][x][y] = value
			case x < 3 && y >= 2:
				board.Quadrants[1][x][y - 3] = value
			case x >= 3 && y < 3:
				board.Quadrants[2][x - 3][y] = value
			case x >= 3 && y >= 3:
				board.Quadrants[3][x - 3][y - 3] = value
			}
		
		}

	}

	// Return our Board object
	return board, nil
}

func GetLineOfBoard(board Board, line int) string {

	quadrantsLine := line % 3
	quadrantNumber := 0
	if line >= 3 {
		quadrantNumber = 2
	}
	return strings.Join(board.Quadrants[quadrantNumber][quadrantsLine][:], "") + 
		strings.Join(board.Quadrants[quadrantNumber + 1][quadrantsLine][:], "")
}

func ToStringBoard(board Board) string {
	boardStringified := ""

	for i := 0; i < 6; i++ {
		boardStringified = boardStringified + GetLineOfBoard(board, i)
	}

	return boardStringified
}

func ConvertQuadrantPositionIntoBoardPosition(position [3]int) ([2]int, error) {
	// Given a position based on [quadrantIndex, x, y]
	// Return a position based on a 2d matrix.
	// e.g [2, 0, 2] should return [3, 2].

	switch position[0] {
	case 0:
		return [2]int{position[1], position[2]}, nil
	case 1:
		return [2]int{position[1], position[2] + 3}, nil
	case 2:
		return [2]int{position[1] + 3, position[2]}, nil
	case 3:
		return [2]int{position[1] + 3, position[2] + 3}, nil
	}

	// If other case, throw an error
	return [2]int{position[1], position[2]}, errors.New("Quadrant doesn't exist")
}

func Transpose(slice [3][3]string) [3][3]string {
    
    result := slice
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            result[i][j] = slice[j][i]
        }
    }

    return result
}

func Reverse(quadrant [3][3]string) [3][3]string {
	for x, row := range(quadrant) {
		quadrant[x][0] = row[2]
		quadrant[x][2] = row[0]
	}

	return quadrant
}

func Rotate(quadrant [3][3]string, direction string) [3][3] string {
	if direction == ROTATE_CLOCKWISE {
		return Reverse(Transpose(quadrant))
	}
	return Transpose(Reverse(quadrant))
}
