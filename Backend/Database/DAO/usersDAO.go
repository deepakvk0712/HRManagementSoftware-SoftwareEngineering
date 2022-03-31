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

func GetTeamDetails(email string) (models.TeamDetails, int) {
	teamDetails := models.TeamDetails{}
	managerId := 0

	row := utils.Db.Raw("SELECT MANAGER_ID, BUSINESS_UNIT FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return teamDetails, 0
	}

	row.Scan(&managerId, &teamDetails.BusinessUnit)

	rows, _ := utils.Db.Raw("SELECT FIRST_NAME || \" \" || LAST_NAME FROM USERS WHERE EMPLOYEE_ID = ? UNION SELECT FIRST_NAME || \" \" || LAST_NAME FROM USERS WHERE MANAGER_ID = ? AND LOWER(OFFICIAL_EMAIL) <> ?", managerId, managerId, strings.ToLower(email)).Rows()
	defer rows.Close()

	rows.Next()
	rows.Scan(&teamDetails.Manager)

	temp := ""
	for rows.Next() {
		rows.Scan(&temp)
		teamDetails.TeamMembers = append(teamDetails.TeamMembers, temp)
	}

	return teamDetails, 1

}

func UpdateProfileDetails(userProfile models.ProfileDetails, email string) int {
	utils.Db.Exec("UPDATE USERS SET FIRST_NAME = ?, LAST_NAME = ?, TITLE = ?, ABOUT_ME = ?, PHONE = ?, PERSONAL_EMAIL = ?, ADDRESS = ? WHERE LOWER(OFFICIAL_EMAIL) = ?", userProfile.FirstName, userProfile.LastName, userProfile.Title, userProfile.AboutMe, userProfile.Phone, userProfile.PersonalEmail, userProfile.Address, strings.ToLower(email))

	return 1
}

func UpdateEmployeeDao(user gormModels.User) int {
	updateTime := time.Now()
	user.UpdatedTS = updateTime
	user.IsOnboard = true

	result := utils.Db.Model(&user).Where("Employee_ID = ?", user.EmployeeID).Updates(&user)
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
		"UpdatedTS",
		"IsFinance").Where("Employee_ID = ?", user.EmployeeID).Updates(gormModels.User{
		RoutingNumber: user.RoutingNumber,
		AccountNumber: user.AccountNumber,
		Bank:          user.Bank,
		UpdatedTS:     updateTime,
		IsFinance:     true})

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
