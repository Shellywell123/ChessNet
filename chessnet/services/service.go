package chessnetservice

import (
	"sync"
	cnr "github.com/Shellywell123/ChessNet/chessnet/chessnetrepository"
)

func play(MOVE string, GAMEFILE string) {

	cnr.printBootLogo()

	// Set up some MPSC stuff - specifically one reciever
	// wait groups - we dont use at all lol (for notifying us when a thread is finished doing its thing)
	// https://gobyexample.com/waitgroups
	wgRx := sync.WaitGroup{}
	wgRx.Add(1)

	// Start
	stdin, _ := cnr.startEngine()

	// Load
	cnr.loadGameToEngine(GAMEFILE, stdin)

	// Play Move
	cnr.sendChessMoveToEngine(MOVE, stdin)
	//fmt.Println()

	// Save
	cnr.saveGameFromEngine(GAMEFILE, stdin)

	// quit engine
	cnr.quitGameEngine(stdin)
}
