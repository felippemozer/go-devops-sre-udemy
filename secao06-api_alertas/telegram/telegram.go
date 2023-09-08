package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text      string `json:"text"`
	ChannelId int64  `json:"channel_id"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func SendTelegram(res http.ResponseWriter, req *http.Request) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Could not retrieve TELEGRAM_BOT_TOKEN environment variable")
	}

	var errorMessage ErrorMessage
	message := Message{}
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		errorMessage.Error = fmt.Sprintf("Request body decode error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot creation error: %v", err)
		errorMessage.Error = "Bot creation error"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	alertText := tgbotapi.NewMessage(message.ChannelId, message.Text)

	botMessage, err := bot.Send(alertText)
	if err != nil {
		log.Fatalf("Send message error: %v", err)
		errorMessage.Error = "Send Message error"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	log.Printf("Message sent: %d", botMessage.MessageID)
}
