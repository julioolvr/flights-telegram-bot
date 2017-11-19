package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Flight struct {
	price int
}

type FlightResponse struct {
	Conversion struct {
		Usd int
	}
	Dtimeutc int
	Atimeutc int
}

type ApiResponse struct {
	Data []FlightResponse
}

func main() {
	url := "https://api.skypicker.com/flights?flyFrom=JFK&to=LAX&dateFrom=01/01/2018&dateTo=10/01/2018&daysInDestinationFrom=10&daysInDestinationTo=15&curr=USD"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request %s\n", err)
		return
	}

	var response ApiResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding response %s\n", err)
		return
	}

	fmt.Printf("Response %v\n", response)
}
