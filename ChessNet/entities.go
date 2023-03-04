// THIS SHOULD:
// - Entity:(own contained functions)
// - Recieve chess notation FROM a chess engine - is this needed if we can passs the board

package chessnet

import (
	"io"
	"log"
	"os"
	"strings"
)

// - Start a chess engine
func startEngine() {

}

// - Send chess notation to a chess engine - gnuchess already does this
func sendChessMoveToEngine(GNUMOVE string, stdin io.WriteCloser) {
	WriteStdOut(strings.Split(GNUMOVE, ":")[1]+"\n", stdin)
}

// - Load a saved game of chess (PGN/FEM)
func loadGameToEngine(GAMEFILE string, stdin io.WriteCloser) {
	if FileExists(GAMEFILE) {
		WriteStdOut("pgnload "+GAMEFILE+"\n", stdin)
	} else {
		emptyFile, err := os.Create(GAMEFILE)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emptyFile)
		emptyFile.Close()
	}
}

// - Save a game of chess (note - this does not mean persistence like on disk) (PGN/FEM)
func saveGameFromEngine(GAMEFILE string, stdin io.WriteCloser) {
	WriteStdOut("pgnsave "+GAMEFILE+"\n", stdin)
}

func quitGameEngine(stdin io.WriteCloser) {
	WriteStdOut("quit\n", stdin)
}
