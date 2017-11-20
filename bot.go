package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/julioolvr/flights-telegram-bot/internal/api"
	godotenv "gopkg.in/joho/godotenv.v1"
	tb "gopkg.in/tucnak/telebot.v1"
)

type flight struct {
	price int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	bot, err := tb.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	res, err := api.FindFlights(api.QueryParams{
		FlyFrom:               "JFK",
		FlyTo:                 "36.1699--115.1398-1000km",
		DateFrom:              "01/01/2018",
		DateTo:                "10/01/2018",
		DaysInDestinationFrom: 10,
		DaysInDestinationTo:   15,
		Limit:                 5,
	})

	if err != nil {
		log.Fatalln(err)
	}

	var message bytes.Buffer

	for _, flight := range res {
		message.WriteString(fmt.Sprintf(
			"$%d %s ✈️ %s / 🛫 %s - 🛬 %s\n",
			flight.Price,
			flight.From.Airport,
			flight.To.Airport,
			flight.Depature.Format("2006-01-02"),
			flight.ReturnArrival.Format("2006-01-02"),
		))
	}

	chatID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	bot.SendMessage(tb.Chat{ID: chatID}, message.String(), nil)
}
