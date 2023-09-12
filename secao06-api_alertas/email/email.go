package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type Email struct {
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Datetime string   `json:"datetime"`
	Server   string   `json:"server"`
	Error    string   `json:"error"`
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

func SendEmail(res http.ResponseWriter, req *http.Request) {
	var errorMessage ErrorMessage
	from := verifyEnvVar("EMAIL_FROM", res)
	host := verifyEnvVar("EMAIL_HOST", res)
	pass := verifyEnvVar("EMAIL_PASSWORD", res)
	port := verifyEnvVar("EMAIL_PORT", res)
	templatePath := verifyEnvVar("EMAIL_TEMPLATE_PATH", res)

	email := Email{}
	err := json.NewDecoder(req.Body).Decode(&email)

	if err != nil {
		errorMessage.Error = fmt.Sprintf("Request body decode error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	auth := smtp.PlainAuth("", from, pass, host)

	t, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Fatalf("Error on template parsing: %v", err)
		return
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", email.Subject, mimeHeaders)))
	t.Execute(&body, email)

	err = smtp.SendMail(fmt.Sprintf("%s:%s", host, port), auth, from, email.To, body.Bytes())

	if err != nil {
		errorMessage.Error = fmt.Sprintf("Send email error: %v", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}

	log.Println("Email enviado com sucesso!")
}

// func SendEmail(to []string, subject string, bodyStruct Body, templatePath string) error {
// 	from := "felippemozer22@gmail.com"
// 	pass := os.Getenv("FELIPPE_PASSWORD")

// 	if pass == "" {
// 		return fmt.Errorf("variável de ambiente FELIPPE_PASSWORD não foi configurada")
// 	}

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth("", from, pass, smtpHost)

// 	t, _ := template.ParseFiles(templatePath)

// 	var body bytes.Buffer

// 	mimeHeaders := "MIME-version: 1.0;\nContent-type: text/html; charset=\"UTF-8\";\n\n"

// 	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))
// 	t.Execute(&body, bodyStruct)

// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

// 	if err != nil {
// 		return fmt.Errorf("erro ao enviar o email: %s", err)
// 	}

// 	fmt.Println("Email enviado com sucesso!")
// 	return nil
// }
