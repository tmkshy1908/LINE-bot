package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Schedule struct {
	Id       int
	Day      int
	Contents string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "usr= dbname= password=")
	if err != nil {
		panic(err)
	}
}

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
					schedule := Schedule{Day: 20220101, Contents: "テスト"}
					schedule.Create()
					replyMessage = "作成しました"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(replyMessage, "更新") {
					replyMessage := "更新します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(replyMessage, "削除") {
					replyMessage := "削除します"
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					schedule := Schedule{Id: 1}
					schedule.Delete()
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

func (schedule *Schedule) Create() (err error) {
	err = Db.QueryRow("insert into schedule (day, contents) values ($1, $2) returning id", schedule.Day, schedule.Contents).Scan(&schedule.Id)
	return
}

func (schedule *Schedule) Delete() (err error) {
	_, err = Db.Exec("delete from schedule where id = $1", schedule.Id)
	return
}
