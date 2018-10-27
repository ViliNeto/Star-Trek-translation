package structures

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEstructures(t *testing.T) {
	correctCharSearchResponse := "{" +
		"    \"page\": {" +
		"        \"pageNumber\": 0," +
		"        \"pageSize\": 50," +
		"        \"numberOfElements\": 1," +
		"        \"totalElements\": 3," +
		"        \"totalPages\": 3," +
		"        \"firstPage\": true," +
		"        \"lastPage\": false" +
		"    }," +
		"    \"sort\": {" +
		"        \"clauses\": []" +
		"    }," +
		"    \"characters\": [" +
		"        {" +
		"            \"uid\": \"CHMA0000023576\"," +
		"            \"name\": \"Nyota Uhura!\"," +
		"            \"gender\": \"F\"," +
		"            \"yearOfBirth\": 2239," +
		"            \"monthOfBirth\": null," +
		"            \"dayOfBirth\": null," +
		"            \"placeOfBirth\": null," +
		"            \"yearOfDeath\": 2268," +
		"            \"monthOfDeath\": null," +
		"            \"dayOfDeath\": null," +
		"            \"placeOfDeath\": null," +
		"            \"height\": null," +
		"            \"weight\": null," +
		"            \"deceased\": null," +
		"            \"bloodType\": null," +
		"            \"maritalStatus\": \"SINGLE\"," +
		"            \"serialNumber\": null," +
		"            \"hologramActivationDate\": null," +
		"            \"hologramStatus\": null," +
		"            \"hologramDateStatus\": null," +
		"            \"hologram\": false," +
		"            \"fictionalCharacter\": false," +
		"            \"mirror\": true," +
		"            \"alternateReality\": false" +
		"        }" +
		"    ]" +
		"}"
	//Convert the Json body into Golang structure
	var charResult CharactersSearchResponse
	err := json.Unmarshal([]byte(correctCharSearchResponse), &charResult)
	if err != nil {
		panic(err)
	} else {
		if charResult.Characters[0].Name != "" {
			fmt.Printf("PASS: Json Unmarshal passed. Eg.: Name = %s\n", charResult.Characters[0].Name)
		} else {
			panic(err)
		}
	}

	incorrectCharSearchResponse := "{" +
		"    \"glossary\": {" +
		"        \"title\": \"example glossary\"," +
		"		\"GlossDiv\": {" +
		"            \"title\": \"S\"," +
		"			\"GlossList\": {" +
		"                \"GlossEntry\": {" +
		"                    \"ID\": \"SGML\"," +
		"					\"SortAs\": \"SGML\"," +
		"					\"GlossTerm\": \"Standard Generalized Markup Language\"," +
		"					\"Acronym\": \"SGML\"," +
		"					\"Abbrev\": \"ISO 8879:1986\"," +
		"					\"GlossDef\": {" +
		"                        \"para\": \"A meta-markup language, used to create markup languages such as DocBook.\"," +
		"						\"GlossSeeAlso\": [\"GML\", \"XML\"]" +
		"                    }," +
		"					\"GlossSee\": \"markup\"" +
		"                }" +
		"            }" +
		"        }" +
		"    }" +
		"}"

	//Convert the Json body into Golang structure
	var charResult2 CharactersSearchResponse
	err = json.Unmarshal([]byte(incorrectCharSearchResponse), &charResult2)
	println("err: ", err)
	if err != nil {
		panic(err)
	} else {
		if charResult2.Page.FirstPage == false {
			fmt.Printf("PASS: Json Unmarshal passed. Eg.: Name = %s\n", charResult2.Characters[0].Name)
		} else {
			panic(err)
		}
	}

}
