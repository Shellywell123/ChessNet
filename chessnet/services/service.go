package services

import (
	"sync"
	
	cnr "github.com/Shellywell123/ChessNet/chessnet/repositories"
	
)

func play(MOVE string, GAMEFILE string) {

	cnr.PrintBootLogo()

	// Set up some MPSC stuff - specifically one reciever
	// wait groups - we dont use at all lol (for notifying us when a thread is finished doing its thing)
	// https://gobyexample.com/waitgroups
	wgRx := sync.WaitGroup{}
	wgRx.Add(1)

	// Start
	stdin, _ := cnr.StartEngine()

	// Load
	cnr.LoadGameToEngine(GAMEFILE, stdin)

	// Play Move
	cnr.SendChessMoveToEngine(MOVE, stdin)
	//fmt.Println()

	// Save
	cnr.SaveGameFromEngine(GAMEFILE, stdin)

	// quit engine
	cnr.QuitGameEngine(stdin)
}
