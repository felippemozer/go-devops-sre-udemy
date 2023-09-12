package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

type Message struct {
	Color    string `json:"color"`
	Pretext  string `json:"pretext"`
	Text     string `json:"text"`
	Error    string `json:"error"`
	Datetime string `json:"datetime"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func verifyEnvVar(envName string, res http.ResponseWriter) string {
	var errorMessage ErrorMessage
	result := os.Getenv(envName)

	if result == "" {
		errorMessage.Error = fmt.Sprintf("Could not retrieve %s environment variable", envName)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errorMessage)
		return ""
	}

	return result
}

func SendSlack(res http.ResponseWriter, req *http.Request) {
	var errorMessage ErrorMessage
	token := verifyEnvVar("SLACK_AUTH_TOKEN", res)
	channelID := verifyEnvVar("SLACK_CHANNEL_ID", res)
	client := slack.New(token, slack.OptionDebug(true))

	message := Message{}

	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		errorMessage.Error = fmt.Sprintf("Request body decode error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	attachment := slack.Attachment{
		Color:   message.Color,
		Pretext: message.Pretext,
		Text:    fmt.Sprintf("%s\nError: %s\nDatetime: %s", message.Text, message.Error, message.Datetime),
	}
	_, _, err = client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		errorMessage.Error = fmt.Sprintf("Send message error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}
	log.Println("Slack enviado com sucesso")
}
