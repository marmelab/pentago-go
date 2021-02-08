package main

import (
	"io/ioutil"
    "log"
	"fmt"
	game "game"
)

func main() {

    content, err := ioutil.ReadFile("./src/datasets/empty_board.txt")

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

	if result == game.GAME_PLAYER1_WON {
		fmt.Println("Player 1 win !")
	} else if result == game.GAME_PLAYER2_WON {
		fmt.Println("Player 2 win !")
	} else if result == game.GAME_DRAW{
		fmt.Println("It's a draw !")
	} else {
		fmt.Println("Game is running...")
	}
}
