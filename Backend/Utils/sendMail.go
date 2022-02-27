package utils

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	"log"
	"os"
)

func SendMail(mailContent models.MailTemplate) int {
	from := mail.NewEmail(mailContent.From, mailContent.FromEmail)
	subject := mailContent.Subject
	to := mail.NewEmail(mailContent.To, mailContent.ToEmail)
	var plainTextContent string
	var htmlContent string
	if mailContent.TextContent != "" {
		plainTextContent = mailContent.TextContent
	}
	if mailContent.HTMLContent != "" {
		htmlContent = mailContent.HTMLContent
	}
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)

		return 0
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return 1
}
