package possiblemoves

import "errors"

//La funcion ischeck determina si el rey esta en jaque tras un movimiento.
//Se asume que el input es un movimineto legal, y no se verifica.

func ischeck(movimiento Move, board [64]int8) int8 {
	newboard := makeMove(movimiento, board)
	wkingpos, err := getwkingpos(newboard)
	if err != nil {
		return -1
	}

	var j int8 = wkingpos - 9
	if j >= 0 && (newboard[j] == -4 || newboard[j] == -9) {
		return 0
	}
	for j >= 9 && j%8 != 0 && newboard[j] == 0 {
		if newboard[j-9] == -4 || newboard[j-9] == -9 {
			return 0
		}
		j = j - 9
	}

	j = wkingpos + 9
	if j < 64 && (newboard[j] == -4 || newboard[j] == -9) {
		return 0
	}
	for j < 55 && j%8 != 7 && newboard[j] == 0 {
		if newboard[j+9] == -4 || newboard[j+9] == -9 {
			return 0
		}
		j = j + 9
	}

	j = wkingpos - 7
	if j >= 0 && (newboard[j] == -4 || newboard[j] == -9) {
		return 0
	}
	for j > 7 && j%8 != 7 && newboard[j] == 0 {
		if newboard[j-9] == -4 || newboard[j-9] == -9 {
			return 0
		}
		j = j - 7
	}

	j = wkingpos + 7
	if j < 64 && (newboard[j] == -4 || newboard[j] == -9) {
		return 0
	}
	for j < 56 && j%8 != 0 && newboard[j] == 0 {
		if newboard[j+9] == -4 || newboard[j+9] == -9 {
			return 0
		}
		j = j + 7
	}

	j = wkingpos + 1
	if j < 64 && (newboard[j] == -5 || newboard[j] == -9) {
		return 0
	}
	for j%8 != 7 && newboard[j] == 0 {
		if newboard[j+1] == -5 || newboard[j+1] == -9 {
			return 0
		}
		j = j + 1
	}

	j = wkingpos - 1
	if j >= 0 && (newboard[j] == -5 || newboard[j] == -9) {
		return 0
	}
	for j%8 != 0 && newboard[j] == 0 {
		if newboard[j-1] == -5 || newboard[j-1] == -9 {
			return 0
		}
		j = j - 1
	}

	j = wkingpos + 8
	if j < 64 && (newboard[j] == -5 || newboard[j] == -9) {
		return 0
	}
	for j < 56 && newboard[j] == 0 {
		if newboard[j+8] == -5 || newboard[j+8] == -9 {
			return 0
		}
		j = j + 8
	}

	j = wkingpos - 8
	if j >= 0 && (newboard[j] == -5 || newboard[j] == -9) {
		return 0
	}
	for j > 7 && newboard[j] == 0 {
		if newboard[j-8] == -5 || newboard[j-8] == -9 {
			return 0
		}
		j = j - 8
	}
	if wkingpos > 7 {
		if wkingpos%8 != 7 && newboard[wkingpos-7] == -1 {

			return 0
		}
		if wkingpos%8 != 0 && newboard[wkingpos-9] == -1 {

			return 0
		}
	}
	if wkingpos > 15 {
		if wkingpos%8 != 7 && newboard[wkingpos-15] == -3 {

			return 0
		}
		if wkingpos%8 != 0 && newboard[wkingpos-17] == -3 {

			return 0
		}
	}
	if wkingpos > 7 {
		if wkingpos%8 < 6 && newboard[wkingpos-6] == -3 {

			return 0
		}
		if wkingpos%8 > 1 && newboard[wkingpos-10] == -3 {

			return 0
		}
	}
	if wkingpos < 48 {
		if wkingpos%8 != 0 && newboard[wkingpos+15] == -3 {

			return 0
		}
		if wkingpos%8 != 7 && newboard[wkingpos+17] == -3 {

			return 0
		}
	}

	if wkingpos < 56 {
		if wkingpos%8 > 1 && newboard[wkingpos+6] == -3 {

			return 0
		}
		if wkingpos%8 < 6 && newboard[wkingpos+10] == -3 {

			return 0
		}
	}

	return 1
}

func getwkingpos(newboard [64]int8) (int8, error) {
	var i int8 = 0
	for i < 64 {
		if newboard[i] == int8(20) {
			return i, nil
		}
		i++
	}
	return -1, errors.New("there is no white king")
}

func makeMove(Movement Move, board [64]int8) [64]int8 {

	switch Movement.Finish {
	case -1:
		// En passant a la izquierda.
		board[Movement.Start-9], board[Movement.Start-1], board[Movement.Start] = 1, 0, 0
	case -2:
		// En passant a la derecha.
		board[Movement.Start-7], board[Movement.Start+1], board[Movement.Start] = 1, 0, 0
	case -3:
		// Enroque corto.
		board[60], board[61], board[62], board[63] = 0, 5, 20, 0
	case -4:
		//Enroque largo.
		board[60], board[58], board[57], board[56] = 0, 5, 20, 0
	default:
		tmp := board[Movement.Start]
		board[Movement.Start], board[Movement.Finish] = 0, tmp
	}

	return board
}
