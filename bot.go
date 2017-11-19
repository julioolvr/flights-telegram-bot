package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type flight struct {
	price int
}

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

func main() {
	url := "https://api.skypicker.com/flights?flyFrom=JFK&to=LAX&dateFrom=01/01/2018&dateTo=10/01/2018&daysInDestinationFrom=10&daysInDestinationTo=15&curr=USD"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request %s\n", err)
		return
	}

	var response apiResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding response %s\n", err)
		return
	}

	fmt.Printf("Response %v\n", response)
}
