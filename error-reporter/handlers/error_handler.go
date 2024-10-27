package handlers

import (
	"encoding/json"
	"error-reporter/models"
	"error-reporter/services"
	"error-reporter/utils"
	"net/http"
)

// HandleErrorReport handles incoming error reports
func HandleErrorReport(w http.ResponseWriter, r *http.Request) {
	var report models.ErrorReport

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Log the error report
	utils.LogError(report)

	// Send alert via email
	services.SendAlert(report)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Error reported successfully"})
}
