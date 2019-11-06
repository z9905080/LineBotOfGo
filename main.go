package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	secretKey := "bc9a61ad5f96567691d97b50e5bc5bff"
	accessToken := "kyV6iGafS/hD2DOAQcVgfe4NWstXgYDdDcEvLRjEjc2PgqtXlFw0PJvLt5qASQIsttAmkXht9I9mp83GOtI4EYqBpK/IVy2FzlRSGqRxXRl2y3emhYXMNT7fGNRNBRl1kKaAxKMahDcw6f9K2oViswdB04t89/1O/w1cDnyilFU="

	client := &http.Client{}
	bot, err := linebot.New(secretKey, accessToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Println("Connect Error:", err)
	}

	// events, requestErr := bot.ParseRequest(req)
	// if requestErr != nil {
	// 	log.Println("Request Error:", requestErr)
	// 	// Do something when something bad happened.
	// }

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					log.Println("Msg:", message.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
