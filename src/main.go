package main

import (
    "log"
	"fmt"
	game "game"
	fileReader "fileReader"
	ai "ai"
	"time"
	"strings"
	// "sort"
)

const MAX_RESULTS = 10
const DEPTH = 2

func PrintBoard(board string) {
	fmt.Println("   0 1 2  3 4 5")
	boardSplitted := strings.Split(board, "\n")
	for x, line := range boardSplitted {
		if x == 0 || x == 4 || x > 7 {
			fmt.Println("  " + line)
		} else {
			var lineNumber int
			if x > 0 && x < 4 {
				lineNumber = x - 1
			} else if x > 4 {
				lineNumber = x - 2
			}
	
			fmt.Println(lineNumber, line)
		}

	}

}

func main() {
	content, err := fileReader.GetFileContent()

	if err != nil {
		log.Fatal(err)
	}

	board, err := game.DeserializeBoard(string(content))
	if err != nil {
		log.Fatal(err)
	}

	PrintBoard(content)

	start := time.Now()

	var startMove ai.Move
	score, move := ai.Minimax(DEPTH, board, "1", startMove)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(score, move)

	elapsed := time.Since(start)

	fmt.Printf("\nFound in %v\n\n", elapsed)
}
