package models

type MailTemplate struct {
	From        string `json:"from"`
	To          string `json:"to"`
	FromEmail   string `json:"fromEmail"`
	ToEmail     string `json:"toEmail"`
	Subject     string `json:"subject"`
	TextContent string `json:"textContent"`
	HTMLContent string `json:"HTMLContent"`
}
