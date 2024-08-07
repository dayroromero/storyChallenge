package notifications

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/dayroromero/storiChallenge/utils"
)

// EmailNotification representa un correo electrónico.
type EmailNotification struct {
	RecipientEmail string
	Subject        string
}

// RenderEmailBody renderiza el cuerpo del correo electrónico utilizando la plantilla y los datos proporcionados.
func RenderEmailBody(data EmailData) (string, error) {
	emailTemplate, err := template.ParseFiles("pkg/notifications/templates/summary_account_email.html")
	if err != nil {
		log.Printf("error loading email template: %v", err)
		return "", fmt.Errorf("error loading email template: %v", err)
	}

	var renderedBody bytes.Buffer
	if err := emailTemplate.Execute(&renderedBody, data); err != nil {
		log.Printf("error reendering email template: %v", err)
		return "", fmt.Errorf("error reendering email template: %v", err)
	}

	return renderedBody.String(), nil
}

// SendEmail envía un correo electrónico utilizando el cuerpo renderizado y los datos proporcionados.
func SendEmail(notification EmailNotification, body string) error {
	smtpHost := utils.GetEnvVar("SMTP_HOST")
	smtpPort := utils.GetEnvVar("SMTP_PORT")
	senderEmail := utils.GetEnvVar("SENDER_EMAIL")
	senderPassword := utils.GetEnvVar("SENDER_PASSWORD")

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	to := []string{notification.RecipientEmail}
	msg := []byte("To: " + notification.RecipientEmail + "\r\n" +
		"Subject: " + notification.Subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return fmt.Errorf("error sending email: %v", err)
	}

	log.Println("Email sent successfully to", notification.RecipientEmail)
	return nil
}

func OrchestrateEmailSending(notification EmailNotification) error {
	data := EmailData{
		ClientName:           "Nombre del Cliente",
		TotalBalance:         39.74,
		TransactionsInJuly:   2,
		TransactionsInAugust: 2,
		AverageDebitAmount:   -15.38,
		AverageCreditAmount:  35.25,
	}

	body, err := RenderEmailBody(data)
	if err != nil {
		return fmt.Errorf("error al renderizar el cuerpo del correo electrónico: %v", err)
	}

	if err := SendEmail(notification, body); err != nil {
		return fmt.Errorf("error al enviar el correo electrónico: %v", err)
	}

	return nil
}
