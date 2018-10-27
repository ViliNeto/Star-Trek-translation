package stapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../structures"
)

//GetSpecieCharacter -> Try to find the first character to have the description of its specie
func GetSpecieCharacter(name string) (string, error) {
	specie := ""
	listUID, error := getCharacterUIDList(name)
	if error != nil {
		return "", error
	}
	println(listUID)
	return specie, nil

}

//Get all UID from searched character
func getCharacterUIDList(name string) (listUID []string, err error) {

	//POST to stapi API to search the desired Char detail
	var list []string
	body := []byte(fmt.Sprintf("name=%s", name))
	resp, err := http.Post("http://stapi.co/api/v1/rest/character/search",
		"x-www-form-urlencoded", bytes.NewBuffer(body))

	//If there is some error on the POST process
	if err != nil {
		return nil, err
	}
	//Close the request
	defer resp.Body.Close()

	//If the obtain result is different than OK
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error on Http request, StatusCode: %d", resp.StatusCode)
	}

	//In case we get CODE 200 (OK) read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	//If err == null means there is no error reading the response body
	if err != nil {
		return nil, err
	}

	//Convert the Json body into Golang structure
	var charResponse structures.CharactersSearchResponse
	err = json.Unmarshal(bodyBytes, &charResponse)

	//Build a list of characters UID
	for _, char := range charResponse.Characters {
		list = append(list, char.UID)
	}
	return list, nil
}
