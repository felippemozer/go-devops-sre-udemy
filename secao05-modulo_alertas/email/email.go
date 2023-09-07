package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type Body struct {
	Server   string
	Error    string
	Datetime string
}

func SendEmail(to []string, subject string, bodyStruct Body, templatePath string) error {
	from := "felippemozer22@gmail.com"
	pass := os.Getenv("FELIPPE_PASSWORD")

	if pass == "" {
		return fmt.Errorf("variável de ambiente FELIPPE_PASSWORD não foi configurada")
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-type: text/html; charset=\"UTF-8\";\n\n"

	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, bodyStruct)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

	if err != nil {
		return fmt.Errorf("erro ao enviar o email: %s", err)
	}

	fmt.Println("Email enviado com sucesso!")
	return nil
}
