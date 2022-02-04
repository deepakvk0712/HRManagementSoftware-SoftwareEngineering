package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func CreateEmployeeDao(employee models.Employee) {

	result := utils.Db.Create(&employee)
	fmt.Println(result.Error, result.RowsAffected)
}
