package flightService

import (
	"time"

	"github.com/julioolvr/flights-telegram-bot/internal/api"
	"github.com/julioolvr/flights-telegram-bot/internal/models/flight"
)

// FindFlightsOptions are the options that can be passed to Search in order to
// filter the flights being looked for.
type FindFlightsOptions struct {
	FlyFrom               string
	FlyTo                 string
	DateFrom              string // TODO: Make dates the time.Time type?
	DateTo                string
	DaysInDestinationFrom int
	DaysInDestinationTo   int
	Currency              string
	Limit                 int
}

// Search uses the API to find flights according to the given options and returns
// a list of flight.Flight entities.
func Search(options FindFlightsOptions) (flights []flight.Flight, err error) {
	response, err := api.FindFlights(api.QueryParams{
		FlyFrom:               options.FlyFrom,
		FlyTo:                 options.FlyTo,
		DateFrom:              options.DateFrom,
		DateTo:                options.DateTo,
		DaysInDestinationFrom: options.DaysInDestinationFrom,
		DaysInDestinationTo:   options.DaysInDestinationTo,
		Currency:              options.Currency,
		Limit:                 options.Limit,
	})

	if err != nil {
		return flights, err
	}

	numberOfFlights := len(response.Data)
	flights = make([]flight.Flight, numberOfFlights)

	for i, flightData := range response.Data {
		flights[i] = buildFlightFromResponse(flightData)
	}

	return flights, err
}

func buildFlightFromResponse(response api.FlightResponse) flight.Flight {
	numberOfSegments := len(response.Route)
	segments := make([]flight.RouteSegment, numberOfSegments)

	for i, segmentData := range response.Route {
		segments[i] = buildRouteSegmentFromResponse(segmentData)
	}

	return flight.Flight{
		Price: response.Conversion.Usd,
		Routes: []flight.Route{flight.Route{
			From:     flight.Location{Airport: flight.AirportCode(response.Flyfrom)},
			To:       flight.Location{Airport: flight.AirportCode(response.Flyto)},
			Segments: segments,
		}},
	}
}

func buildRouteSegmentFromResponse(response api.RouteResponse) flight.RouteSegment {
	return flight.RouteSegment{
		From:      flight.Location{Airport: flight.AirportCode(response.FlyFrom)},
		To:        flight.Location{Airport: flight.AirportCode(response.FlyTo)},
		Departure: time.Unix(response.Dtimeutc, 0),
		Arrival:   time.Unix(response.Atimeutc, 0),
	}
}
