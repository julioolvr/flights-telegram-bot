package main

import (
	"fmt"
	"os"

	"github.com/julioolvr/flights-telegram-bot/api"
)

type flight struct {
	price int
}

func main() {
	res, err := api.FindFlights()

	if err == nil {
		fmt.Printf("Response %v\n", res)
	} else {
		fmt.Fprintf(os.Stderr, "Error with request to API %s\n", err)
	}
}
