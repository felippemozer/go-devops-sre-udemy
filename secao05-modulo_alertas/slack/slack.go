package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendSlack(alertText string) error {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		return fmt.Errorf("variável de ambiente SLACK_AUTH_TOKEN não foi configurada")
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		return fmt.Errorf("variável de ambiente SLACK_CHANNEL_ID não foi configurada")
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
		return fmt.Errorf("erro ao postar mensagem: %s", err)
	}

	fmt.Println("Mensagem enviada com sucesso para o canal ID", channelID, "às", timestamp)
	return nil
}
