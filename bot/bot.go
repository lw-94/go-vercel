package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botListener(token string) {
	bot, err := tgBotApi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
func Handler(w http.ResponseWriter, r *http.Request) {
	g := gin.Default()

	token := "6706237172:AAFZyrXsYjMg2ion8MH2rG99Pf-Cjf-DjVw"
	g.Any("/*", func(c *gin.Context) {
		fmt.Println(c)
		botListener(token)
	})

	g.ServeHTTP(w, r)
}
