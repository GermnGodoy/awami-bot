package main

import (
	"errors"
	"fmt"
	"strconv"
)

var InitialPos = "rhbkqbhrpppppppp88888PPPPPPPP-0-0-"

var Pieces = map[string]int8{
	// Map with all the CR7 equivalences, lowercase refers to black pieces, uppercase to white ones
	"P": 1,
	"H": 2, // Used "H"orse instead of "K"night for not having duplicated keys
	"B": 3,
	"R": 4,
	"Q": 5,
	"K": 6,
	"p": -1,
	"h": -2, // Used "h"orse instead of "k"night for not having duplicated keys
	"b": -3,
	"r": -4,
	"q": -5,
	"k": -6,
}

var invPieces = invMap(Pieces)

var SeedLength = 4

func invMap(m map[string]int8) map[int8]string {
	var reversedMap map[int8]string

	for key, value := range m {
		reversedMap[value] = key
	}

	return reversedMap
}

func CR7ToArray(cr string) ([64]int8, [2]byte, string, error) {

	// TODO: Añadir tabla de contexto

	var board [64]int8
	var castling [2]byte

	offset := 0

	for index, rune := range cr {

		if index >= len(cr)-SeedLength {
			break
		}

		char := string(rune)

		if num, err := strconv.Atoi(char); err == nil {
			if index >= len(cr)-1+SeedLength {
				castling[len(cr)-index+SeedLength] = byte(num)
			}
			offset += num - 1
			continue
		}

		if Pieces[char] == 0 {
			return [64]int8{}, [2]byte{}, "", errors.New("Invalid CR7 code")
		}

		board[index] = Pieces[char]
	}

	seed := cr[len(cr)-SeedLength:]

	return board, castling, seed, nil
}

func ArrayToCR7(board [64]int8, castling [2]byte, seed string) (string, error) {

	CR7string := ""

	offset := 0
	for _, value := range board {

		if value == 0 {
			offset++
			continue
		}

		if invPieces[value] == "" {
			return "", errors.New("Invalid Board")
		}

		if offset != 0 {
			CR7string += fmt.Sprint(offset)
		}

		CR7string += invPieces[value]
	}

	for _, castle := range castling {

		if string(castle) != "0" && string(castle) != "1" {
			return "", errors.New("Bad castling table")
		}

		CR7string += string(castle)
	}

	CR7string += seed

	return CR7string, nil
}
