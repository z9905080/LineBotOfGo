package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	secretKey := "bc9a61ad5f96567691d97b50e5bc5bff"
	accessToken := "kyV6iGafS/hD2DOAQcVgfe4NWstXgYDdDcEvLRjEjc2PgqtXlFw0PJvLt5qASQIsttAmkXht9I9mp83GOtI4EYqBpK/IVy2FzlRSGqRxXRl2y3emhYXMNT7fGNRNBRl1kKaAxKMahDcw6f9K2oViswdB04t89/1O/w1cDnyilFU="

	bot, err := linebot.New(secretKey, accessToken)
	if err != nil {
		log.Println("Connect Error:", err)
	}

	// events, requestErr := bot.ParseRequest(req)
	// if requestErr != nil {
	// 	log.Println("Request Error:", requestErr)
	// 	// Do something when something bad happened.
	// }

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Writer.WriteHeader(400)
			} else {
				c.Writer.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					log.Println("ID:", message.ID)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	r.Run()
}
