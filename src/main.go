package main

import (
    "log"
	"fmt"
	game "game"
	fileReader "fileReader"
	ai "ai"
	"time"
	"math/bits"
)

func main() {
	number1 := 0b10000100000100000100000100000000000
	number2 := 0b10000110000100000100000100000000000

	numberResult := bits.OnesCount64(uint64(number1 & number2))
	fmt.Println("OK", numberResult)

	content, err := fileReader.GetFileContent()

	if err != nil {
		log.Fatal(err)
	}

	board, err := game.DeserializeBoard(string(content))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

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
