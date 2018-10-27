package util

import (
	"fmt"
	"strings"
	"testing"
)

//Testing the util package
func TestTranslator(t *testing.T) {

	allStrings := []string{
		"a",
		"v",
		"b",
		"D",
		"e",
		"H",
		"I",
		"j",
		"l",
		"m",
		"n",
		"#",
		"o",
		"p",
		"q",
		"Q",
		"r",
		"S​",
		"s",
		"t​",
		"u",
		"y​",
		"$",
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"!",
		"@",
		"w",
		"’",
		".",
		",",
		"s",
		" ",
		"ç",
		"-", //Doesn't exists in the original map
		"}", //Doesn't exists in the original map
	}

	//Test all dictionary to check if is there any char that is not being contemplated
	klingonAlphabet := LoadKlingonAlphabet()
	for i := range allStrings {
		klingonRespectiveChar := klingonAlphabet[allStrings[i]]
		if klingonRespectiveChar != "" {
			fmt.Printf("PASSED: for %s\n", allStrings[i])
		} else {
			if (allStrings[i] == "-") || (allStrings[i] == "}") {
				fmt.Printf("PASSED: for %s, because it doesn't exists in the original map\n", allStrings[i])
			} else {
				fmt.Printf("FAILED: for %s\n", allStrings[i])
			}
		}
	}
	//Predecode the user input
	decode := PreDecode("&!#$ç4")
	if decode == "4" {
		println("PASSED into 1st decode test")
	} else {
		println("FAILED into 1st decode test")
	}
	decode = PreDecode("chghngtlh")
	if decode == "&!#$" {
		println("PASSED into 2nd decode test")
	} else {
		println("FAILED into 2nd decode test")
	}

	//Translating to Hexadecimal
	translation := TranslateToHexadecimal("Uhura")
	if strings.Join(translation, " ") == "​​​0xF8E5 0x​F8D6 ​​​0xF8E5 ​​0xF8E1 0xF8D0" {
		println("PASSED into 1st Translation test")
	} else {
		println("FAILED into 1st Translation test")
	}

	//Trying to translate using character not registered ("-")
	translation = TranslateToHexadecimal("Uhura-")
	println(strings.Join(translation, " "))
	if strings.Join(translation, " ") == "​​​0xF8E5 0x​F8D6 ​​​0xF8E5 ​​0xF8E1 0xF8D0 " {
		println("PASSED into 2nd Translation test")
	} else {
		println("FAILED into 2nd Translation test")
	}
}
