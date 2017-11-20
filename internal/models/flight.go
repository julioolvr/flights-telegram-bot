package models

import (
	"fmt"
)

// Flight represents a single Flight with its price, dates and related
// information
type Flight struct {
	Price int
}

func (flight Flight) String() string {
	return fmt.Sprintf("Flight for $%d", flight.Price)
}
