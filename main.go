package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	cn "github.com/shellywell123/ChessNet/ChessNet"
)

const GNUMOVE = "My move is : e4"
const GAMEFILE = "GAME.PGN"

// Relay input from gnuchess to the channel
func readGnuChessInput(r io.Reader, channel chan string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		channel <- scanner.Text() // syntax for stuff being sent to the channel (sender/transmitter) Ref: https://gobyexample.com/channels
	}
}

// Function runs an infinate loop:
// - It is THE listener/reciever of the channel
// - when a string is sent INTO channel from elsewhere, it is assigned into the variable `msg`
func listener(channel chan string, stdin io.WriteCloser) {
	// go doesnt have any other form of loops - so has quirky stuff like this
	for {
		// msg is assigned the value written INTO the channel from elsewhere. It's the reciever.
		msg := <-channel // syntax for stuff bing recievd bt the channel (listener/reciever) Ref: https://gobyexample.com/channels
		// If the msg (which is a string) contains the substring GNUMOVE "My move is :", then write a second move - in this case d4
		if strings.Contains(msg, GNUMOVE) {
			WriteStdOut("d4\n", stdin)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {

	// Set up some MPSC stuff - specifically one reciever
	// wait groups - we dont use at all lol (for notifying us when a thread is finished doing its thing)
	// https://gobyexample.com/waitgroups
	wgRx := sync.WaitGroup{}
	wgRx.Add(1)

	events := make(chan string) // chan is a type (channel)
	cmd := exec.Command("gnuchess")
	cmd.Stderr = os.Stderr        // gnuchess standard errors are pasesed as cmd errors
	stdin, err := cmd.StdinPipe() // assign stdin variable to the command stdinput - the thing we'll message gnuchess for. This is the same type of 'pipe' as | in bash
	if err != nil {
		log.Fatalf("Error getting stdin: %s", err.Error()) // err.Error() think of as err.ToString()
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error getting stdout: %s", err.Error())
	}

	// go = a go routine
	go listener(events, stdin) // We are running a goroutine ("think: magic thread") calling the function writeOut

	// Load
	e.loadGameToEngine(GAMEFILE, stdin)

	// Play Move
	e.sendChessMoveToEngine(GNUMOVE, stdin)

	// Save
	e.saveGameFromEngine(GAMEFILE, stdin)

	// quit engine
	e.quitGameEngine(stdin)

	// Create a stdout reader
	r := bufio.NewReader(stdout) // input out put buffer from lib we have imported
	go readGnuChessInput(r, events)
	cmd.Start()
	cmd.Wait()
}

// Server side
// - API - Use case that will interface with... entity

// Client side
// - CLI - Use case that will interface with... entity
