package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strings"
	"time"
)

func UpdateLeave(leaveInfo models.ApplyLeave, email string) int {
	startDate, _ := time.Parse("2006-01-02", leaveInfo.StartDate)
	endDate, _ := time.Parse("2006-01-02", leaveInfo.EndDate)

	numOfDays := int(endDate.Sub(startDate).Hours()/24) + 1

	if leaveInfo.LeaveType == utils.PaidLeave {
		utils.Db.Exec("UPDATE USERS SET LEAVES_REMAINING = LEAVES_REMAINING - ?, PAID_LEAVES = PAID_LEAVES - ? WHERE LOWER(OFFICIAL_EMAIL) = ?", numOfDays, numOfDays, strings.ToLower(email))
	} else {
		utils.Db.Exec("UPDATE USERS SET LEAVES_REMAINING = LEAVES_REMAINING - ?, UNPAID_LEAVES = UNPAID_LEAVES - ? WHERE LOWER(OFFICIAL_EMAIL) = ?", numOfDays, numOfDays, strings.ToLower(email))
	}

	leave := gormModels.Leave{
		StartDate: startDate,
		EndDate:   endDate,
		Subject:   leaveInfo.Subject,
		Reason:    leaveInfo.Reason,
		Email:     email,
		LeaveType: leaveInfo.LeaveType,
	}

	result := utils.Db.Omit("LeaveID").Create(&leave)
	if result.Error != nil {
		fmt.Println(result.Error)

		return 0
	}

	return 1
}

func GetLeaveStatus(email string) (int, int, int) {
	paidLeaves := 0
	unpaidLeaves := 0

	row := utils.Db.Raw("SELECT PAID_LEAVES, UNPAID_LEAVES FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return 0, 0, 0
	}

	row.Scan(&paidLeaves, &unpaidLeaves)
	return paidLeaves, unpaidLeaves, 1
}
