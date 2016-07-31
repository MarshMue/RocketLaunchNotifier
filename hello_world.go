package main

import (
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	//"gopkg.in/telegram-bot-api.v4"
	"fmt"
	//"log"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
)

func main() {
	/*
		bot, err := tgbotapi.NewBotAPI("249474117:AAHPf1gj7olfzcEs9WmmMPSW34eHM8FF32s")
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}



		db, err := sql.Open("sqlite3", "./launchbot.db")

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		_, err = db.Exec("insert into launches(Date, Rocket, Payload, Description) values('07-30-2016', 'Falcon 9', 'GOES-R', 'Sample')")

		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("select Date, Rocket, Payload, Description from launches")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var date string
			var rocket string
			var payload string
			var desc string
			err = rows.Scan(&date, &rocket, &payload, &desc)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(date, rocket, payload, desc)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/

	resp, err := http.Get("http://spaceflightnow.com/launch-schedule/")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcherDate := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.Span {
			return scrape.Attr(n, "class") == "launchdate"
		}
		return false
	}

	matcherMission := func(n *html.Node) bool {
		if n.DataAtom == atom.Span {
			return scrape.Attr(n, "class") == "mission"
		}
		return false
	}

	// grab all launch dates
	dates := scrape.FindAll(root, matcherDate)
	for _, date := range dates {
		fmt.Printf("Date: %s\n", scrape.Text(date))
	}
	fmt.Printf("Length of dates: %d\n", len(dates))

	// grab all missions
	missions := scrape.FindAll(root, matcherMission)
	for _, mission := range missions {
		fmt.Printf("Mission: %s\n", scrape.Text(mission))
	}
	fmt.Printf("Length of missions: %d\n", len(missions))

}
