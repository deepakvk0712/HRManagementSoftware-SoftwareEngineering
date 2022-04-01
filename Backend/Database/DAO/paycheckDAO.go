package Dao

import (
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func GetPaycheck(employeeId int, start, end string) []gormModels.Paycheck {
	var paychecks []gormModels.Paycheck

	if start != "" && end != "" {
		utils.Db.Where("EMPLOYEE_ID = ? AND CHECK_DATE >= ? AND CHECK_DATE <= ?", employeeId, start, end).Find(&paychecks)

		return paychecks
	}

	utils.Db.Where("EMPLOYEE_ID = ?", employeeId).Find(&paychecks)

	return paychecks
}

func GetSalaries(managerId int, start, end string) []models.TeamSalary {
	var teamSalaries []models.TeamSalary
	employeeSalary := models.TeamSalary{}

	if start != "" && end != "" {
		rows, _ := utils.Db.Raw("SELECT U.FIRST_NAME || \" \" || U.LAST_NAME, U.EMPLOYEE_ID, P.CHECK_DATE, P.PAY_BEGIN_DATE, P.PAY_END_DATE, P.AMOUNT_PAID FROM USERS U JOIN (SELECT EMPLOYEE_ID, CHECK_DATE, PAY_BEGIN_DATE, PAY_END_DATE, AMOUNT_PAID FROM PAYCHECKS WHERE CHECK_DATE >= ? AND CHECK_DATE <= ?) P ON U.EMPLOYEE_ID = P.EMPLOYEE_ID WHERE U.MANAGER_ID = ?", start, end, managerId).Rows()
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&employeeSalary.EmployeeName, &employeeSalary.EmployeeID, &employeeSalary.CheckDate, &employeeSalary.PayBeginDate, &employeeSalary.PayEndDate, &employeeSalary.AmountPaid)
			teamSalaries = append(teamSalaries, employeeSalary)
		}
	} else {
		rows, _ := utils.Db.Raw("SELECT U.FIRST_NAME || \" \" || U.LAST_NAME, U.EMPLOYEE_ID, P.CHECK_DATE, P.PAY_BEGIN_DATE, P.PAY_END_DATE, P.AMOUNT_PAID FROM USERS U JOIN (SELECT EMPLOYEE_ID, CHECK_DATE, PAY_BEGIN_DATE, PAY_END_DATE, AMOUNT_PAID FROM PAYCHECKS WHERE (EMPLOYEE_ID, CHECK_DATE) in (SELECT EMPLOYEE_ID, MAX(CHECK_DATE) FROM PAYCHECKS GROUP BY EMPLOYEE_ID)) P ON U.EMPLOYEE_ID = P.EMPLOYEE_ID WHERE U.MANAGER_ID = ?", managerId).Rows()
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&employeeSalary.EmployeeName, &employeeSalary.EmployeeID, &employeeSalary.CheckDate, &employeeSalary.PayBeginDate, &employeeSalary.PayEndDate, &employeeSalary.AmountPaid)
			teamSalaries = append(teamSalaries, employeeSalary)
		}
	}

	return teamSalaries
}

func UpdateSalary(employeeId int, newSalary float32) int {
	utils.Db.Exec("UPDATE USERS SET SALARY = ? WHERE EMPLOYEE_ID = ?", newSalary, employeeId)

	return 1
}
