package slack

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func SendSlack(alertText string, envPath string) {
	godotenv.Load()

	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("SLACK_AUTH_TOKEN não foi configurado")
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("SLACK_CHANNEL_ID não foi configurado")
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alerta de servidor down",
		Text:    alertText,
	}

	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensagem enviada com sucesso para o canal ID", channelID, "às", timestamp)
}
