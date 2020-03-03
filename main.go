package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type AccountST struct {
	Account  string
	Password string
}

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

	dataMap := map[string]AccountST{
		"U3fbaefa437ebd2c4d89c79b898ef5129": {Account: "20190881", Password: "a9906296"},       // Shouting
		"U8ba30807b36213e65214d9c894b10b73": {Account: "20190883", Password: "00000000"},       // Ray
		"U4cf6460eab6a5a60eb19221bdf2a816b": {Account: "20190982", Password: "blue0981944899"}, // 藍
		"U73989a4a17ea8a26a5cc774b7c66347c": {Account: "20190481", Password: "www714556"},      // 貴軒
		"Uc354ea11210fcc33ea09e89240235879": {Account: "20190884", Password: "ab8063352l"},     // 松儒
		"U9bb4d80b44bfcf04c4de3e8bbc7510c4": {Account: "20190882", Password: "qwe123"},         // Shank
		"Ue4a925aed5130d864eb77a4bc8fa1932": {Account: "20190682", Password: "111450"},         // DINDIN
		"Ua6cc33837be8bc6a55b2c3e190e897d1": {Account: "20190885", Password: "2lgidoal"},       // Kevin
		"Udbf9c361db4c8ca4e6c98f52f983c2e1": {Account: "20190781", Password: "hln06012"},       // 企鵝
		"Ua67e25918327a060f2d5a7105a8f8a1d": {Account: "20190981", Password: "j4163010"},       // 翠翠子
		"U8dbfd56afd8b693938b74ecb712bf424": {Account: "20200282", Password: "GARYgary2020"},   // Gary
		"U8f68f3a3dd4958f266f130ecd1e86561": {Account: "20200281", Password: "Max0958828605"},  // 建仲
		"Ub8c9b012af988e6215f2c3b42ce2aada": {Account: "20200381", Password: "a1030410304"},    // 建仲
	}

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

			log.Println("UserID:", event.Source.UserID)

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					log.Println("MsgID:", message.ID)
					resp := ""
					switch message.Text {
					case "#打卡":
						{
							resp = "http://goyu.ddns.net:1688/sign.html"
							if data, isExist := dataMap[event.Source.UserID]; isExist {
								resp += "?account=" + data.Account + "&password=" + data.Password
							}
						}
					case "#拜託給我加入":
						{
							resp = "既然你誠心誠意的發問了，我就給你ID吧\n" +
								"ID: " + event.Source.UserID +
								"\n去找我的徒弟 Shouting 把ID給他"
						}
					default:
						{
							resp = message.Text
						}
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resp)).Do(); err != nil {
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
