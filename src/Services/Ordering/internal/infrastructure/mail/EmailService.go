package mail

import (
	"fmt"
	"log"
)

type EmailService struct {
	emailSettings *EmailSettings
	logger        *log.Logger
}

func NewEmailService(mailSettings *EmailSettings, logger *log.Logger) *EmailService {
	return &EmailService{
		emailSettings: mailSettings,
		logger:        logger,
	}
}

func (svc *EmailService) SendEmail(email *Email) error {
	client := sendgrid.NewSendClient(svc.emailSettings.ApiKey)

	subject := email.Subject
	to := mail.NewEmail("", email.To)
	emailBody := email.Body

	from := mail.NewEmail(svc.emailSettings.FromName, svc.emailSettings.FromAddress)

	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	response, err := client.Send(message)

	if err != nil {
		svc.logger.Errorf("Error sending email: %s", err.Error())
		return err
	}

	svc.logger.Info("Email sent.")

	if response.StatusCode == 202 || response.StatusCode == 200 {
		return nil
	}

	svc.logger.Error("Email sending failed.")
	return fmt.Errorf("failed to send email: %v", response)
}
