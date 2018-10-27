package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// GET Input from command line
	if len(os.Args) == 1 {
		fmt.Println("Char name not found")
		os.Exit(-1)
	}
	//Join all inputs typed on the terminal
	characterName := strings.Join(os.Args[1:], " ")

}
