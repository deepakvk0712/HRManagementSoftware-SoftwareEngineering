package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
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
func GetTodayWorkingHoursByID(employeeID int) (int, int) {
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
func GetPreSentDays(EID int, WD models.WorkingDetails) int {

	workingHours := make([]gormModels.WorkingHours, 1)
	result := utils.Db.Where("employee_id = ? AND start_time >=? AND end_time <= ?", EID, WD.StartDate, WD.EndDate).Find(&workingHours)

	if result.Error != nil {
		fmt.Println(result.Error)
		return -1
	}
	fmt.Println(result.Error, result.RowsAffected)

	for i := 0; i < len(workingHours)-1; i++ {
		if workingHours[i].StartTime.Day() == workingHours[i+1].StartTime.Day() {
			workingHours = append(workingHours[:i], workingHours[i+1:]...)
			i--
		}
	}
	return len(workingHours)
}
func GetAbsentDays(EID int, WD models.WorkingDetails) int {
	var totalDays int
	var presentDays int
	var holidays int
	holidays = 0
	totalDays = WD.EndDate.YearDay() - WD.StartDate.YearDay()
	for i := 0; i < totalDays; i++ {
		temp := WD.StartDate.AddDate(0, 0, i)
		if IsHoliday(temp) {
			holidays++
		}
	}

	presentDays = GetPreSentDays(EID, WD)
	if presentDays == -1 {
		return -1
	}
	absentDays := totalDays - presentDays - holidays
	return absentDays
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
	fmt.Println(result.Error, result.RowsAffected)
	for _, data := range workingHours {
		totalHourWorked += data.TimeWorked
		hoursWorked := data.EndTime.Sub(data.StartTime).Minutes() / 60.0
		//Total hours of a day must bigger than 8 in order to count for overtime
		// If employee works on holiday, whole working hour counted as overtime
		if IsHoliday(data.StartTime) {
			totalOvertimeHour += hoursWorked
			continue
		} else { // Regular working days
			if hoursWorked < 8 {
				totalRegularHour += hoursWorked
			} else {
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
