// THIS SHOULD:
// - Entity:(own contained functions)
// - Recieve chess notation FROM a chess engine - is this needed if we can passs the board

package chessnet

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// - Start a chess engine
func StartEngine() (io.WriteCloser, io.ReadCloser) {
	cmd := exec.Command("gnuchess --xboard")
	cmd.Stderr = os.Stderr        // gnuchess standard errors are pasesed as cmd errors
	stdin, err := cmd.StdinPipe() // assign stdin variable to the command stdinput - the thing we'll message gnuchess for. This is the same type of 'pipe' as | in bash
	if err != nil {
		log.Fatalf("Error getting stdin: %s", err.Error()) // err.Error() think of as err.ToString()
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error getting stdout: %s", err.Error())
	}
	return stdin, stdout
}

// - Send chess notation to a chess engine - gnuchess already does this
func SendChessMoveToEngine(MOVE string, stdin io.WriteCloser) {
	fmt.Println("moving " + MOVE)
	WriteStdOut(MOVE+"\n", stdin)
}

// - Load a saved game of chess (PGN/FEM)
func LoadGameToEngine(GAMEFILE string, stdin io.WriteCloser) {
	if FileExists(GAMEFILE) {
		WriteStdOut("pgnload "+GAMEFILE+" \n", stdin)
		fmt.Println("loaded " + GAMEFILE)
	} else {
		emptyFile, err := os.Create(GAMEFILE)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("created " + GAMEFILE)
		emptyFile.Close()
	}
}

// - Save a game of chess (note - this does not mean persistence like on disk) (PGN/FEM)
func SaveGameFromEngine(GAMEFILE string, stdin io.WriteCloser) {
	WriteStdOut("pgnsave "+GAMEFILE+"\n", stdin)
	fmt.Println("saved " + GAMEFILE)
}

// exit out of gnuchess
func QuitGameEngine(stdin io.WriteCloser) {
	WriteStdOut("quit\n", stdin)
	fmt.Println("exited.")
}
