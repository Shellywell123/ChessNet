package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	cnr "github.com/Shellywell123/ChessNet/src/repositories/server"
)

type test_struct struct {
	Move string
	User string
}

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

func handle(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	GAMEFILE := t.User + ".PGN"

	// print server side
	log.Println("New Incoming Request:", t.User, t.Move)

	// print client side
	fmt.Fprintf(rw, "Move recieved: %s from %s\n", t.Move, t.User)

	play(t.Move, GAMEFILE) // not sure why these are blank
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "check mate.\n")
}

func StartServer() {
	fmt.Println("Server version v0.01")
	cnr.PrintBootLogo()

	// handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/game", handle)

	// start server
	port := ":3000"
	fmt.Printf("listening on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
