package repositories

import (
	"fmt"
	"log"
	"os"
)

func PrintBootLogo() {
	contents, err := os.ReadFile("src/assets/ascii-boot-logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(contents))
}
