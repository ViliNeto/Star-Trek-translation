package structures

//CharactersSearchResponse -> Response of the first API (search character) (http://stapi.co/api/v1/rest/character/search)
type CharactersSearchResponse struct {
	Page page `json:"page"`
	Sort struct {
		Clauses []interface{} `json:"clauses"`
	} `json:"sort"`
	Characters []character `json:"characters"`
}

//CharUIDspecie -> Detailed API based on UID to get information from original char (http://stapi.co/api/v1/rest/character/?uid=*******)
type CharUIDspecie struct {
	Character struct {
		CharacterSpecies []struct {
			Name string `json:"name"`
		} `json:"characterSpecies"`
	} `json:"character"`
}

type page struct {
	PageNumber       int  `json:"pageNumber"`
	PageSize         int  `json:"pageSize"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	TotalPages       int  `json:"totalPages"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
}
type character struct {
	UID                    string      `json:"uid"`
	Name                   string      `json:"name"`
	Gender                 string      `json:"gender"`
	YearOfBirth            int         `json:"yearOfBirth"`
	MonthOfBirth           interface{} `json:"monthOfBirth"`
	DayOfBirth             interface{} `json:"dayOfBirth"`
	PlaceOfBirth           interface{} `json:"placeOfBirth"`
	YearOfDeath            int         `json:"yearOfDeath"`
	MonthOfDeath           interface{} `json:"monthOfDeath"`
	DayOfDeath             interface{} `json:"dayOfDeath"`
	PlaceOfDeath           interface{} `json:"placeOfDeath"`
	Height                 interface{} `json:"height"`
	Weight                 interface{} `json:"weight"`
	Deceased               interface{} `json:"deceased"`
	BloodType              interface{} `json:"bloodType"`
	MaritalStatus          string      `json:"maritalStatus"`
	SerialNumber           interface{} `json:"serialNumber"`
	HologramActivationDate interface{} `json:"hologramActivationDate"`
	HologramStatus         interface{} `json:"hologramStatus"`
	HologramDateStatus     interface{} `json:"hologramDateStatus"`
	Hologram               bool        `json:"hologram"`
	FictionalCharacter     bool        `json:"fictionalCharacter"`
	Mirror                 bool        `json:"mirror"`
	AlternateReality       bool        `json:"alternateReality"`
}
