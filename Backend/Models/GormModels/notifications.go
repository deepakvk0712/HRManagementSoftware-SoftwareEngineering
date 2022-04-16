package gormModels

type Notification struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
	Read     bool   `json:"read"`
}
