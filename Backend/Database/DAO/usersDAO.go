package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateUserDAO(u models.User, personalEmail string) (string, int) {
	var num int
	emailSlice := strings.Split(personalEmail, "@")

	row := utils.Db.Raw("SELECT COUNT(1) FROM USERS WHERE LOWER(PERSONAL_EMAIL) LIKE \"" + strings.ToLower(emailSlice[0]) + "@%\"").Row()
	if row.Err() != nil {
		fmt.Println(row.Err())
		return "", 0
	}

	row.Scan(&num)

	email := emailSlice[0] + strconv.Itoa(num+1) + os.Getenv("EMAIL_DOMAIN")

	user := gormModels.User{
		OfficialEmail: email,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		BusinessUnit:  u.BusinessUnit,
		ManagerId:     u.ManagerId,
		Grade:         u.Grade,
		Location:      u.Location,
		Country:       u.Country,
		Title:         u.Title,
		Type:          u.Type,
		PersonalEmail: u.PersonalEmail,
		CreatedTS:     time.Now(),
		UpdatedTS:     time.Now(),
	}

	result := utils.Db.Omit("EmployeeID").Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)

		return "", 0
	}

	return email, 1
}

func DeleteUserDAO(email string) int {
	result := utils.Db.Where("OFFICIAL_EMAIL = ?", email).Delete(gormModels.User{})

	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}
