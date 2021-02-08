package main

import (
	"io/ioutil"
    "log"
	"os"
	"fmt"
	game "game"
)

func main() {
	fileName := os.Args[1]

	fmt.Println(fileName)
    
	content, err := ioutil.ReadFile("./src/datasets/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	board, err := game.DeserializeBoard(string(content))
	if err != nil {
		log.Fatal(err)
	}

	result, err := game.DetectWinner(board)
	if err != nil {
		log.Fatal(err)
	}

	switch result {
	case game.GAME_PLAYER1_WON:
		fmt.Println("Player 1 win !")
	case game.GAME_PLAYER2_WON:
		fmt.Println("Player 2 win !")
	case game.GAME_DRAW:
		fmt.Println("It's a draw")
	default:
		fmt.Println("Game is running...")
	}
}
