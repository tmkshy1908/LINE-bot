package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

			//テキストの場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				if strings.Contains(replyMessage, "作成") {
					replyMessage := "作成します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(replyMessage, "更新") {
					replyMessage := "更新します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(replyMessage, "削除") {
					replyMessage := "削除します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(replyMessage, "取得") {
					replyMessage := "取得します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else {
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}

			//スタンプの場合
			case *linebot.StickerMessage:
				replyMessage := "良いスタンプだねぇ"
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}

			//画像の場合
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
