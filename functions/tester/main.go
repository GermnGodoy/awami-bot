package main

import (
	"fmt"
	"tester/test/possiblemoves"
)

var board = [64]int8{-5, -3, -4, -9, -20, -4, -3, -5,
	-1, -1, -1, -1, -1, 0, -1, -1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, -1, 0, 0,
	0, 3, 0, 1, 0, 0, 5, 0,
	4, 0, 0, 0, -5, 0, 0, 0,
	1, 1, 1, 0, 9, 1, 1, 1,
	5, 0, 0, 0, 20, 0, 0, 5}

var enpassant int8 = 29
var possiblecastles = [4]uint8{1, 1, 0, 0}

func main() {
	var Moves []possiblemoves.Move = possiblemoves.SliceOfPossibleMoves(board, enpassant, possiblecastles)
	i := 0
	for i < len(Moves) {
		fmt.Printf("[%d, %d]\n", Moves[i].Start, Moves[i].Finish)
		i++
	}
}
