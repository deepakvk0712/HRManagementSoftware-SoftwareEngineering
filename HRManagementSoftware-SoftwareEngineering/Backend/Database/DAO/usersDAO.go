package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"time"
)

func CreateUserDAO(u models.User, email string) int {
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

		return 0
	}

	return 1
}
func UpdateEmployeeDao(user gormModels.User) int {

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
		"RoutingNumber",
		"AccountNumber",
		"Bank").Updates(gormModels.User{
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
		RoutingNumber:   user.RoutingNumber,
		AccountNumber:   user.AccountNumber,
		Bank:            user.Bank})

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1

}
