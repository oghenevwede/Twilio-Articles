package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"error-reporter/models"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// loads email configuration from config.json
func LoadConfig() (map[string]string, string, string) {
	var config struct {
		TeamEmails     map[string]string `json:"team_emails"`
		DefaultEmail   string            `json:"default_email"`
		SendGridAPIKey string            `json:"sendgrid_api_key"`
	}

	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config.TeamEmails, config.DefaultEmail, config.SendGridAPIKey
}

// sends an alert email for the given error report
func SendAlert(report models.ErrorReport) {
	teamEmails, defaultEmail, sendgridAPIKey := LoadConfig()

	// Determine the recipients based on the error report's Notify field
	recipients := report.Notify

	// If no specific recipients are provided, fallback to default email
	if len(recipients) == 0 {
		recipients = []string{defaultEmail} // Fallback to default email
	}

	from := mail.NewEmail("Error Reporter", "test@email.com")
	subject := fmt.Sprintf("Error Report: %s in %s", report.ErrorType, report.Application)
	plainTextContent := fmt.Sprintf("Application: %s\nSeverity: %s\nMessage: %s\nContext: %+v",
		report.Application, report.Severity, report.Message, report.Context)

	// Log the length of the API key for debugging
	log.Printf("SENDGRID_API_KEY length: %d", len(sendgridAPIKey))
	if sendgridAPIKey == "" {
		log.Fatal("SENDGRID_API_KEY is not set")
	}

	// Send an email to each recipient specified in the report
	for _, recipient := range recipients {
		var email string
		// If the recipient is a known team, verify if it's in the teamEmails map
		if knownEmail, exists := teamEmails[recipient]; exists {
			email = knownEmail
		} else {
			log.Printf("Unknown team recipient: %s. Using default email: %s", recipient, defaultEmail)
			email = defaultEmail
		}

		to := mail.NewEmail("Developer", email)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

		client := sendgrid.NewSendClient(sendgridAPIKey)
		response, err := client.Send(message)

		if err != nil {
			log.Printf("Failed to send email to %s: %v", email, err)
		} else {
			log.Printf("Alert email sent to %s: Status Code %d", email, response.StatusCode)
		}
	}
}
