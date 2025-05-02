package address

import (
	"encoding/json"
	"fmt"
	"io"
)

type ZipDatabase map[string]Address

type Address struct {
	ZipCode         string   `json:"cep"`
	State           string   `json:"state"`
	InitialDistrict District `json:"initialDistrict"`
	FinalDistrict   District `json:"finalDistrict"`
	AddressLine     string   `json:"addressLine"`
	City            string   `json:"city"`
	Deleted         string   `json:"deleted"`
}

type District struct {
	Key  int    `json:"key"`
	Name string `json:"name"`
}

func FromJson(jsonData io.Reader) (database ZipDatabase, err0 error) {
	jsonParser := json.NewDecoder(jsonData)

	err := jsonParser.Decode(&database)
	if err != nil {
		return nil, fmt.Errorf("parsing file '%s': %w", jsonData, err)
	}

	return database, nil
}

func (address *Address) ToJson() (jsonData string, err0 error) {
	jsonBytes, err := json.Marshal(address)
	if err != nil {
		return "", fmt.Errorf("parsing file '%s': %w", jsonData, err)
	}

	return string(jsonBytes), nil
}
