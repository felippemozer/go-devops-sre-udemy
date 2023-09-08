package main

import (
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
	log.Fatal(http.ListenAndServe(":9090", router))

	// felippeEmail := os.Getenv("FELIPPE_EMAIL")
	// if felippeEmail != "" {
	// 	log.Printf("Enviando email para %s\n", felippeEmail)
	// 	err := email.SendEmail(
	// 		[]string{
	// 			felippeEmail,
	// 		},
	// 		"Alerta: Servidor Caiu!",
	// 		email.Body{
	// 			Server:   "Google",
	// 			Error:    "Erro ao conectar o servidor",
	// 			Datetime: "06/03/2023 20:53",
	// 		},
	// 		"./email/template.html",
	// 	)
	// 	if err != nil {
	// 		log.Printf("Não foi possível enviar o email para %s\n", felippeEmail)
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("Continuando execução...")
	// } else {
	// 	log.Println("Variável de ambiente FELIPPE_EMAIL não foi configurada. Pulando execução de envio de email...")
	// }

	// log.Println("Enviando Slack para canal Alertas")
	// err := slack.SendSlack("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/09/2023 21:35")
	// if err != nil {
	// 	log.Printf("Não foi possível enviar o Slack para o canal Alertas\n")
	// 	fmt.Println(err)
	// }
	// fmt.Println("Continuando execução...")

	// felippePhone := os.Getenv("FELIPPE_PHONE")
	// if felippePhone != "" {
	// 	err := sms.SendSMS("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 07/09/2023 15:27", felippePhone)
	// 	if err != nil {
	// 		log.Printf("Não foi possível enviar o SMS para o telefone %s\n", felippePhone)
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("Continuando execução...")
	// } else {
	// 	log.Println("Variável de ambiente FELIPPE_PHONE não foi configurada. Pulando execução de envio de SMS...")
	// }

}
