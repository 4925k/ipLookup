package iplookup

import (
	"encoding/json"
	"fmt"
	"io"
	"lookup/flatten"
	"net/http"
)

//uses ipregistry.com to look up given address
//the first 100k requests are free
//currently uses token provided to Dibek's account
//Returns a struct of the response
func IpLookup(ip string) (map[string]string, error) {
	response := make(map[string]interface{})
	apiKey := "w568xr2nj82mbl" //apikey is generated in your ip registry user account
	fields := "connection,location,security,-location.country.population,-location.country.population_density,-location.country.flag,-location.in_eu"

	//create url request
	request := fmt.Sprintf("https://api.ipregistry.co/%s?key=%s&fields=%s", ip, apiKey, fields)
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	//flatten struct
	flat, err := flatten.Flatten(response, "", flatten.DotStyle)
	if err != nil {
		return nil, err
	}

	//flatten nested map to map[string]string
	//to go with the main program
	return flat, nil
}
