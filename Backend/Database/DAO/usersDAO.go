package Dao

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
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
		Salary:        u.Salary,
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

	rows, _ := utils.Db.Raw("SELECT FIRST_NAME || \" \" || LAST_NAME, LOWER(OFFICIAL_EMAIL), EMPLOYEE_ID FROM USERS WHERE EMPLOYEE_ID = ? UNION SELECT FIRST_NAME || \" \" || LAST_NAME, LOWER(OFFICIAL_EMAIL), EMPLOYEE_ID FROM USERS WHERE MANAGER_ID = ? AND LOWER(OFFICIAL_EMAIL) <> ?", managerId, managerId, strings.ToLower(email)).Rows()
	defer rows.Close()

	//rows.Next()
	//rows.Scan(&teamDetails.Manager)

	teamMember := models.TeamMember{}

	for rows.Next() {
		rows.Scan(&teamMember.Name, &teamMember.EmailID, &teamMember.EmployeeID)
		if teamMember.EmployeeID == managerId {
			teamDetails.Manager = teamMember.Name
			continue
		}

		teamDetails.TeamMembers = append(teamDetails.TeamMembers, teamMember)
	}

	return teamDetails, 1

}

func GetEmployeeID(email string) (int, int) {
	employeeId := 0

	row := utils.Db.Raw("SELECT EMPLOYEE_ID FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return employeeId, 0
	}

	row.Scan(&employeeId)

	return employeeId, 1
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
func GetEmployeeIDByEmail(officialEmail string) int {

	var EID int
	result := utils.Db.Raw("SELECT EMPLOYEE_ID FROM users WHERE OFFICIAL_EMAIL = ?", officialEmail).Scan(&EID)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	fmt.Println(result.Error, result.RowsAffected)

	return EID
}

func GetFormStatus(email string) (int, bool, bool) {
	var isOnboard bool
	var isFinance bool

	row := utils.Db.Raw("SELECT IS_ONBOARD, IS_FINANCE FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return 0, false, false
	}

	row.Scan(&isOnboard, &isFinance)

	return 1, isOnboard, isFinance
}

func GetBusinessUnits() ([]string, int) {
	var businessUnits []string

	rows, _ := utils.Db.Raw("SELECT DISTINCT BUSINESS_UNIT FROM USERS").Rows()
	defer rows.Close()

	var businessUnit string

	for rows.Next() {
		rows.Scan(&businessUnit)

		businessUnits = append(businessUnits, businessUnit)
	}

	return businessUnits, 1

}

func GetTeamDetailsByBU(businessUnit, email string) ([]models.TeamMember, int) {
	var teamMembers []models.TeamMember

	rows, _ := utils.Db.Raw("SELECT FIRST_NAME || \" \" || LAST_NAME, LOWER(OFFICIAL_EMAIL), EMPLOYEE_ID FROM USERS WHERE LOWER(BUSINESS_UNIT) = ?", strings.ToLower(businessUnit)).Rows()
	defer rows.Close()

	teamMember := models.TeamMember{}

	for rows.Next() {

		rows.Scan(&teamMember.Name, &teamMember.EmailID, &teamMember.EmployeeID)
		if teamMember.EmailID == email {
			continue
		}

		teamMembers = append(teamMembers, teamMember)
	}

	return teamMembers, 1

}

func GetEmployeeNameByEmail(officialEmail string) string {

	var fn string
	var ln string
	result1 := utils.Db.Raw("SELECT FIRST_NAME FROM users WHERE OFFICIAL_EMAIL = ?", officialEmail).Scan(&fn)
	result2 := utils.Db.Raw("SELECT LAST_NAME FROM users WHERE OFFICIAL_EMAIL = ?", officialEmail).Scan(&ln)

	if result1.Error != nil && result2.Error != nil {
		fmt.Println(result1.Error, result2.Error)
		return "0"
	}
	fmt.Println(result1.RowsAffected, result2.RowsAffected)

	name1 := fn + " " + ln

	return name1
}
