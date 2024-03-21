package possiblemoves

func castling(board [64]int8, enroques [4]uint8, listofmovesptr **node) {
	// Enroque corto
	if (enroques[0] == 1 && board[61] == 0 && board[62] == 0 && ischeck(move{60, 60}, board) == 1 &&
		ischeck(move{60, 61}, board) == 1 && ischeck(move{60, 62}, board) == 1) {
		push(move{60, -3}, listofmovesptr)
	}

	// Enroque largo
	if (enroques[1] == 1 && board[59] == 0 && board[58] == 0 && board[57] == 0 && ischeck(move{60, 60}, board) == 1 &&
		ischeck(move{60, 59}, board) == 1 && ischeck(move{60, 58}, board) == 1 && ischeck(move{60, 57}, board) == 1) {
		push(move{60, -4}, listofmovesptr)
	}
}

func kingmoves(i int8, board [64]int8, listofmovesptr **node) {
	var movlegal move
	movlegal.start = i
	if (i+1)%8 != 0 && board[i+1] <= 0 {
		movlegal.finish = i + 1
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if (i-1)%8 != 7 && board[i-1] <= 0 {
		movlegal.finish = i - 1
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i-8 >= 0 && board[i-8] <= 0 {
		movlegal.finish = i - 8
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i+8 < 64 && board[i+8] <= 0 {
		movlegal.finish = i + 8
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i+8 < 64 && (i-1)%8 != 0 && board[i+7] <= 0 {
		movlegal.finish = i + 7
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i-8 >= 0 && (i-1)%8 != 0 && board[i-9] <= 0 {
		movlegal.finish = i - 9
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i+8 < 64 && (i+1)%8 != 7 && board[i+9] <= 0 {
		movlegal.finish = i + 9
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i-8 >= 0 && (i+1)%8 != 7 && board[i-7] <= 0 {
		movlegal.finish = i - 7
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

}

func pawnmoves(i int8, board [64]int8, enpassant int8, listofmovesptr **node) {
	var movlegal move
	movlegal.start = i
	if board[i-8] == 0 {
		movlegal.finish = i - 8
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}
	if i < 56 && i > 47 && board[i-16] == 0 && board[i-8] == 0 {
		movlegal.finish = i - 16
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	//La variable enpassant tiene la posicion del peon que puede ser capturad por enpassant (en caso de que sea posible).
	if enpassant != 0 && board[enpassant-8] == 0 {
		if enpassant%8 != 0 && i+1 == enpassant {
			movlegal.finish = -1
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if enpassant%8 != 7 && i-1 == enpassant {
			movlegal.finish = -2
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}

		if i%8 != 7 && board[i-7] < 0 {
			movlegal.finish = i - 7
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if i%8 != 0 && board[i-9] < 0 {
			movlegal.finish = i - 7
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}

	}
}

func horsemoves(i int8, board [64]int8, listofmovesptr **node) {
	var movlegal move
	movlegal.start = i
	if i > 15 {
		if board[i-15] <= 0 && i%8 != 7 {
			movlegal.finish = i - 15
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if board[i-17] <= 0 && i%8 != 0 {
			movlegal.finish = i - 17
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
	}
	if i > 7 {
		if board[i-6] <= 0 && i%8 < 6 {
			movlegal.finish = i - 6
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if board[i-10] <= 0 && i%8 > 1 {
			movlegal.finish = i - 10
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
	}
	if i < 48 {
		if board[i+15] <= 0 && i%8 != 0 {
			movlegal.finish = i + 15
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if board[i+17] <= 0 && i%8 != 7 {
			movlegal.finish = i + 17
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
	}
	if i < 56 {
		if board[i+6] <= 0 && i%8 > 1 {
			movlegal.finish = i + 6
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
		if board[i+10] <= 0 && i%8 < 6 {
			movlegal.finish = i + 10
			if ischeck(movlegal, board) == 1 {
				push(movlegal, listofmovesptr)
			}
		}
	}

}

func bishopmoves(i int8, board [64]int8, listofmovesptr **node) {
	var movlegal move
	movlegal.start = i
	var j int8 = i + 9

	for j < 64 && (j-9)%8 != 7 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j + 9
	}
	if j < 64 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i - 9
	for j >= 0 && (j+9)%8 != 0 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j - 9
	}
	if j >= 0 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i + 7
	for j < 64 && (j-7)%8 != 0 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j + 7
	}
	if j < 64 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i - 7
	for j >= 0 && (j+7)%8 != 7 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j - 7
	}
	if j >= 0 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

}
func rookmoves(i int8, board [64]int8, listofmovesptr **node) {
	var movlegal move
	movlegal.start = i
	var j int8 = i + 1

	for j%8 != 0 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j++
	}
	if j%8 != 0 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i - 1
	for j%8 != 7 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j - 1
	}
	if j%8 != 7 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i + 8
	for j < 64 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j + 8
	}
	if j < 64 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

	j = i - 8
	for j >= 0 && board[j] == 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
		j = j - 8
	}
	if j >= 0 && board[j] < 0 {
		movlegal.finish = j
		if ischeck(movlegal, board) == 1 {
			push(movlegal, listofmovesptr)
		}
	}

}
