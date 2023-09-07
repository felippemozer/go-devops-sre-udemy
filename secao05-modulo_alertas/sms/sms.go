package sms

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func SendSMS(message string, phone string) error {

	apiKey := os.Getenv("VONAGE_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("variável de ambiente VONAGE_API_KEY não foi configurada")
	}

	apiSecret := os.Getenv("VONAGE_API_SECRET")
	if apiSecret == "" {
		return fmt.Errorf("variável de ambiente VONAGE_API_SECRET não foi configurada")
	}

	data := url.Values{}
	data.Set("api_key", apiKey)
	data.Set("api_secret", apiSecret)
	data.Set("from", "Vonage APIs")
	data.Set("text", message)
	data.Set("to", phone)
	fmt.Println(data.Encode())

	endpoint := "https://rest.nexmo.com/sms/json"
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("requisição não permitida: %s", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("erro de envio de requisição: %s", err)
	}
	defer res.Body.Close()

	log.Println(res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("erro ao ler corpo de resposta: %s", err)
	}

	fmt.Println(string(body))
	return nil
}
