package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julioolvr/flights-telegram-bot/internal/models"
)

type flightResponse struct {
	Conversion struct {
		Usd int
	}
	Dtimeutc int64
	Atimeutc int64
}

type apiResponse struct {
	Data []flightResponse
}

// FindFlights finds flights (TODO: Real comment)
func FindFlights(fromAirport string, toAirport string) (flights []models.Flight, err error) {
	url := fmt.Sprintf(
		"https://api.skypicker.com/flights?flyFrom=%s&to=%s&dateFrom=01/01/2018&dateTo=10/01/2018&daysInDestinationFrom=10&daysInDestinationTo=15&curr=USD",
		fromAirport,
		toAirport,
	)

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
		flights[i] = models.Flight{
			Price:    response.Data[i].Conversion.Usd,
			Depature: time.Unix(response.Data[i].Dtimeutc, 0),
			Arrival:  time.Unix(response.Data[i].Atimeutc, 0),
		}
	}

	return flights, err
}
