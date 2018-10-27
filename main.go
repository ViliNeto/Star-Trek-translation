package main

import (
	"fmt"
	"os"
	"strings"

	"./util"
)

func main() {
	// GET Input from command line
	if len(os.Args) == 1 {
		fmt.Println("Char name not found")
		os.Exit(-1)
	}
	//Join all inputs typed on the terminal
	characterName := strings.Join(os.Args[1:], " ")

	klingonTranslation(characterName)

}

func klingonTranslation(characterName string) {

	//Pre decode the input to change character like 'ch', 'tlh', etc. That variable is gonna be used to translate the input into hexadecimal
	characterNamePreDecode := util.PreDecode(characterName)

	//Original input
	fmt.Println(characterName)
	//Hexadecimal output (Using strings join to remove the braces on this array otherwise the output would be something like [0xF8E5, 0x​F8D6, ​​​0xF8E5, ​​0xF8E1, 0xF8D0])
	fmt.Println(strings.Join(util.TranslateToHexadecimal(characterNamePreDecode), ", "))

}
