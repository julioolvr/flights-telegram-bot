package flight

import (
	"fmt"
	"time"
)

// AirportCode is an alias for a string that represents an airport code
// (e.g. JFK)
type AirportCode string

// A Route for a flight is a collection of segments that goes from one
// location to another (and potentially back to the origin)
type Route struct {
	From     Location
	To       Location
	Segments []RouteSegment
}

// A RouteSegment is each individual flight that forms part of a Route.
type RouteSegment struct {
	From      Location
	To        Location
	Departure time.Time
	Arrival   time.Time
}

// Flight represents a single Flight with its price, dates and related
// information
type Flight struct {
	Price  int
	Routes []Route
}

// Location represents an airport visited in some route during the flight
type Location struct {
	Airport AirportCode
}

// DepartsFrom returns the Location from which the first flight departs from.
func (flight Flight) DepartsFrom() Location {
	return flight.Routes[0].From
}

// Destination returns the Location where the trip finishes (or the point at
// which the round trip is done).
func (flight Flight) Destination() Location {
	return flight.Routes[0].To
}

// DepartsAt returns the Time at which the first flight departs.
func (flight Flight) DepartsAt() time.Time {
	return flight.Routes[0].Segments[0].Departure
}

// ReturnsAt returns the Time at which the last flight arrives.
func (flight Flight) ReturnsAt() time.Time {
	numberOfRoutes := len(flight.Routes)
	lastRoute := flight.Routes[numberOfRoutes-1]
	segmentsInRoute := len(lastRoute.Segments)
	return lastRoute.Segments[segmentsInRoute-1].Arrival
}

func (flight Flight) String() string {
	return fmt.Sprintf(
		"Flight for $%d, departs from %s on %s to %s and returns on %s",
		flight.Price,
		flight.DepartsFrom().Airport,
		flight.DepartsAt(),
		flight.Destination().Airport,
		flight.ReturnsAt(),
	)
}
