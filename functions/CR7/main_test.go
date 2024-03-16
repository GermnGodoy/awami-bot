package main

import "testing"

type DecoderTest struct {
	CR7               string
	expected_board    [64]int8
	expected_pasant   int8
	expected_castling [4]byte
	expected_seed     string
}

type EncoderTest struct {
	board        [64]int8
	pasant       int8
	castling     [4]byte
	seed         string
	expected_CR7 string
}

var DecoderTests = []DecoderTest{
	{
		CR7: "rhbqkbhrpppppppp8888PPPPPPPPRHBQKBHR000000seed",
		expected_board: [64]int8{Pieces["r"], Pieces["h"], Pieces["b"], Pieces["q"], Pieces["k"], Pieces["b"], Pieces["h"], Pieces["r"],
			Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"],
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"],
			Pieces["R"], Pieces["H"], Pieces["B"], Pieces["Q"], Pieces["K"], Pieces["B"], Pieces["H"], Pieces["R"]},
		expected_pasant:   0,
		expected_castling: [4]byte{0, 0, 0, 0},
		expected_seed:     "seed",
	},
}

var EncoderTests = []EncoderTest{
	{
		board: [64]int8{Pieces["r"], Pieces["h"], Pieces["b"], Pieces["q"], Pieces["k"], Pieces["b"], Pieces["h"], Pieces["r"],
			Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"], Pieces["p"],
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"], Pieces["P"],
			Pieces["R"], Pieces["H"], Pieces["B"], Pieces["Q"], Pieces["K"], Pieces["B"], Pieces["H"], Pieces["R"]},
		pasant:       0,
		castling:     [4]byte{0, 0, 0, 0},
		seed:         "seed",
		expected_CR7: "rhbqkbhrpppppppp8888PPPPPPPPRHBQKBHR000000seed",
	},
}

func TestEncoder(t *testing.T) {
	for _, test := range EncoderTests {
		output, err := CR7Encoder(test.board, test.pasant, test.castling, test.seed)

		t.Log(output)

		if err != nil {
			t.Errorf("[!] %d", err)
			continue
		}

		if output != test.expected_CR7 {
			t.Errorf("[!] CR7 returned %q not equal to expected %q", output, test.expected_CR7)
		}

	}
}

func TestDecoder(t *testing.T) {
	for _, test := range DecoderTests {
		output_board, output_pasant, output_castling, output_seed, err := CR7Decoder(test.CR7)

		if err != nil {
			t.Errorf("[!] %d", err)
			continue
		}

		if output_seed != test.expected_seed {
			t.Errorf("[!] Seed returned %q not equal to expected %q", output_seed, test.expected_seed)
		}

		if output_board != test.expected_board {
			t.Errorf("[!] Board returned not equal to expected")
		}

		if output_pasant != test.expected_pasant {
			t.Errorf("[!] Pasant returned not equal to expected")
		}

		if output_castling != test.expected_castling {
			t.Errorf("[!] Castling returned not equal to expected")
		}

	}
}
