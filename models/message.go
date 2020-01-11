package models

// Message -> communication with client
type Message struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
