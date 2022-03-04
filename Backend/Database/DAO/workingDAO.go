package Dao

import (
	"fmt"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"time"
)

func InsertWorkingRecord(workinghours gormModels.WorkingHours) int {
	result := utils.Db.Create(&workinghours)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
func SetSchedule(schedule gormModels.Schedule) int {
	result := utils.Db.Create(&schedule)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
func GetSchedule(schedule gormModels.Schedule) (int, []gormModels.Schedule) {
	var s []gormModels.Schedule
	result := utils.Db.Model(schedule).Where("employee_id = ?", schedule.EmployeeID).Find(&s)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, s
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1, s
}
func DelSchedule(schedule gormModels.Schedule) int {

	result := utils.Db.Delete(schedule)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
func GetWeekWorkingRecordByID(employeeID int64) (int, []gormModels.WorkingHours) {

	workingHours := make([]gormModels.WorkingHours, 1)
	now := time.Now()
	weekday := now.Weekday()
	var startTime time.Time
	var monday time.Time
	switch weekday {
	case 1:
		monday = now
		startTime = now.AddDate(0, 0, -1)
	case 2:
		monday = now.AddDate(0, 0, -1)
		startTime = now.AddDate(0, 0, -2)
	case 3:
		monday = now.AddDate(0, 0, -2)
		startTime = now.AddDate(0, 0, -3)
	case 4:
		monday = now.AddDate(0, 0, -3)
		startTime = now.AddDate(0, 0, -4)
	case 5:
		monday = now.AddDate(0, 0, -4)
		startTime = now.AddDate(0, 0, -5)
	case 6:
		monday = now.AddDate(0, 0, -5)
		startTime = now.AddDate(0, 0, -6)
	case 7:
		monday = now.AddDate(0, 0, -6)
		startTime = now.AddDate(0, 0, -7)
	default:
		fmt.Println("Something wrong!")
	}

	result := utils.Db.Where("employee_id = ? AND start_time >? AND end_time < ?", employeeID, startTime, now).Order("ID asc").Find(&workingHours)
	// SELECT * FROM working_hours WHERE employee_id = employeeID AND start_time > startTime AND end_time < now order by ID asc;
	fmt.Println(workingHours)
	length := len(workingHours)
	for i := 0; i < length; i++ {
		if workingHours[i].StartTime.YearDay() < monday.YearDay() {
			workingHours = append(workingHours[:i], workingHours[i+1:]...)
			i--
			length--
		}
	}

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, workingHours
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1, workingHours
}
func GetTodayWorkingHoursByID(employeeID int64) (int, int) {
	now := time.Now()
	y1, m1, d1 := now.Year(), now.Month(), now.Day()
	workingHours := make([]gormModels.WorkingHours, 1)
	result := utils.Db.Where("employee_id = ? AND start_time >? AND end_time < ?", employeeID, now.AddDate(0, 0, -1), now).Find(&workingHours)

	var hourWorked int = 0
	for i, _ := range workingHours {
		y2, m2, d2 := workingHours[i].EndTime.Year(), workingHours[i].EndTime.Month(), workingHours[i].EndTime.Day()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			hourWorked += int(workingHours[i].TimeWorked)
		}
	}

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, hourWorked
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1, hourWorked
}
