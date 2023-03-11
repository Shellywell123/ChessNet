package main

import (
	cns "github.com/Shellywell123/ChessNet/chessnet/services"
)

const GAMEFILE = "GAME.PGN"

func main() {
	cns.Play("e2e4", GAMEFILE)
}
