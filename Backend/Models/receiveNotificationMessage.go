package models

type ReceiveNotificationMessage struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}
