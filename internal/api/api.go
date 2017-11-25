package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/imdario/mergo"

	"github.com/google/go-querystring/query"
)

type FlightResponse struct {
	Conversion struct {
		Usd int
	}
	Dtimeutc     int64
	Atimeutc     int64
	Flyfrom      string
	Flyto        string
	Route        []RouteResponse
	BookingToken string `json:"booking_token"`
}

type RouteResponse struct {
	Atimeutc int64
	Dtimeutc int64
	FlyFrom  string
	FlyTo    string
}

type ApiResponse struct {
	Data []FlightResponse
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
	Limit                 int    `url:"limit"`
}

// FindFlights finds flights (TODO: Real comment)
func FindFlights(userOptions QueryParams) (response ApiResponse, err error) {
	options := QueryParams{
		Currency: "USD",
		Limit:    200,
	}

	err = mergo.MergeWithOverwrite(&options, userOptions)

	if err != nil {
		return response, err
	}

	querystring, err := query.Values(options)

	if err != nil {
		return response, err
	}

	url := fmt.Sprintf(
		"https://api.skypicker.com/flights?%s",
		querystring.Encode(),
	)

	resp, err := http.Get(url)

	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err
}
