package Dao

import (
	"fmt"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strings"
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

func GetPasswordDAO(email string) (string, int) {
	var password string
	row := utils.Db.Raw("SELECT PASSWORD FROM LOGIN_USERS WHERE LOWER(EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		return "", 0
	}

	row.Scan(&password)

	return password, 1
}

func UpdatePasswordDAO(email, password string) int {
	row := utils.Db.Model(&gormModels.LoginUser{}).Where("EMAIL = ?", email).Update("PASSWORD", password)
	if row.Error != nil {
		return 0
	}

	return 1
}
