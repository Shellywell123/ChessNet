package repositories

import (
	"io"
	"os"
)

func writeStdOut(message string, stdin io.WriteCloser) {
	stdin.Write([]byte(message))
}

func fileExists(path string) bool {
	_, err := os.Open(path) // For read access.
	return err == nil
}
