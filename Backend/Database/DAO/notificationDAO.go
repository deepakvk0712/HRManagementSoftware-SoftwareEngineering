package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func StoreMessage(message models.NotificationMessage, sender string) int {
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
