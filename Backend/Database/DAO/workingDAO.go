package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strconv"
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
func GetWeekWorkingRecordByID(employeeID int) (int, []gormModels.WorkingHours) {

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
func GetTodayWorkingHoursByID(employeeID int) (int, float64, float64, float64) {
	now := time.Now()
	y1, m1, d1 := now.Year(), now.Month(), now.Day()
	startTime := time.Date(y1, m1, d1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(y1, m1, d1, 23, 59, 59, 0, time.Local)
	workingHours := make([]gormModels.WorkingHours, 1)
	result := utils.Db.Where("employee_id = ? AND start_time >? AND end_time < ?", employeeID, startTime, endTime).Find(&workingHours)

	var hourWorked float64 = 0
	for _, data := range workingHours {
		hourWorked += data.EndTime.Sub(data.StartTime).Minutes() / 60
	}
	var regularHour float64
	var overtime float64
	if hourWorked > 8 {
		regularHour = 8
		overtime = hourWorked - 8
	} else {
		regularHour = hourWorked
		overtime = 0
	}
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, -1, -1, -1
	}

	fmt.Println(result.Error, result.RowsAffected)

	return 1, hourWorked, regularHour, overtime
}
func GetPreSentDays(EID int, WD models.WorkingDetails) int {

	workingHours := make([]gormModels.WorkingHours, 1)
	result := utils.Db.Where("employee_id = ? AND start_time >=? AND end_time <= ?", EID, WD.StartDate, WD.EndDate).Find(&workingHours)

	if result.Error != nil {
		fmt.Println(result.Error)
		return -1
	}
	fmt.Println(result.Error, result.RowsAffected)

	for i := 0; i < len(workingHours)-1; i++ {
		if workingHours[i].StartTime.YearDay() == workingHours[i+1].StartTime.YearDay() {
			workingHours = append(workingHours[:i], workingHours[i+1:]...)
			i--
		}
	}
	return len(workingHours)
}
func GetAbsentAndHoliAndTotalDays(EID int, WD models.WorkingDetails) (int, int, int) {
	fmt.Println(WD)
	var totalDays int
	var presentDays int
	var holidays int
	holidays = 0
	totalDays = WD.EndDate.YearDay() - WD.StartDate.YearDay() + 1

	for i := 0; i < totalDays; i++ {
		temp := WD.StartDate.AddDate(0, 0, i)
		if IsHoliday(temp) {
			holidays++
		}
	}
	fmt.Println("holidays : " + strconv.Itoa(holidays))
	presentDays = GetPreSentDays(EID, WD)
	if presentDays == -1 {
		return -1, -1, -1
	}
	absentDays := totalDays - presentDays - holidays
	return absentDays, holidays, totalDays
}
func Get3WorkingHours(EID int, WD models.WorkingDetails) (float64, float64, float64) {
	var totalHourWorked float64
	var totalRegularHour float64
	var totalOvertimeHour float64
	totalHourWorked = 0
	totalRegularHour = 0
	totalOvertimeHour = 0

	workingHours := make([]gormModels.WorkingHours, 1)
	result := utils.Db.Where("employee_id = ? AND start_time >=? AND end_time <= ?", EID, WD.StartDate, WD.EndDate).Order("ID asc").Find(&workingHours)
	// SELECT * FROM working_hours WHERE employee_id = employeeID AND start_time > startTime AND end_time < now order by ID asc;
	if result.Error != nil {
		fmt.Println(result.Error)
		return -1, -1, -1
	}
	fmt.Print(workingHours)
	fmt.Println(result.Error, result.RowsAffected)
	for _, data := range workingHours {
		hoursWorked := data.EndTime.Sub(data.StartTime).Minutes() / 60.0
		totalHourWorked += hoursWorked
		//Total hours of a day must bigger than 8 in order to count for overtime
		// If employee works on holiday, whole working hour counted as overtime
		if IsHoliday(data.StartTime) {
			totalOvertimeHour += hoursWorked
		} else { // Regular working days
			if hoursWorked <= 8 {
				totalRegularHour += hoursWorked
			} else {
				totalRegularHour += 8
				totalOvertimeHour += hoursWorked - 8
			}
		}
	}

	return totalHourWorked, totalRegularHour, totalOvertimeHour
}

func IsHoliday(day time.Time) bool {
	var result bool
	result = false
	//Is weekend
	if day.Weekday() == 6 || day.Weekday() == 7 {
		result = true
	}
	//Is holiday
	holiday := make([]gormModels.Holiday, 1)

	result1 := utils.Db.Find(&holiday)
	if result1.Error != nil {
		fmt.Println(result1.Error)
		return false
	}
	fmt.Println(result1.Error, result1.RowsAffected)

	for _, data := range holiday {
		if day.YearDay() == data.Date.YearDay() {
			result = true
		}
	}

	return result
}
func SetHoliday(holiday gormModels.Holiday) int {

	result := utils.Db.Create(&holiday)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	fmt.Println(result.Error, result.RowsAffected)
	return 1
}
