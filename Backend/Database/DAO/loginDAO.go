package Dao

import (
	"fmt"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"time"
)

func CreateLoginDAO(email, password string) int {
	loginUser := gormModels.LoginUser{
		Email:     email,
		Password:  password,
		CreatedTS: time.Now(),
		UpdatedTS: time.Now(),
	}

	result := utils.Db.Create(&loginUser)

	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}

func DeleteLoginDAO(email string) int {
	result := utils.Db.Where("EMAIL = ?", email).Delete(gormModels.LoginUser{})

	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}
