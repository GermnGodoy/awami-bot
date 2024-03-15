package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

var InitialPos = "rhbqkbhrpppppppp8888PPPPPPPPRHBQKBHR000000seed"

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

func invMap(m map[string]int8) map[int8]string {
	var reversedMap = make(map[int8]string)

	for key, value := range m {
		reversedMap[value] = key
	}

	return reversedMap
}

var invPieces = invMap(Pieces)

var SeedLength = 4

func CR7Decoder(cr string) ([64]int8, int8, [4]byte, string, error) {

	// TODO: Añadir tabla de contexto

	seed := cr[utf8.RuneCountInString(cr)-SeedLength:]

	var board [64]int8
	var castling [4]byte
	var pasant int8

	var offset int

	for index, rune := range cr {

		if index+offset == 64 {
			if num, err := strconv.Atoi(cr[index : index+2]); err == nil {
				pasant = int8(num)
				continue
			}
			return board, 0, [4]byte{}, cr[index : index+2], errors.New("invalid CR7 code")
		}

		if index+offset == 65 {
			continue
		}

		if index >= utf8.RuneCountInString(cr)-SeedLength {
			break
		}

		char := string(rune)

		if num, err := strconv.Atoi(char); err == nil {
			if index >= utf8.RuneCountInString(cr)-1+SeedLength {
				castling[utf8.RuneCountInString(cr)-index+SeedLength] = byte(num)
			}
			offset += num - 1
			continue
		}

		if Pieces[char] == 0 {
			return [64]int8{}, 0, [4]byte{}, char, errors.New("invalid CR7 code")
		}

		board[index+offset] = Pieces[char]
	}

	return board, pasant, castling, seed, nil
}

func CR7Encoder(board [64]int8, pasant int8, castling [4]byte, seed string) (string, error) {

	CR7string := ""

	offset := 0
	for _, value := range board {

		if value == 0 {
			if offset == 8 {
				CR7string += fmt.Sprint(offset)
				offset = 0
			}
			offset++
			continue
		}

		if invPieces[value] == "" {
			return fmt.Sprint(value), errors.New("invalid Board")
		}

		if offset != 0 {
			CR7string += fmt.Sprint(offset)
			offset = 0
		}

		CR7string += invPieces[value]
	}

	CR7string += fmt.Sprintf("%02d", pasant)

	for _, castle := range castling {

		if castle != 0 && castle != 1 {
			return "", errors.New("bad castling table")
		}

		CR7string += fmt.Sprint(castle)
	}

	CR7string += seed

	return CR7string, nil
}
