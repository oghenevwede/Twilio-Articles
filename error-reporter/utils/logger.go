package utils

import (
	"error-reporter/models"
	"log"
	"os"
)

// LogError logs the error report to a file
func LogError(report models.ErrorReport) {
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Failed to open log file")
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("[ERROR] %s | App: %s | Severity: %s | Message: %s\n", report.Timestamp, report.Application, report.Severity, report.Message)
}
