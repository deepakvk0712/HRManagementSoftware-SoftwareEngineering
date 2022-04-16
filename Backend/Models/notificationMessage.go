package models

type NotificationMessage struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}
