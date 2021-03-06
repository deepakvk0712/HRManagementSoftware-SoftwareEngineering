package Dao

import (
	"fmt"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func InsertFeedback(fb gormModels.Feedback) int {
	result := utils.Db.Create(&fb)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	EID := fb.EmployeeID
	utils.Db.Model(&gormModels.User{}).Where("EMPLOYEE_ID = ?", EID).Update("is_resigned", 1)

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
