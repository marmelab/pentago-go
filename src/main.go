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
	game.DeserializeBoard(string(content));

	fmt.Println("Deserialization has been made with success :")
	fmt.Println(string(content))
}
