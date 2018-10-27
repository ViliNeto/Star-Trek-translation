package stapi

import (
	"os"
	"testing"
)

func TestStapiAPI(t *testing.T) {
	//Happy scenario
	specie, err := GetSpecieCharacter("Uhura")
	if err != nil {
		println("FAILED: Error to retreive Char specie")
		os.Exit(1)
	} else {
		if specie == "Human" {
			println("PASS: Uhura is a Human")
		}
	}
	//When the character doesn't exist
	specie, err = GetSpecieCharacter("Vili")
	if err != nil {
		println("FAILED: Error to retreive Char specie")
		os.Exit(1)
	} else {
		if specie == "Specie Not found..." {
			println("PASS: Vili is not a Star Trek character")
		}
	}
}
