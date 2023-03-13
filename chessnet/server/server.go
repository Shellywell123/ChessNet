package server

import (
	"fmt"
	"net/http"

	cns "github.com/Shellywell123/ChessNet/chessnet/server/services"
)

const GAMEFILE = "GAME.PGN"

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func StartServer() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
	cns.Play("e2e4", GAMEFILE)
}
