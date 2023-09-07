package telegram

import (
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text      string `json:"text"`
	ChannelId int64  `json:"channel_id"`
}

func SendTelegram(apiKey string, textMessage string) error {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return fmt.Errorf("erro ao criar novo bot: %s", err)
	}
	message := Message{}
	message.Text = textMessage

	channelId := os.Getenv("TELEGRAM_CHANNEL_ID")
	if channelId == "" {
		return fmt.Errorf("variável de ambiente TELEGRAM_CHANNEL_ID não foi configurada")
	}

	message.ChannelId, err = strconv.ParseInt(channelId, 10, 64)
	if err != nil {
		return fmt.Errorf("não foi possível converter ID do canal em int64: %s", err)
	}
	bot.Debug = true
	alertText := tgbotapi.NewMessage(message.ChannelId, message.Text)
	_, err = bot.Send(alertText)
	if err != nil {
		return fmt.Errorf("erro ao enviar texto para ao canal")
	}

	return nil
}
