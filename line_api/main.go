package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	http.HandleFunc("/calback", lineHandler)
	fmt.Println("http://localhost:8080 で起動中...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		"チャンネルシークレット",
		"アクセストークン",
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
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
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}

			case *linebot.StickerMessage:
				replyMessage := "良いスタンプだねぇ"
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}

			case *linebot.ImageMessage:
				replyMessage := "素敵な写真だなぁ"
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}
