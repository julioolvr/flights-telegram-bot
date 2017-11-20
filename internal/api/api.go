package api

import (
	"encoding/json"
	"net/http"

	"github.com/julioolvr/flights-telegram-bot/internal/models"
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
func FindFlights() (flights []models.Flight, err error) {
	url := "https://api.skypicker.com/flights?flyFrom=JFK&to=LAX&dateFrom=01/01/2018&dateTo=10/01/2018&daysInDestinationFrom=10&daysInDestinationTo=15&curr=USD"

	resp, err := http.Get(url)

	if err != nil {
		return flights, err
	}

	var response apiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return flights, err
	}

	numberOfFlights := len(response.Data)
	flights = make([]models.Flight, numberOfFlights)

	for i := 0; i < numberOfFlights; i++ {
		flights[i] = models.Flight{Price: response.Data[i].Conversion.Usd}
	}

	return flights, err
}
