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

func GetProfileDetails(email string) (models.ProfileDetails, int) {
	profileDetails := models.ProfileDetails{}

	row := utils.Db.Raw("SELECT FIRST_NAME, LAST_NAME, TITLE, ABOUT_ME FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return profileDetails, 0
	}

	row.Scan(&profileDetails.FirstName, &profileDetails.LastName, &profileDetails.Title, &profileDetails.AboutMe)
	return profileDetails, 1
}

func UpdateProfileDetails(userProfile models.ProfileDetails, email string) int {
	utils.Db.Exec("UPDATE USERS SET FIRST_NAME = ?, LAST_NAME = ?, TITLE = ?, ABOUT_ME = ?WHERE LOWER(OFFICIAL_EMAIL) = ?", userProfile.FirstName, userProfile.LastName, userProfile.Title, userProfile.AboutMe, strings.ToLower(email))

	return 1
}
