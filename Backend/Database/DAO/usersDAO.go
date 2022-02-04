package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
)

func CreateEmployeeDao(employee models.Employee, EID int) {

	result := db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	fmt.Println(result.Error, result.RowsAffected)
}
