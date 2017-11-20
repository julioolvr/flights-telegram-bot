package models

import (
	"fmt"
	"time"
)

// Flight represents a single Flight with its price, dates and related
// information
type Flight struct {
	Price    int
	Depature time.Time
	Arrival  time.Time
}

func (flight Flight) String() string {
	return fmt.Sprintf(
		"Flight for $%d, departs at %s and arrives at %s",
		flight.Price,
		flight.Depature,
		flight.Arrival,
	)
}
