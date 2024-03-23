package randomboy

import (
	"github.com/GermnGodoy/awami-bot/services/CR7"
)

func getMove(CR7pos string) {
	board, pasant, castling, seed, err := CR7.CR7Decoder(CR7pos)
}
