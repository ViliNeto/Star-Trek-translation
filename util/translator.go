package util

import (
	"strings"
)

//TranslateToHexadecimal -> Translate the name of the Character to hexadecimal
func TranslateToHexadecimal(charNamePreDecoded string) []string {
	//Load Klingon alphabet to hexadecimal representation
	klingonAlphabet := LoadKlingonAlphabet()

	var translation []string
	//runs for every single letter into the char name to build hexadecimal representation (eg: Uhura -> [0]U [1]=h [2]u [3]r [4]a)
	for i := range charNamePreDecoded {
		singleChar := string(charNamePreDecoded[i])
		klingonRespectiveChar := klingonAlphabet[singleChar]

		if klingonRespectiveChar == "" {
			if strings.ToUpper(singleChar) == singleChar {
				singleChar = strings.ToLower(singleChar)
			} else {
				singleChar = strings.ToUpper(singleChar)
			}
			klingonRespectiveChar = klingonAlphabet[singleChar]
		}

		translation = append(translation, klingonRespectiveChar)
	}

	return translation
}
