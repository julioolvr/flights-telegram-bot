package main

import (
	"fmt"
	"os"

	"github.com/julioolvr/flights-telegram-bot/internal/api"
)

type flight struct {
	price int
}

func main() {
	res, err := api.FindFlights(api.QueryParams{FlyFrom: "JFK", FlyTo: "LAX"})

	if err == nil {
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error with request to API %s\n", err)
	}
}
