package chessnet

import (
	"io"
	"os"
)

func WriteStdOut(message string, stdin io.WriteCloser) {
	stdin.Write([]byte(message))
}

func FileExists(path string) bool {
	_, err := os.Open(path) // For read access.
	return err == nil

}
