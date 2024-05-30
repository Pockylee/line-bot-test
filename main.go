package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {
  bot, err := linebot.New(
    os.Getenv("LINE_CHANNEL_SECRET"),
    os.Getenv("LINE_CHANNEL_TOKEN"),
  )
  if err != nil {
    log.Fatal(err)
  }

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
          if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
            log.Print(err)
          }
        }
      }
    }
  })

  log.Println("Server is running at :8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}