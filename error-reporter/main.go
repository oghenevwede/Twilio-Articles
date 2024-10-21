package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ErrorReport struct {
	ErrorType   string   `json:"error_type"`
	Message     string   `json:"message"`
	Application string   `json:"application"`
	Severity    string   `json:"severity"`
	Notify      []string `json:"notify"` // Emails to notify
}

// API Handler to receive error reports
func reportError(w http.ResponseWriter, r *http.Request) {
	var report ErrorReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Log received error
	log.Printf("Received error report: %+v\n", report)

	// Send alert based on error report
	sendAlert(report)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Error reported successfully"))
}

// SendGrid Integration for sending alert emails
func sendAlert(report ErrorReport) {
	from := mail.NewEmail("Error Reporter", "no-reply@yourdomain.com")
	subject := fmt.Sprintf("Error Report: %s in %s", report.ErrorType, report.Application)
	plainTextContent := fmt.Sprintf("Application: %s\nSeverity: %s\nMessage: %s",
		report.Application, report.Severity, report.Message)

	// Sending email to each recipient specified in the report
	for _, recipient := range report.Notify {
		to := mail.NewEmail("Developer", recipient)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

		// Replace this line with your actual SendGrid API key
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", recipient, err)
		} else {
			log.Printf("Alert email sent to %s: Status Code %d", recipient, response.StatusCode)
		}
	}
}

// Custom email routing based on error type (Optional feature)
// func routeAlert(report ErrorReport) []string {
// 	switch report.ErrorType {
// 	case "database":
// 		return []string{"db-team@yourdomain.com"}
// 	case "network":
// 		return []string{"network-team@yourdomain.com"}
// 	case "authentication":
// 		return []string{"security-team@yourdomain.com"}
// 	default:
// 		return []string{"general@yourdomain.com"}
// 	}
// }

func main() {
	// Add this line to load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/report-error", reportError)

	fmt.Println("Error reporting service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
