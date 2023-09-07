package main

import (
	"alertmanager/slack"
)

func main() {
	// email.SendEmail(
	// 	[]string{
	// 		"felippemozer22@gmail.com",
	// 	},
	// 	"Alerta: Servidor Caiu!",
	// 	email.Body{
	// 		Server:   "Google",
	// 		Error:    "Erro ao conectar o servidor",
	// 		Datetime: "06/03/2023 20:53",
	// 	},
	// 	"./email/template.html",
	// )

	slack.SendSlack("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/09/2023 21:35", ".env")
}
