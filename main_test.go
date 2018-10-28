package main

import (
	"fmt"
	"strings"
	"testing"

	"./util"
)

//TestMain -> Test "main.go" mainly interactions
func TestMain(t *testing.T) {
	//Uhura will find a specie
	klingonTranslation("Uhura")
	characterNamePreDecode := util.PreDecode("Uhura")
	if specie == "Human" && strings.Join(util.TranslateToHexadecimal(characterNamePreDecode), ", ") == "​​​0xF8E5, 0x​F8D6, ​​​0xF8E5, ​​0xF8E1, 0xF8D0" {
		fmt.Println("PASSED: Uhura is Human")
	}
	//Loren Shriver exists but will not find a specie
	klingonTranslation("Loren Shriver")
	characterNamePreDecode = util.PreDecode("Loren Shriver")
	if specie == "Specie Not found..." && strings.Join(util.TranslateToHexadecimal(characterNamePreDecode), ", ") == "0x​F8D9, 0x​F8DD, ​​0xF8E1, 0xF8D4, 0x​F8DB, 0x0020, ​​​0xF8E2, 0x​F8D6, ​​0xF8E1, 0x​F8D7, 0xF8D0, 0xF8D4, ​​0xF8E1" {
		fmt.Println("PASSED: Loren Shriver has no specie registered")
	}
	//Dwuas does not exists [I made that name, sorry. =) ]
	klingonTranslation("Dwuas")
	characterNamePreDecode = util.PreDecode("Dwuas")
	if specie == "Specie Not found..." && strings.Join(util.TranslateToHexadecimal(characterNamePreDecode), ", ") == "0xF8D3, 0x​​​F8E7, ​​​0xF8E5, 0xF8D0, ​​​0xF8E2" {
		fmt.Println("PASSED: Dwuast has no specie registered")
	}

}
