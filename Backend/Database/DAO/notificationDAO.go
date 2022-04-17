package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strings"
)

func StoreMessage(message models.ReceiveNotificationMessage, sender string) int {
	notification := gormModels.Notification{
		Sender:   sender,
		Receiver: message.Receiver,
		Message:  message.Message,
		Read:     false,
	}

	result := utils.Db.Create(&notification)
	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}

func MarkAsRead(id string) int {
	utils.Db.Exec("UPDATE NOTIFICATIONS SET READ = TRUE WHERE MESSAGE_ID = ?", id)

	return 1
}

func GetMessages(email string) ([]models.SendNotificationMessage, int) {
	var messages []models.SendNotificationMessage

	rows, _ := utils.Db.Raw("SELECT SENDER, MESSAGE, MESSAGE_ID FROM NOTIFICATIONS WHERE LOWER(RECEIVER) = ? AND READ = FALSE", strings.ToLower(email)).Rows()
	defer rows.Close()

	message := models.SendNotificationMessage{}

	for rows.Next() {
		rows.Scan(&message.Sender, &message.Message, &message.MessageID)

		messages = append(messages, message)
	}

	return messages, 1
}
