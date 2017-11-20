package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/imdario/mergo"

	"github.com/google/go-querystring/query"

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

// QueryParams are the parameters that can be used to query the API
type QueryParams struct {
	FlyFrom               string `url:"flyFrom"`
	FlyTo                 string `url:"to"`
	DateFrom              string `url:"dateFrom"` // TODO: Make dates the time.Time type?
	DateTo                string `url:"dateTo"`
	DaysInDestinationFrom int    `url:"daysInDestinationFrom"`
	DaysInDestinationTo   int    `url:"daysInDestinationTo"`
	Currency              string `url:"curr"`
}

// FindFlights finds flights (TODO: Real comment)
func FindFlights(userOptions QueryParams) (flights []models.Flight, err error) {
	options := QueryParams{
		Currency: "USD",
	}

	err = mergo.MergeWithOverwrite(&options, userOptions)

	if err != nil {
		return flights, err
	}

	querystring, err := query.Values(options)

	if err != nil {
		return flights, err
	}

	url := fmt.Sprintf(
		"https://api.skypicker.com/flights?%s",
		querystring.Encode(),
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
