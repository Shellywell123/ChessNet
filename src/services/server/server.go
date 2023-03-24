package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	cnr "github.com/Shellywell123/ChessNet/src/repositories/server"
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

	GAMEFILE := t.user + ".PGN"
	log.Println("New Incoming Request:", t.user, t.move)
	play(t.move, GAMEFILE) // not sure why these are blank
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StartServer() {
	fmt.Println("Server version v0.01")
	cnr.PrintBootLogo()

	// handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/chessnet", handle)

	// start server
	fmt.Println("listening on 3001")
	http.ListenAndServe(":1234", nil)
}
