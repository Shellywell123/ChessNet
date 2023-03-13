package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sync"

	cnr "github.com/Shellywell123/ChessNet/chessnet/repositories/server/"
)

func play(MOVE string, GAMEFILE string) {

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

type test_struct struct {
	move string
	user string
}

func handle(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.move)
	GAMEFILE := t.user + ".PGN"
	play(t.move, GAMEFILE)
}

func StartServer() {

	http.HandleFunc("/chessnet", handle)

	fmt.Println("servering on 8090")
	http.ListenAndServe(":8090", nil)

	fmt.Println("server>")
}

func main() {
	StartServer()
}
