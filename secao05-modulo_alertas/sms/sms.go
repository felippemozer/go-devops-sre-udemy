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

func SendSMS(message string, phone string) {

	apiKey := os.Getenv("VONAGE_API_KEY")
	if apiKey == "" {
		panic("VONAGE_API_KEY não foi configurado")
	}

	apiSecret := os.Getenv("VONAGE_API_SECRET")
	if apiSecret == "" {
		panic("VONAGE_API_SECRET não foi configurado")
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
		panic(fmt.Sprintf("Request not allowed. See error below\n%s", err))
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(fmt.Sprintf("Error on URL invoke. See error below\n%s", err))
	}
	defer res.Body.Close()

	log.Println(res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Sprintf("Error when read response body. See error below\n%s", err))
	}

	fmt.Println(string(body))

}
