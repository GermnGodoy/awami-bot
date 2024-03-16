package possiblemoves

//La funcion ischeck determina si el rey esta en jaque tras un movimiento.
//Se asume que el input es un movimineto legal, y no se verifica.

func ischeck(movimiento move, board [64]int8) int8 {
	var wkingpos int8 = getwkingpos(board)
	makemove(movimiento, board)

	var j int8 = wkingpos - 9
	for j >= 9 && j%8 != 0 && board[j] == 0 {
		if board[j-9] == -4 || board[j-9] == -9 {
			return 0
		}
		j = j - 9
	}

	j = wkingpos + 9
	for j < 55 && j%8 != 7 && board[j] == 0 {
		if board[j+9] == -4 || board[j+9] == -9 {
			return 0
		}
		j = j + 9
	}

	j = wkingpos - 7
	for j > 7 && j%8 != 7 && board[j] == 0 {
		if board[j-9] == -4 || board[j-9] == -9 {
			return 0
		}
		j = j - 7
	}

	j = wkingpos + 7
	for j < 56 && j%8 != 0 && board[j] == 0 {
		if board[j+9] == -4 || board[j+9] == -9 {
			return 0
		}
		j = j + 7
	}

	j = wkingpos + 1
	for j%8 != 7 && board[j] == 0 {
		if board[j+1] == -5 || board[j+1] == -9 {
			return 0
		}
		j = j + 1
	}

	j = wkingpos - 1
	for j%8 != 0 && board[j] == 0 {
		if board[j-1] == -5 || board[j-1] == -9 {
			return 0
		}
		j = j - 1
	}

	j = wkingpos + 8
	for j < 56 && board[j] == 0 {
		if board[j+8] == -5 || board[j+8] == -9 {
			return 0
		}
		j = j + 8
	}

	j = wkingpos - 8
	for j > 7 && board[j] == 0 {
		if board[j-8] == -5 || board[j-8] == -9 {
			return 0
		}
		j = j - 8
	}
	if wkingpos > 7 {
		if wkingpos%8 != 7 && board[wkingpos-7] == -1 {

			return 0
		}
		if wkingpos%8 != 0 && board[wkingpos-9] == -1 {

			return 0
		}
	}
	if wkingpos > 15 {
		if board[wkingpos-15] == -3 && wkingpos%8 != 7 {

			return 0
		}
		if board[wkingpos-17] == -3 && wkingpos%8 != 0 {

			return 0
		}
	}
	if wkingpos > 7 {
		if board[wkingpos-6] == -3 && wkingpos%8 < 6 {

			return 0
		}
		if board[wkingpos-10] == -3 && wkingpos%8 > 1 {

			return 0
		}
	}
	if wkingpos < 48 {
		if board[wkingpos+15] == -3 && wkingpos%8 != 0 {

			return 0
		}
		if board[wkingpos+17] == -3 && wkingpos%8 != 7 {

			return 0
		}
	}
	if wkingpos < 56 {
		if board[wkingpos+6] == -3 && wkingpos%8 > 1 {

			return 0
		}
		if board[wkingpos+10] == -3 && wkingpos%8 < 6 {

			return 0
		}
	}

	return 1
}

func getwkingpos(board [64]int8) int8 {
	var i int8 = 0
	for i < 64 {
		if board[i] == int8(20) {
			break
		}
		i++
	}
	return i
}

func makemove(movement move, board [64]int8) {

	switch movement.finish {
	case -1:
		// En passant a la izquierda.
		board[movement.start-9], board[movement.start-1], board[movement.start] = 1, 0, 0
	case -2:
		// En passant a la derecha.
		board[movement.start-7], board[movement.start+1], board[movement.start] = 1, 0, 0
	case -3:
		// Enroque corto.
		board[60], board[61], board[62], board[63] = 0, 5, 20, 0
	case -4:
		//Enroque largo.
		board[60], board[58], board[57], board[56] = 0, 5, 20, 0
	default:
		board[movement.finish] = board[movement.start]
		board[movement.start] = 0
	}
}
