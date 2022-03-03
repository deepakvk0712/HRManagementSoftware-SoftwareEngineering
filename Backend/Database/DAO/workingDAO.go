package Dao

import (
	"fmt"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func InsertWorkingRecord(workinghours gormModels.WorkingHours) int {
	result := utils.Db.Create(&workinghours)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
