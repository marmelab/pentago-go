package game

import (
	"strings"
	"strconv"
	"regexp"
)

func DeserializeBoard(str string) [][]int {
	var board [][]int
	// Split string into array based on \n
	strArr := strings.Fields(str)
	// Remove first line, last line & middle line which are board's delimiters.
	strArr = append(strArr[1:4], strArr[5:8]...)

	re := regexp.MustCompile("[0-9]+")

	for _, line := range strArr {
		// Remove all non-digits chars
		lineStrings := re.FindAllString(line, -1)

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
	return board
}
