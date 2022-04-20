package models

type SendNotificationMessage struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	MessageID int    `json:"messageID"`
}
