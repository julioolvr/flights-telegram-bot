package flightService

import (
	"time"

	"github.com/julioolvr/flights-telegram-bot/internal/api"
	"github.com/julioolvr/flights-telegram-bot/internal/models/flight"
)

// findFlightsOptions are the options that can be passed to Search in order to
// filter the flights being looked for.
type findFlightsOptions struct {
	FlyFrom               string
	FlyTo                 string
	DateFrom              string // TODO: Make dates the time.Time type?
	DateTo                string
	DaysInDestinationFrom int
	DaysInDestinationTo   int
	Currency              string
	Limit                 int
}

// FlightFinder is used to build a query for flights and then execute
// such query in order to retrieve the results.
type FlightFinder struct {
	options findFlightsOptions
}

// Find initializes a FlightFinder, on which calls can be chained in
// order to filter flights to retrieve.
func Find() FlightFinder {
	return FlightFinder{options: findFlightsOptions{}}
}

// From filters flights by the origin airport.
func (finder FlightFinder) From(from string) FlightFinder {
	finder.options.FlyFrom = from
	return finder
}

// To filters flights by the destination airport.
func (finder FlightFinder) To(to string) FlightFinder {
	finder.options.FlyTo = to
	return finder
}

// DepartureDateRange sets an initial and end date to search for departure
// dates of flights.
func (finder FlightFinder) DepartureDateRange(dateFrom string, dateTo string) FlightFinder {
	finder.options.DateFrom = dateFrom
	finder.options.DateTo = dateTo
	return finder
}

// DateFrom filters for flights with a departure date starting on this date.
func (finder FlightFinder) DateFrom(dateFrom string) FlightFinder {
	finder.options.DateFrom = dateFrom
	return finder
}

// DateTo filters for flights with a departure date ending on this date.
func (finder FlightFinder) DateTo(dateTo string) FlightFinder {
	finder.options.DateTo = dateTo
	return finder
}

// DaysInDestination sets a minimum and maximum dates of stay at the destination.
func (finder FlightFinder) DaysInDestination(daysFrom int, daysTo int) FlightFinder {
	finder.options.DaysInDestinationFrom = daysFrom
	finder.options.DaysInDestinationTo = daysTo
	return finder
}

// DaysInDestinationFrom sets a minimum days of stay at the destination.
func (finder FlightFinder) DaysInDestinationFrom(daysFrom int) FlightFinder {
	finder.options.DaysInDestinationFrom = daysFrom
	return finder
}

// DaysInDestinationTo sets a maximum days of stay at the destination.
func (finder FlightFinder) DaysInDestinationTo(daysTo int) FlightFinder {
	finder.options.DaysInDestinationTo = daysTo
	return finder
}

// Currency sets which currency should be included for the flight prices.
func (finder FlightFinder) Currency(currency string) FlightFinder {
	finder.options.Currency = currency
	return finder
}

// Limit sets the maximum number of flights to retrieve.
func (finder FlightFinder) Limit(limit int) FlightFinder {
	finder.options.Limit = limit
	return finder
}

// Exec executes the request to the API with all the parameters set previously.
func (finder FlightFinder) Exec() (flights []flight.Flight, err error) {
	return search(finder.options)
}

// search uses the API to find flights according to the given options and returns
// a list of flight.Flight entities.
func search(options findFlightsOptions) (flights []flight.Flight, err error) {
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
