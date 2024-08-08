package notifications

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/dayroromero/storiChallenge/pkg/models"
	emailrepository "github.com/dayroromero/storiChallenge/pkg/notifications/repository"
	"github.com/dayroromero/storiChallenge/utils"
)

//go:embed templates/summary_account_email.html
var emailTemplate string

type EmailNotification struct {
	RecipientEmail string
	Subject        string
}

// RenderEmailBody renders the body of the email using the template and data provided.
func RenderEmailBody(data models.EmailData) (string, error) {
	emailTemplate, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Printf("error loading email template: %v", err)
		return "", fmt.Errorf("error loading email template: %v", err)
	}

	var renderedBody bytes.Buffer
	if err := emailTemplate.Execute(&renderedBody, data); err != nil {
		log.Printf("error rendering email template: %v", err)
		return "", fmt.Errorf("error rendering email template: %v", err)
	}

	return renderedBody.String(), nil
}

// SendEmail sends an email using the rendered body and data provided.
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

// OrchestrateEmailSending orchestrates templating and email sending
func OrchestrateEmailSending(notification EmailNotification) error {
	data := emailrepository.GetSummary(1)

	body, err := RenderEmailBody(data)
	if err != nil {
		return fmt.Errorf("error al renderizar el cuerpo del correo electrónico: %v", err)
	}

	if err := SendEmail(notification, body); err != nil {
		return fmt.Errorf("Error sendig email: %v", err)
	}

	return nil
}

func SendSummary(UserID int) {
	user, err := emailrepository.GetUser(UserID)
	if err != nil {
		log.Printf("Error getting email from user: %v", err)
		return
	}

	email := EmailNotification{
		RecipientEmail: user.Email,
		Subject:        "Stori - Transactions Summary",
	}

	err = OrchestrateEmailSending(email)
	if err != nil {
		log.Printf("Error email ochestation: %v", err)
	}
}
