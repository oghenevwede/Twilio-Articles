package models

//represents the type of error
type ErrorType string

const (
	UnknownError    ErrorType = "UnknownError"
	DeveloperError  ErrorType = "DeveloperError"
	NetworkError    ErrorType = "NetworkError"
	DatabaseError   ErrorType = "DatabaseError"
	ValidationError ErrorType = "ValidationError"
)

//represents the severity of the error
type Severity string

const (
	Info     Severity = "Info"
	Warning  Severity = "Warning"
	Critical Severity = "Critical"
)

//defines the structure of an error report
type ErrorReport struct {
	ErrorType   ErrorType      `json:"error_type"`
	Message     string         `json:"message"`
	Application string         `json:"application"`
	Severity    Severity       `json:"severity"`
	Timestamp   string         `json:"timestamp"`
	Context     RequestContext `json:"context"`
	Notify      []string       `json:"notify"`
}

//captures additional context about the error
type RequestContext struct {
	IP        string            `json:"ip"`
	UserAgent string            `json:"user_agent"`
	URL       string            `json:"url"`
	Headers   map[string]string `json:"headers"`
}
