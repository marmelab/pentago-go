package main

import (
    "log"
	"fmt"
	"time"
	"strings"
	"sync"
	"sort"
	"game"
	"fileReader"
	"ai"
	"constants"
)



type Result struct {
	move ai.Move
	score int
}

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

func StartFirstNodeMinimax(board game.Board, move ai.Move, ch chan Result, key int, wg *sync.WaitGroup) {
	
	defer wg.Done()

	newBoard := ai.ApplyMoveOnBoard(board, move, constants.PLAYER1_VALUE)
    // fmt.Printf("Worker %d starting\n", key)
	score, _ := ai.Minimax(
		constants.DEPTH - 1,
		newBoard,
		constants.PLAYER2_VALUE,
		move,
		-constants.SCORE_ALIGNED[4],
		constants.SCORE_ALIGNED[4],
	)
	// fmt.Println("finished", key)
	ch <- Result{
		move: move,
		score: score,
	}
}

func EstimateBrowsedNodes(rootTreeLength int, iteration int) int {
	result := rootTreeLength

	for i := 1; i < iteration; i++ {
		result = result * (rootTreeLength - i)
	}

	return result
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

	var wg sync.WaitGroup
	start := time.Now()

	moves := ai.GetAllPossibleMoves(board);
	
	var results []Result
	ch := make(chan Result)

	for x, move := range moves {
		wg.Add(1)
		go StartFirstNodeMinimax(board, move, ch, x, &wg)
	}
	
	go func() {
		for result := range ch {
			results = append(results, result)
        }
	}()
		
	wg.Wait()
	elapsed := time.Since(start)


	numberOfNodesToCompute := EstimateBrowsedNodes(len(results), constants.DEPTH)
	fmt.Printf("Results analyzed : %d, Depth : %d\n", numberOfNodesToCompute, constants.DEPTH)
	sort.Slice(
		results,
		func(i, j int) bool { return results[i].score > results[j].score },
	)
	fmt.Println("Suggested moves :")
	for _, result := range(results[:constants.MAX_RESULTS]) {
		placeMarble, _ := game.ConvertQuadrantPositionIntoBoardPosition(result.move.PlaceMarble)
		rotate := result.move.Rotate

		fmt.Printf(
			"=> %d : Place a marble in %d %d and rotate quadrant %v in %v \n",
			result.score,
			placeMarble[0],
			placeMarble[1],
			rotate[0],
			rotate[1],
		)
	}

	close(ch)


	fmt.Printf("\nFound in %v\n\n", elapsed)
}
