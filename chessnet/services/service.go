package chessnet

import (
	"sync"

	cnr "github.com/Shellywell123/ChessNet/chessnet/repositories"
)

// const GNUMOVED = "My move is :"
// const GNUSTARTED = "There is NO WARRANTY, to the extent permitted by law."

// // Relay input from gnuchess to the channel
// func readGnuChessInput(r io.Reader, channel chan string) {
// 	scanner := bufio.NewScanner(r)
// 	for scanner.Scan() {
// 		channel <- scanner.Text() // syntax for stuff being sent to the channel (sender/transmitter) Ref: https://gobyexample.com/channels
// 	}
// }

// // listener Function runs an infinate loop:
// // - It is THE listener/reciever of the channel
// // - when a string is sent INTO channel from elsewhere, it is assigned into the variable `msg`
// func listener(channel chan string, stdin io.WriteCloser) {
// 	// go doesnt have any other form of loops - so has quirky stuff like this
// 	for {
// 		// msg is assigned the value written INTO the channel from elsewhere. It's the reciever.
// 		// msg := <-channel // syntax for stuff bing recievd bt the channel (listener/reciever) Ref: https://gobyexample.com/channels
// 		// If the msg (which is a string) contains the substring GNUMOVE "My move is :", then write a second move - in this case d4
// 		// select {
// 		// case msg := <-channel:
// 		// 	//fmt.Println("# " + msg)
// 		// 	switch {
// 		// 	case strings.Contains(msg, GNUSTARTED):
// 		// 		cn.LoadGameToEngine("test.pgn", stdin)
// 		// 		break

// 		// 	case strings.Contains(msg, ""):
// 		// 		cn.SendChessMoveToEngine("e4", stdin)
// 		// 		break

// 		// 	// case strings.Contains(msg, GNUSTARTED):
// 		// 	// 	cn.SendChessMoveToEngine("e4", stdin)
// 		// 	// 	break

// 		// 	case strings.Contains(msg, GNUMOVED):
// 		// 		cn.QuitGameEngine(stdin)
// 		// 		break
// 		// 	}

// 		// 	// if strings.Contains(msg, GNUSTARTED) {

// 		// 	// }
// 		// 	// if strings.Contains(msg, GNUSTARTED) {
// 		// 	// 	cn.SendChessMoveToEngine("e4", stdin)
// 		// 	// }
// 		// 	// if strings.Contains(msg, GNUMOVED) {
// 		// 	// 	cn.QuitGameEngine(stdin)
// 		// 	// }
// 		// 	fmt.Printf("%s\n", msg)
// 		// 	break
// 		// }

// 	}
// }

func Play(MOVE string, GAMEFILE string) {

	cnr.PrintBootLogo()

	// Set up some MPSC stuff - specifically one reciever
	// wait groups - we dont use at all lol (for notifying us when a thread is finished doing its thing)
	// https://gobyexample.com/waitgroups
	wgRx := sync.WaitGroup{}
	wgRx.Add(1)

	//events := make(chan string) // chan is a type (channel)

	// go = a go routine
	//go listener(events, stdin) // We are running a goroutine ("think: magic thread") calling the function writeOut
	stdin, _ := cnr.StartEngine()

	//cmd.Start()

	// Create a stdout reader
	// r := bufio.NewReader(stdout) // input out put buffer from lib we have imported
	// go readGnuChessInput(r, events)

	// Load
	cnr.LoadGameToEngine(GAMEFILE, stdin)

	// Play Move
	cnr.SendChessMoveToEngine(MOVE, stdin)
	//fmt.Println()

	// Save
	cnr.SaveGameFromEngine(GAMEFILE, stdin)

	// quit engine
	cnr.QuitGameEngine(stdin)

	//cmd.Wait() // waits for gnuchess to exit
	//time.Sleep(20 * time.Second)
}

// Server side
// - API - Use case that will interface with... entity

// Client side
// - CLI - Use case that will interface with... entity
