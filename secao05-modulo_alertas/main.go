package main

import (
	"alertmanager/telegram"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// felippeEmail := os.Getenv("FELIPPE_EMAIL")
	// if felippeEmail == "" {
	// 	panic("FELIPPE_EMAIL não foi configurado")
	// }
	// email.SendEmail(
	// 	[]string{
	// 		felippeEmail,
	// 	},
	// 	"Alerta: Servidor Caiu!",
	// 	email.Body{
	// 		Server:   "Google",
	// 		Error:    "Erro ao conectar o servidor",
	// 		Datetime: "06/03/2023 20:53",
	// 	},
	// 	"./email/template.html",
	// )

	// slack.SendSlack("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/09/2023 21:35")

	// felippePhone := os.Getenv("FELIPPE_PHONE")
	// if felippePhone == "" {
	// 	panic("FELIPPE_PHONE não foi configurado")
	// }
	// sms.SendSMS("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 07/09/2023 15:27", felippePhone)

	telegramApiKey := os.Getenv("TELEGRAM_API_KEY")
	if telegramApiKey == "" {
		panic("TELEGRAM_API_KEY não foi configurado")
	}
	telegram.SendTelegram(telegramApiKey, "Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 07/09/2023 16:39")
}
