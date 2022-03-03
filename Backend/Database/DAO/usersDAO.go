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
func UpdateEmployeeDao(user gormModels.User) int {
	updateTime := time.Now()

	result := utils.Db.Model(&user).Select(
		"DriversLicense",
		"SSN",
		"StateID",
		"Address",
		"Phone",
		"AlternateEmails",
		"FirstName",
		"LastName",
		"Gender",
		"DateOfBirth",
		"Race",
		"Ethnicity",
		"Citizenship",
		"Nationality",
		"Pronouns",
		"UpdatedTS",
	).Where("Employee_ID = ?", user.EmployeeID).Updates(gormModels.User{
		DriversLicense:  user.DriversLicense,
		SSN:             user.SSN,
		StateID:         user.StateID,
		Address:         user.Address,
		Phone:           user.Phone,
		AlternateEmails: user.AlternateEmails,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Gender:          user.Gender,
		DateOfBirth:     user.DateOfBirth,
		Race:            user.Race,
		Ethnicity:       user.Ethnicity,
		Citizenship:     user.Citizenship,
		Nationality:     user.Nationality,
		Pronouns:        user.Pronouns,
		UpdatedTS:       updateTime})
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1

}

func UpdateEmployeeBankingDao(user gormModels.User) int {

	updateTime := time.Now()

	result := utils.Db.Model(&user).Select(
		"RoutingNumber",
		"AccountNumber",
		"Bank",
		"UpdatedTS").Where("Employee_ID = ?", user.EmployeeID).Updates(gormModels.User{
		RoutingNumber: user.RoutingNumber,
		AccountNumber: user.AccountNumber,
		Bank:          user.Bank,
		UpdatedTS:     updateTime})

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
