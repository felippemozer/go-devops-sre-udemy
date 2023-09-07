package main

import (
	"alertmanager/sms"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// email.SendEmail(
	// 	[]string{
	// 		os.Getenv("FELIPPE_EMAIL"),
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

	sms.SendSMS("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/09/2023 21:35", os.Getenv("FELIPPE_PHONE"))
}
