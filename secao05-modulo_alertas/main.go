package main

import (
	"alertmanager/email"
	"alertmanager/slack"
	"alertmanager/sms"
	"alertmanager/telegram"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	felippeEmail := os.Getenv("FELIPPE_EMAIL")
	if felippeEmail != "" {
		log.Printf("Enviando email para %s\n", felippeEmail)
		err := email.SendEmail(
			[]string{
				felippeEmail,
			},
			"Alerta: Servidor Caiu!",
			email.Body{
				Server:   "Google",
				Error:    "Erro ao conectar o servidor",
				Datetime: "06/03/2023 20:53",
			},
			"./email/template.html",
		)
		if err != nil {
			log.Printf("Não foi possível enviar o email para %s\n", felippeEmail)
			fmt.Println(err)
		}
		fmt.Println("Continuando execução...")
	} else {
		log.Println("Variável de ambiente FELIPPE_EMAIL não foi configurada. Pulando execução de envio de email...")
	}

	log.Println("Enviando Slack para canal Alertas")
	err := slack.SendSlack("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/09/2023 21:35")
	if err != nil {
		log.Printf("Não foi possível enviar o Slack para o canal Alertas\n")
		fmt.Println(err)
	}
	fmt.Println("Continuando execução...")

	felippePhone := os.Getenv("FELIPPE_PHONE")
	if felippePhone != "" {
		err := sms.SendSMS("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 07/09/2023 15:27", felippePhone)
		if err != nil {
			log.Printf("Não foi possível enviar o SMS para o telefone %s\n", felippePhone)
			fmt.Println(err)
		}
		fmt.Println("Continuando execução...")
	} else {
		log.Println("Variável de ambiente FELIPPE_PHONE não foi configurada. Pulando execução de envio de SMS...")
	}

	telegramApiKey := os.Getenv("TELEGRAM_API_KEY")
	if telegramApiKey != "" {
		err = telegram.SendTelegram(telegramApiKey, "Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 07/09/2023 16:39")
		if err != nil {
			log.Printf("Não foi possível enviar o Telegram para o canal Alertas\n")
			log.Fatal(err)
		}
	} else {
		log.Fatal("Variável de ambiente TELEGRAM_API_KEY não foi configurada. Pulando execução de envio de Telegram...")
	}
}
