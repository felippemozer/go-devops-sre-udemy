package sms

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Message struct {
	Text  string `json:"text"`
	Phone string `json:"phone"`
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

func SendSMS(res http.ResponseWriter, req *http.Request) {
	var errorMessage ErrorMessage
	apiKey := verifyEnvVar("VONAGE_API_KEY", res)
	apiSecret := verifyEnvVar("VONAGE_API_SECRET", res)

	message := Message{}
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		errorMessage.Error = fmt.Sprintf("Request body decode error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	data := url.Values{}
	data.Set("api_key", apiKey)
	data.Set("api_secret", apiSecret)
	data.Set("from", "Vonage APIs")
	data.Set("text", message.Text)
	data.Set("to", message.Phone)

	endpoint := "https://rest.nexmo.com/sms/json"
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		errorMessage.Error = fmt.Sprintf("Request not allowed: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		errorMessage.Error = fmt.Sprintf("Send message error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	log.Println("SMS sent successfully")

}
