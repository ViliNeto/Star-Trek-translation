package util

import "strings"

//LoadKlingonAlphabet -> Function responsible to map all chars to its respective hexadecimal representation
func LoadKlingonAlphabet() map[string]string {

	klingonAlphabet := map[string]string{
		"a":  "0xF8D0",
		"v":  "0xF8D0",
		"b":  "0x​F8D1",
		"D":  "0xF8D3",
		"e":  "0xF8D4",
		"H":  "0x​F8D6",
		"I":  "0x​F8D7",
		"j":  "0x​F8D8",
		"l":  "0x​F8D9",
		"m":  "0x​F8DA",
		"n":  "0x​F8DB",
		"#":  "0x​F8DC",
		"o":  "0x​F8DD",
		"p":  "0x​F8DE",
		"q":  "0x​F8DF",
		"Q":  "0x​F8E0",
		"r":  "​​0xF8E1",
		"S​": "​​​0xF8E2",
		"s":  "​​​0xF8E2",
		"t​": "​​​0xF8E3",
		"u":  "​​​0xF8E5",
		"y​": "0x​​​F8E8",
		"$":  "​​​0xF8E4",
		"0":  "0xF8F0",
		"1":  "0xF8F1",
		"2":  "0xF8F2",
		"3":  "0xF8F3",
		"4":  "0xF8F4",
		"5":  "0xF8F5",
		"6":  "0xF8F6",
		"7":  "0xF8F7",
		"8":  "0xF8F8",
		"9":  "0xF8F9",
		"!":  "0x​​​F8D2",
		"@":  "​0xF8D5",
		"w":  "0x​​​F8E7",
		"’":  "0x​​F8E9",
		".":  "0xF8FD",
		",":  "0xF8FE",
		"t":  "0x0074",
		" ":  "0x0020",
		"ç":  "0xF8FF",
	}
	return klingonAlphabet
}

//PreDecode -> Decode main chars to its representations
func PreDecode(charName string) string {
	charName = strings.Replace(charName, "&", "", -1)
	charName = strings.Replace(charName, "!", "", -1)
	charName = strings.Replace(charName, "#", "", -1)
	charName = strings.Replace(charName, "$", "", -1)
	charName = strings.Replace(charName, "ç", "", -1)

	charName = strings.Replace(charName, "ch", "&", -1)
	charName = strings.Replace(charName, "gh", "!", -1)
	charName = strings.Replace(charName, "ng", "#", -1)
	charName = strings.Replace(charName, "tlh", "$", -1)

	return charName
}
