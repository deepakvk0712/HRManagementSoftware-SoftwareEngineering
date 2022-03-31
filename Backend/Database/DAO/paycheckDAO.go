package Dao

import (
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func GetPaycheck(employeeId int, param models.PaycheckQuery) []gormModels.Paycheck {
	var paychecks []gormModels.Paycheck

	utils.Db.Where("EMPLOYEE_ID = ? AND CHECK_DATE >= ? AND CHECK_DATE <= ?", employeeId, param.PayBeginDate, param.PayEndDate).Find(&paychecks)

	return paychecks
}

func GetSalaries(managerId int, param models.PaycheckQuery) []models.TeamSalary {
	var teamSalaries []models.TeamSalary
	employeeSalary := models.TeamSalary{}

	rows, _ := utils.Db.Raw("SELECT U.FIRST_NAME || \" \" || U.LAST_NAME, U.EMPLOYEE_ID, P.CHECK_DATE, P.PAY_BEGIN_DATE, P.PAY_END_DATE, P.AMOUNT_PAID FROM USERS U JOIN (SELECT CHECK_DATE, PAY_BEGIN_DATE, PAY_END_DATE, AMOUNT_PAID FROM PAYCHECKS WHERE CHECK_DATE >= ? AND CHECK_DATE <= ?) P ON U.EMPLOYEE_ID = P.EMPLOYEE_ID WHERE U.MANAGER_ID = ?", param.PayBeginDate, param.PayEndDate, managerId).Rows()
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&employeeSalary.EmployeeName, &employeeSalary.EmployeeID, &employeeSalary.CheckDate, &employeeSalary.PayBeginDate, &employeeSalary.PayEndDate, &employeeSalary.AmountPaid)
		teamSalaries = append(teamSalaries, employeeSalary)
	}

	return teamSalaries
}
