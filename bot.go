package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/julioolvr/flights-telegram-bot/internal/services/flightService"
	godotenv "gopkg.in/joho/godotenv.v1"
	"gopkg.in/robfig/cron.v2"
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

	flyFrom := flag.String("from", "", "Fly from")
	flyTo := flag.String("to", "", "Fly to")
	dateFrom := flag.String("leavingFrom", "", "Leaving starting on date")
	dateTo := flag.String("leavingTo", "", "Leaving ending on date")
	limit := flag.Int("limit", 5, "Number of results")
	chatID := flag.Int("chat", 0, "Telegram chat id")

	flag.Parse()

	c := cron.New()

	c.AddFunc("TZ=America/Argentina/Buenos_Aires 0 11 * * *", func() {
		res, err := flightService.
			Find().
			From(*flyFrom).
			To(*flyTo).
			DepartureDateRange(*dateFrom, *dateTo).
			DaysInDestination(10, 15).
			Limit(*limit).
			Exec()

		if err != nil {
			log.Fatalln(err)
		}

		var message bytes.Buffer

		for _, flight := range res {
			message.WriteString(fmt.Sprintf(
				"💸 $%d 📅 %s - %s\n🛫 %s - 🛬 %s\n\n",
				flight.Price,
				flight.DepartsAt().Format("2006-01-02"),
				flight.ReturnsAt().Format("2006-01-02"),
				flight.DepartsFrom().Airport,
				flight.Destination().Airport,
			))
		}

		bot.SendMessage(tb.Chat{ID: int64(*chatID)}, message.String(), nil)
	})

	c.Start()

	// TODO: Is this a good way to leave a program running?
	select {}
}
