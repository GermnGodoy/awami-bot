package possiblemoves

func castling(board [64]int8, PossibleCastles [4]uint8) []Move {
	// Enroque corto
	var LegalCastles []Move
	if (PossibleCastles[0] == 1 && board[61] == 0 && board[62] == 0 && ischeck(Move{60, 60}, board) == 1 &&
		ischeck(Move{60, 61}, board) == 1 && ischeck(Move{60, 62}, board) == 1) {
		LegalCastles = append(LegalCastles, Move{60, -3})
	}

	// Enroque largo
	if (PossibleCastles[1] == 1 && board[59] == 0 && board[58] == 0 && board[57] == 0 && ischeck(Move{60, 60}, board) == 1 &&
		ischeck(Move{60, 59}, board) == 1 && ischeck(Move{60, 58}, board) == 1 && ischeck(Move{60, 57}, board) == 1) {
		LegalCastles = append(LegalCastles, Move{60, -4})
	}
	return LegalCastles
}

func kingMoves(square int8, board [64]int8) []Move {
	var movlegal Move
	var LegalMoves []Move
	movlegal.Start = square
	if (square)%8 != 7 && board[square+1] <= 0 {
		movlegal.Finish = square + 1
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if (square)%8 != 0 && board[square-1] <= 0 {
		movlegal.Finish = square - 1
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square-8 >= 0 && board[square-8] <= 0 {
		movlegal.Finish = square - 8
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square+8 < 64 && board[square+8] <= 0 {
		movlegal.Finish = square + 8
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square+8 < 64 && (square-1)%8 != 0 && board[square+7] <= 0 {
		movlegal.Finish = square + 7
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square-8 >= 0 && (square-1)%8 != 0 && board[square-9] <= 0 {
		movlegal.Finish = square - 9
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square+8 < 64 && (square+1)%8 != 7 && board[square+9] <= 0 {
		movlegal.Finish = square + 9
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square-8 >= 0 && (square+1)%8 != 7 && board[square-7] <= 0 {
		movlegal.Finish = square - 7
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	return LegalMoves
}

func pawnMoves(square int8, board [64]int8, enpassant int8) []Move {
	var movlegal Move
	var LegalMoves []Move
	movlegal.Start = square
	if board[square-8] == 0 {
		movlegal.Finish = square - 8
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	if square < 56 && square > 47 && board[square-16] == 0 && board[square-8] == 0 {
		movlegal.Finish = square - 16
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	//La variable enpassant tiene la posicion del peon que puede ser capturad por enpassant (en caso de que sea posible).
	if enpassant != 0 && board[enpassant-8] == 0 {
		if enpassant%8 != 0 && square+1 == enpassant {
			movlegal.Finish = -1
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if enpassant%8 != 7 && square-1 == enpassant {
			movlegal.Finish = -2
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}

		if square%8 != 7 && board[square-7] < 0 {
			movlegal.Finish = square - 7
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if square%8 != 0 && board[square-9] < 0 {
			movlegal.Finish = square - 7
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}

	}
	return LegalMoves
}

func horseMoves(square int8, board [64]int8) []Move {
	var movlegal Move
	var LegalMoves []Move
	movlegal.Start = square
	if square > 15 {
		if square%8 != 7 && board[square-15] <= 0 {
			movlegal.Finish = square - 15
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if square%8 != 0 && board[square-17] <= 0 {
			movlegal.Finish = square - 17
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
	}
	if square > 7 {
		if square%8 < 6 && board[square-6] <= 0 {
			movlegal.Finish = square - 6
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if square%8 > 1 && board[square-10] <= 0 {
			movlegal.Finish = square - 10
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
	}
	if square < 48 {
		if square%8 != 0 && board[square+15] <= 0 {
			movlegal.Finish = square + 15
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if square%8 != 7 && board[square+17] <= 0 {
			movlegal.Finish = square + 17
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
	}
	if square < 56 {
		if square%8 > 1 && board[square+6] <= 0 {
			movlegal.Finish = square + 6
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
		if square%8 < 6 && board[square+10] <= 0 {
			movlegal.Finish = square + 10
			if ischeck(movlegal, board) == 1 {
				LegalMoves = append(LegalMoves, movlegal)
			}
		}
	}
	return LegalMoves
}

func bishopMoves(square int8, board [64]int8) []Move {
	var movlegal Move
	var LegalMoves []Move
	movlegal.Start = square
	var j int8 = square + 9

	for j < 64 && (j-9)%8 != 7 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j + 9
	}
	if j < 64 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square - 9
	for j >= 0 && (j+9)%8 != 0 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j - 9
	}
	if j >= 0 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square + 7
	for j < 64 && (j-7)%8 != 0 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j + 7
	}
	if j < 64 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square - 7
	for j >= 0 && (j+7)%8 != 7 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j - 7
	}
	if j >= 0 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	return LegalMoves
}
func rookMoves(square int8, board [64]int8) []Move {
	var movlegal Move
	var LegalMoves []Move
	movlegal.Start = square
	var j int8 = square + 1

	for j%8 != 0 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j++
	}
	if j%8 != 0 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square - 1
	for j%8 != 7 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j - 1
	}
	if j%8 != 7 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square + 8
	for j < 64 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j + 8
	}
	if j < 64 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}

	j = square - 8
	for j >= 0 && board[j] == 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
		j = j - 8
	}
	if j >= 0 && board[j] < 0 {
		movlegal.Finish = j
		if ischeck(movlegal, board) == 1 {
			LegalMoves = append(LegalMoves, movlegal)
		}
	}
	return LegalMoves
}
