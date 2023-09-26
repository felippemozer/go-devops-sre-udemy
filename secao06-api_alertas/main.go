package main

import (
	"alertmanager/email"
	"alertmanager/slack"
	"alertmanager/sms"
	"alertmanager/telegram"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	f, err := os.OpenFile("alertmanager.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Starting alert manager")

	router := mux.NewRouter()
	router.HandleFunc("/telegram", telegram.SendTelegram).Methods("POST")
	router.HandleFunc("/email", email.SendEmail).Methods("POST")
	router.HandleFunc("/slack", slack.SendSlack).Methods("POST")
	router.HandleFunc("/sms", sms.SendSMS).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
