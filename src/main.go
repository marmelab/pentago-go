package main

import (
    "log"
	"fmt"
	"strings"
	"strconv"
	"sync"
	"sort"
	"encoding/json"
	"net/http"
	"game"
	"ai"
	"constants"
)



type Result struct {
	move ai.Move
	opponentMove ai.Move
	Score int
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

func StartFirstNodeMinimax(board game.Board, currentPlayer string, move ai.Move, ch chan Result, key int, wg *sync.WaitGroup) {
	
	defer wg.Done()
	newBoard := ai.ApplyMoveOnBoard(board, move, currentPlayer)
	opponent := ai.SwitchPlayer(currentPlayer)

	score, opponentMove := ai.Minimax(
		constants.DEPTH - 1,
		newBoard,
		currentPlayer,
		opponent,
		move,
		-constants.SCORE_ALIGNED[4],
		constants.SCORE_ALIGNED[4],
	)
	ch <- Result{
		move: move,
		opponentMove: opponentMove,
		Score: score,
	}
}

type ResultData struct {
	Score int
	PlaceMarble [2]int
	Rotate [2]string
}

func getBestMoveForPlayer(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	
	arr := game.TwoDimensionArrayBoard{}
	json.NewDecoder(r.Body).Decode(&arr)
	currentPlayer := strconv.Itoa(arr.CurrentPlayer)

	board := game.DeserializeTwoDimensionArrayBoard(arr)
	// Make communication between main & routines possible
	ch := make(chan Result)
	
	var results []Result

	// The WaitGroup will wait len(moves) routines
	moves := ai.GetAllPossibleMoves(board);


	wg.Add(len(moves))
	for x, move := range moves {
		go StartFirstNodeMinimax(board, currentPlayer, move, ch, x, &wg)
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
	
	sort.Slice(
		results,
		func(i, j int) bool { return results[i].Score > results[j].Score },
	)
	result := results[0]
	fmt.Println(result)
	placeMarble, _ := game.ConvertQuadrantPositionIntoBoardPosition(result.move.PlaceMarble)

	data := ResultData{
		Score: result.Score,
		PlaceMarble: placeMarble,
		Rotate: result.move.Rotate,
	}
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", getBestMoveForPlayer)

    log.Println("Listening on :8083")

    log.Fatal(http.ListenAndServe(":8083", nil))

}
