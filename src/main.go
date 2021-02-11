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
	opponentMove ai.Move
	score int
}

func PrintBoardFromFile(board string) {
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

func PrintBoard(board game.Board) {
	boardStr := game.ToStringBoard(board)

	fmt.Println("┌────────────┐")
	for i, r := range boardStr {
		if i == 18 {
			fmt.Println("|────────────|")
		}
		fmt.Printf("|%c", r)
		if (i + 1) %6 == 0 {
			fmt.Printf("|\n")
		} else if (i+1) %3 == 0 {
			fmt.Printf("|")
		}
	}
	fmt.Println("└────────────┘")
}

func StartFirstNodeMinimax(board game.Board, move ai.Move, ch chan Result, key int, wg *sync.WaitGroup) {
	
	defer wg.Done()

	newBoard := ai.ApplyMoveOnBoard(board, move, constants.FIRST_PLAYER_VALUE)
	opponent := ai.SwitchPlayer(constants.FIRST_PLAYER_VALUE)
    // fmt.Printf("Worker %d starting\n", key)
	score, opponentMove := ai.Minimax(
		constants.DEPTH - 1,
		newBoard,
		opponent,
		move,
		-constants.SCORE_ALIGNED[4],
		constants.SCORE_ALIGNED[4],
	)
	ch <- Result{
		move: move,
		opponentMove: opponentMove,
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
	var wg sync.WaitGroup

	content, err := fileReader.GetFileContent()
	PrintBoardFromFile(content)
	
	if err != nil {
		log.Fatal(err)
	}

	board, err := game.DeserializeBoard(string(content))
	if err != nil {
		log.Fatal(err)
	}


	start := time.Now()

	moves := ai.GetAllPossibleMoves(board);
	
	var results []Result

	// Make communication between main & routines possible
	ch := make(chan Result)

	// The WaitGroup will wait len(moves) routines
	wg.Add(len(moves))

	for x, move := range moves {
		go StartFirstNodeMinimax(board, move, ch, x, &wg)
	}

	// Compute result sent in the channel
	go func() {
		for result := range ch {
			results = append(results, result)
        }
	}()

	// Wait all workers has been closed
	wg.Wait()
	close(ch)

	elapsed := time.Since(start)


	numberOfNodesToCompute := EstimateBrowsedNodes(len(results), constants.DEPTH)
	fmt.Printf("Results analyzed : %d, Depth : %d\n", numberOfNodesToCompute, constants.DEPTH)
	sort.Slice(
		results,
		func(i, j int) bool { return results[i].score > results[j].score },
	)
	fmt.Println("Suggested moves :")
	for _, result := range(results[:constants.MAX_RESULTS]) {


		fmt.Printf(
			"\n===> %d : ",
			result.score,
		)

		move := result.move

		placeMarble, _ := game.ConvertQuadrantPositionIntoBoardPosition(move.PlaceMarble)
		rotate := move.Rotate

		fmt.Printf("You should place a marble in %d %d and rotate quadrant %v in %v \n",
			placeMarble[0],
			placeMarble[1],
			rotate[0],
			rotate[1],
		)

		if constants.PRINT_BOARD_FOR_MOVES == true {
			newBoard := ai.ApplyMoveOnBoard(board, move, constants.FIRST_PLAYER_VALUE)
			PrintBoard(newBoard)
	
			move = result.opponentMove
			placeMarble, _ = game.ConvertQuadrantPositionIntoBoardPosition(move.PlaceMarble)
			rotate = move.Rotate
			
			fmt.Printf("\nOpponent should play in %d %d and rotate quadrant %v in %v to get the following result : \n",
				placeMarble[0],
				placeMarble[1],
				rotate[0],
				rotate[1],
			)
	
			opponent := ai.SwitchPlayer(constants.FIRST_PLAYER_VALUE)
			newBoard = ai.ApplyMoveOnBoard(newBoard, result.opponentMove, opponent)
	
			PrintBoard(newBoard)
		}

	}

	fmt.Printf("\nFound in %v\n\n", elapsed)
}
