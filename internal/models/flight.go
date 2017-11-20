package models

import (
	"fmt"
	"time"
)

// Flight represents a single Flight with its price, dates and related
// information
type Flight struct {
	Price         int
	Depature      time.Time
	Arrival       time.Time
	ReturnArrival time.Time // TODO: Make a "Trip" that includes both back and forth flights
	From          FlightLocation
	To            FlightLocation
}

// FlightLocation represents a departure or destination location of
// a flight
type FlightLocation struct {
	Airport string
}

func (flight Flight) String() string {
	return fmt.Sprintf(
		"Flight for $%d, departs from %s on %s and arrives at %s on %s",
		flight.Price,
		flight.From.Airport,
		flight.Depature,
		flight.To.Airport,
		flight.Arrival,
	)
}
