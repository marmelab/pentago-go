package game

import (
	"strings"
	"strconv"
	"regexp"
	"errors"
)

func DeserializeBoard(str string) ([][]int, error) {
	var board [][]int
	// Split string into array based on \n
	strArr := strings.Fields(str)
	// Remove first line, last line & middle line which are board's delimiters.
	if len(strArr) != 9 {
		return nil, errors.New("Given string is not correctly formatted and have to contain 9 lines.")
	}
	

	arrayWithoutBorders := append(make([]string, 0, 6), strArr[1:4]...)
	arrayWithoutBorders = append(arrayWithoutBorders, strArr[5:8]...)

	re := regexp.MustCompile("[0-9]+")

	for _, line := range arrayWithoutBorders {
		// Remove all non-digits chars
		lineStrings := re.FindAllString(line, -1)

		if len(lineStrings) != 6 {
			return nil, errors.New("All line should contains 6 values.")
		}

		// This loop allow us to cast all digits into integers
		var lineInt = make([]int, 0, len(lineStrings))

		for _, value := range lineStrings {
			value, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			lineInt = append(lineInt, value)
		}
		board = append(board, lineInt)
	}

	// Return our 2-d arrays of integers
	return board, nil
}
