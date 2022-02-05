package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"time"
)

func CreateUserDAO(u models.User, email string) int {
	//var user = struct {
	//	models.User
	//	string
	//}{
	//	u,
	//	email,
	//}

	//_ = utils.Db.AutoMigrate(&gormModels.User{})

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

	//result := utils.Db.Select("FirstName", "LastName", "BusinessUnit", "ManagerId", "Grade", "Location", "Country", "Title", "Type", "PersonalEmail", "OfficialEmail").Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}
