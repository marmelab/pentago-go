package main

import (
    "log"
	"fmt"
	game "game"
	fileReader "fileReader"
	ai "ai"
	"time"
)

func main() {

	content, err := fileReader.GetFileContent()

	if err != nil {
		log.Fatal(err)
	}

	board, err := game.DeserializeBoard(string(content))
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	results := ai.PlayAllPossibleMoves(board)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You win if you :")
	for _, result := range(results.Win) {
		fmt.Printf(
			"Place a marble in %d %d and rotate quadrant %v in %v \n",
			result.PlaceMarble[0],
			result.PlaceMarble[1],
			result.Rotate[0],
			result.Rotate[1],
		)
	}

	fmt.Println("\nYou loose if you :")
	for _, result := range(results.Loose) {
		fmt.Printf(
			"Place a marble in %d %d and rotate quadrant %v in %v \n",
			result.PlaceMarble[0],
			result.PlaceMarble[1],
			result.Rotate[0],
			result.Rotate[1],
		)
	}

	elapsed := time.Since(start)

	fmt.Printf("Found in %v\n\n", elapsed)
}
