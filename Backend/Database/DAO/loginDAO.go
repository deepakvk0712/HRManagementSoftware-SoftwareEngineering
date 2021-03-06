package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strings"
	"time"
)

func CreateLoginDAO(email, password string, role byte) int {
	loginUser := gormModels.LoginUser{
		Email:      email,
		Password:   password,
		CreatedTS:  time.Now(),
		UpdatedTS:  time.Now(),
		Role:       role,
		FirstLogin: true,
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

func GetUserDetailsDAO(email string) (models.UserDetails, int) {
	userDetails := models.UserDetails{}

	isResigned := false

	row := utils.Db.Raw("SELECT IS_RESIGNED FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return userDetails, 0
	}

	row.Scan(&isResigned)

	if isResigned {
		return userDetails, 1
	}

	row = utils.Db.Raw("SELECT EMAIL, PASSWORD, ROLE, FIRST_LOGIN FROM LOGIN_USERS WHERE LOWER(EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return userDetails, 0
	}

	row.Scan(&userDetails.Email, &userDetails.Password, &userDetails.Role, &userDetails.FirstLogin)
	return userDetails, 1
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

func GetEmailDAO(email string) (string, int) {
	var count int
	row := utils.Db.Raw("SELECT COUNT(1) FROM LOGIN_USERS WHERE LOWER(EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		return "", 0
	}

	row.Scan(&count)

	if count == 0 {
		return "", 0
	}

	return email, 1
}

func UpdatePasswordDAO(email, password string) int {
	row := utils.Db.Model(&gormModels.LoginUser{}).Where("EMAIL = ?", email).Update("PASSWORD", password)
	if row.Error != nil {
		return 0
	}

	return 1
}

func UpdateFirstLoginDAO(email string) int {
	row := utils.Db.Model(&gormModels.LoginUser{}).Where("EMAIL = ?", email).Update("FIRST_LOGIN", false)
	if row.Error != nil {
		return 0
	}

	return 1
}
