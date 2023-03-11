package chessnetrepository

import (
	"fmt"
	"log"
	"os"
)

func printBootLogo() {
	contents, err := os.ReadFile("assets/ascii-boot-logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(contents))
}
