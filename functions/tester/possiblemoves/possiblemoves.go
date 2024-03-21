package possiblemoves

import "fmt"

type move struct {
	start  int8
	finish int8
}

type node struct {
	move move
	next *node
}

//pointer to the linked list

func Printpossiblemoves(board [64]int8, enpassant int8, possiblecastles [4]uint8) {
	var listofmoves *node = nil

	//chequeamos primero posibles ernoques.
	castling(board, possiblecastles, &listofmoves)

	var i int8 = 0
	for i < 64 {
		//printf("%i\n", board[i]);
		for board[i] == 0 {
			i++
		}
		switch board[i] {
		//Caso peon blanco
		case 1:
			pawnmoves(i, board, enpassant, &listofmoves)
		// Caso caballo blanco.
		case 3:
			horsemoves(i, board, &listofmoves)
		// Caso alfil.
		case 4:
			bishopmoves(i, board, &listofmoves)
		// Caso torre.
		case 5:
			rookmoves(i, board, &listofmoves)
			//El enroque se hace a aparte (al principio).
		// Caso reina.
		case 9:
			bishopmoves(i, board, &listofmoves)
			rookmoves(i, board, &listofmoves)
		// Caso rey.
		case 20:
			kingmoves(i, board, &listofmoves)
			//El enroque se hace a aparte (al principio).
		default:
		}
		i++
	}

	// ahora imprimimos los mivimientos obtenidos
	ptr := listofmoves

	if ptr == nil {
		fmt.Printf("No legal moves.")
		return
	}

	for ptr.next != nil {
		fmt.Printf("[%d, %d]\n", ptr.move.start, ptr.move.finish)
		ptr = ptr.next
	}
	fmt.Printf("[%d, %d]\n", ptr.move.start, ptr.move.finish)
}

func push(movimiento move, listofmovesptr **node) {
	newelement := &node{
		move: movimiento,
		next: *listofmovesptr,
	}
	*listofmovesptr = newelement
}
