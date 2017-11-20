package api

import (
	"encoding/json"
	"net/http"
)

type flightResponse struct {
	Conversion struct {
		Usd int
	}
	Dtimeutc int
	Atimeutc int
}

type apiResponse struct {
	Data []flightResponse
}

// FindFlights finds flights (TODO: Real comment)
func FindFlights() (response apiResponse, err error) {
	url := "https://api.skypicker.com/flights?flyFrom=JFK&to=LAX&dateFrom=01/01/2018&dateTo=10/01/2018&daysInDestinationFrom=10&daysInDestinationTo=15&curr=USD"

	resp, err := http.Get(url)

	if err != nil {
		return apiResponse{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}
