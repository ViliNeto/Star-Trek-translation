//Kinglon Specie translator is an application that look for Star Trek characters into Stapi API and show the character specie
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"./stapi"
	"./util"
)

//Defining this variable as global to reuse it on the unit tests
var specie string

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

//klingonTranslation -> Function used to do all translations requested
func klingonTranslation(characterName string) {

	//Pre decode the input to change character like 'ch', 'tlh', etc. That variable is gonna be used to translate the input into hexadecimal
	characterNamePreDecode := util.PreDecode(characterName)

	//Get the character specie through Stapi (character search and character UID)
	sp, err := stapi.GetSpecieCharacter(characterName)
	specie = sp
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error found: %s", err.Error()))
	}

	//Hexadecimal output (Using strings join to remove the braces on this array otherwise the output would be something like [0xF8E5, 0x​F8D6, ​​​0xF8E5, ​​0xF8E1, 0xF8D0])
	fmt.Println(strings.Join(util.TranslateToHexadecimal(characterNamePreDecode), ", "))
	//Char specie
	fmt.Println(specie)

}
