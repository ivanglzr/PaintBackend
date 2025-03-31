package models

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
