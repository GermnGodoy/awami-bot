package possiblemoves

type Move struct {
	Start  int8
	Finish int8
}

//pointer to the linked list

func SliceOfPossibleMoves(board [64]int8, enpassant int8, possiblecastles [4]uint8) []Move {
	var LegalMoves []Move

	//chequeamos primero posibles ernoques.
	LegalMoves = append(LegalMoves, castling(board, possiblecastles)...)

	var square int8 = 0
	for square < 64 {
		//printf("%i\n", board[i]);
		for board[square] == 0 {
			square++
		}
		switch board[square] {
		//Caso peon blanco
		case 1:
			LegalMoves = append(LegalMoves, pawnMoves(square, board, enpassant)...)
		// Caso caballo blanco.
		case 3:
			LegalMoves = append(LegalMoves, horseMoves(square, board)...)
		// Caso alfil.
		case 4:
			LegalMoves = append(LegalMoves, bishopMoves(square, board)...)
		// Caso torre.
		case 5:
			LegalMoves = append(LegalMoves, rookMoves(square, board)...)
			//El enroque se hace a aparte (al principio).
		// Caso reina.
		case 9:
			LegalMoves = append(LegalMoves, bishopMoves(square, board)...)
			LegalMoves = append(LegalMoves, rookMoves(square, board)...)
		// Caso rey.
		case 20:
			LegalMoves = append(LegalMoves, kingMoves(square, board)...)
			//El enroque se hace a aparte (al principio).
		default:
		}
		square++
	}

	return LegalMoves
}
