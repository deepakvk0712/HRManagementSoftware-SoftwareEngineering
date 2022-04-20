package Controller

import (
	"encoding/json"
	"fmt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var starTime time.Time
var endTime time.Time
var isWorking bool

func StartWorking(w http.ResponseWriter, r *http.Request) {
	isWorking = true
	starTime.Format("2006-01-02 15:03:04")
	//User can only click starWorking once in 3hs
	if !starTime.IsZero() {
		now := time.Now().Sub(starTime).Hours()
		if now < 3 {
			errorResponses.SendBadRequestResponse(w, "You can only click once every 3 hours!")
			return
		}
	}

	starTime = time.Now()

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Start working!"
	res.Data = starTime.String()

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func StopWorking(w http.ResponseWriter, r *http.Request) {
	starTime.Format("2006-01-02 15:03:04")
	if isWorking == false {
		errorResponses.SendBadRequestResponse(w, "You haven't started working yet! Please start working first!")
		return
	} else {
		//User can only click stopWorking once in 3hs
		if !endTime.IsZero() {
			now := time.Now().Sub(endTime).Hours()
			if now < 3 {
				errorResponses.SendBadRequestResponse(w, "You can only click once every 3 hours!")
				return
			}
		}
		//calculate today's working hour
		isWorking = false
		endTime = time.Now()
		workingHours := endTime.Sub(starTime).Truncate(time.Hour).Hours()

		//Database Operations
		var workingH gormModels.WorkingHours
		//Get EmployeeID from token
		email := r.Context().Value("email").(string)
		workingH.EmployeeID = Dao.GetEmployeeIDByEmail(email)

		workingH.StartTime = starTime
		workingH.EndTime = endTime
		workingH.TimeWorked = workingHours

		if Dao.InsertWorkingRecord(workingH) == 0 {
			errorResponses.SendInternalServerErrorResponse(w)
			return
		}

		res := models.JsonResponse{}

		res.Error = ""
		res.Msg = "Leave work at:" + endTime.String()
		res.Data = ""

		jsonResponse, jsonError := json.Marshal(res)

		if jsonError != nil {
			fmt.Println(jsonError)

			errorResponses.SendInternalServerErrorResponse(w)
			return
		}

		utils.MessageHandler(w, jsonResponse, http.StatusCreated)

	}
}
func SetSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule gormModels.Schedule
	//Get schedule from front end
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "")
		return
	}
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	schedule.EmployeeID = Dao.GetEmployeeIDByEmail(email)

	//Database operation
	if Dao.SetSchedule(schedule) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Schedule added!"
	res.Data = schedule.ScheduleTitle + schedule.ScheduleContent

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)

}
func GetSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule gormModels.Schedule
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	schedule.EmployeeID = Dao.GetEmployeeIDByEmail(email)

	//Database operation
	s := make([]gormModels.Schedule, 1)
	var result int
	result, s = Dao.GetSchedule(schedule)
	if result == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	res := models.JsonResponseObject{}
	res.Error = ""
	res.Msg = "Success!"
	res.Data = s

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)

}
func DelSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule gormModels.Schedule
	//Get ID from front end
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}
	//Delete from Database
	if Dao.DelSchedule(schedule) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Schedule deleted!"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)

}
func GetWeekWorkingByID(w http.ResponseWriter, r *http.Request) {
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)

	//Get working records from database
	wkhr := make([]gormModels.WorkingHours, 1)
	var result int
	result, wkhr = Dao.GetWeekWorkingRecordByID(employee_id)
	if result == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	var totalWorkingHours float64
	workingStatus := make(map[string]int)
	for _, WH := range wkhr {
		totalWorkingHours += WH.TimeWorked
	}
	workingStatus["TotalHours"] = int(totalWorkingHours)

	var m1 float64 = 0
	var t2 float64 = 0
	var w3 float64 = 0
	var t4 float64 = 0
	var f5 float64 = 0
	var s6 float64 = 0
	var s7 float64 = 0
	//calculate each day's sum
	for _, ws := range wkhr {
		weekday := ws.StartTime.Weekday()
		switch weekday {
		case 1:
			m1 += ws.TimeWorked
		case 2:
			t2 += ws.TimeWorked
		case 3:
			w3 += ws.TimeWorked
		case 4:
			t4 += ws.TimeWorked
		case 5:
			f5 += ws.TimeWorked
		case 6:
			s6 += ws.TimeWorked
		case 7:
			s7 += ws.TimeWorked
		default:
			fmt.Println("Something Wrong！")
		}
	}

	workingStatus["Monday"] = int(m1)
	workingStatus["Tuesday"] = int(t2)
	workingStatus["Wednesday"] = int(w3)
	workingStatus["Thursday"] = int(t4)
	workingStatus["Friday"] = int(f5)
	workingStatus["Saturday"] = int(s6)
	workingStatus["Sunday"] = int(s7)

	res := models.JsonResponseObject{}

	res.Error = ""
	res.Msg = "Working detail of this week"
	res.Data = workingStatus

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func GetTodayWorkingHoursByID(w http.ResponseWriter, r *http.Request) {
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)

	//get today's total hours form database
	result, hourWorked, regular, overtime := Dao.GetTodayWorkingHoursByID(employee_id)
	if result == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	res := models.JsonResponseObject{}

	workingHours := make(map[string]float64)
	workingHours["TotalHour"] = hourWorked
	workingHours["RegularHours"] = regular
	workingHours["Overtime"] = overtime

	res.Error = ""
	res.Msg = "Today's working hours:"
	res.Data = workingHours

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func GetWorkingDetailsBetween(w http.ResponseWriter, r *http.Request) {
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)

	v := r.URL.Query()

	StartDate := v.Get("StartDate")
	EndDate := v.Get("EndDate")

	//Get Date start && end time from front end
	workingDetails := models.WorkingDetails{}

	workingDetails.StartDate, _ = time.Parse("2006-01-02", StartDate)
	workingDetails.EndDate, _ = time.Parse("2006-01-02", EndDate)

	if workingDetails.EndDate.Before(workingDetails.StartDate) {
		errorResponses.SendBadRequestResponse(w, "Star date must before End Date!")
		return
	}
	if workingDetails.StartDate.Equal(workingDetails.EndDate) {
		errorResponses.SendBadRequestResponse(w, "Star date and End Date cannot be the same day!")
		return
	}
	//Get Number of Present Day
	workingDetails.PresentDays = Dao.GetPreSentDays(employee_id, workingDetails)
	if workingDetails.PresentDays == -1 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	//Get Number of Absent Day
	workingDetails.AbsentDays, workingDetails.Holidays, workingDetails.Totaldays = Dao.GetAbsentAndHoliAndTotalDays(employee_id, workingDetails)
	if workingDetails.AbsentDays == -1 || workingDetails.Holidays == -1 || workingDetails.Totaldays == -1 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	//Get total hour worked, Total Regular Hour, Total Overtime Hour
	workingDetails.TotalWorkHour, workingDetails.TotalRegularHour, workingDetails.TotalOvertimeHour = Dao.Get3WorkingHours(employee_id, workingDetails)
	if workingDetails.TotalWorkHour == -1 || workingDetails.TotalRegularHour == -1 || workingDetails.TotalOvertimeHour == -1 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	//Calculate average work hour
	if workingDetails.PresentDays == 0 { // divide 0
		workingDetails.AverageWorkHour = 0
	} else {
		workingDetails.AverageWorkHour = workingDetails.TotalWorkHour / float64(workingDetails.PresentDays)
	}

	res := models.JsonResponseObject{}

	res.Error = ""
	res.Msg = "working details:"
	res.Data = workingDetails

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)

}
func SetHolidays(w http.ResponseWriter, r *http.Request) {

	//Get holiday information from front end
	holiday := gormModels.Holiday{}
	if err := json.NewDecoder(r.Body).Decode(&holiday); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "")
		return
	}
	//Database operation
	holiday.Date.Format("2006-01-02")
	fmt.Println(holiday.Date.String())
	if Dao.SetHoliday(holiday) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Holiday added!"
	res.Data = holiday.Date.String() + " " + holiday.Comment

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func SetWorkingHours(w http.ResponseWriter, r *http.Request) {
	//Get start&end date from front end
	WH := gormModels.WorkingHours{}
	FT := models.FrontendTime{}
	if err := json.NewDecoder(r.Body).Decode(&FT); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "")
		return
	}
	arr := strings.Split(FT.Date, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])

	arr2 := strings.Split(FT.StartTime, ":")
	h1, _ := strconv.Atoi(arr2[0])
	m1, _ := strconv.Atoi(arr2[1])

	arr3 := strings.Split(FT.EndTime, ":")
	h2, _ := strconv.Atoi(arr3[0])
	m2, _ := strconv.Atoi(arr3[1])

	if h1 > h2 {
		errorResponses.SendBadRequestResponse(w, "End time must greater than start time!")
		return
	}

	WH.StartTime = time.Date(year, time.Month(month), day, h1, m1, 0, 0, time.Local)
	WH.EndTime = time.Date(year, time.Month(month), day, h2, m2, 0, 0, time.Local)

	WH.TimeWorked = WH.EndTime.Sub(WH.StartTime).Minutes() / 60
	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)
	WH.EmployeeID = employee_id

	//Database operation
	if Dao.InsertWorkingRecord(WH) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Record added!"
	res.Data = "Employee " + strconv.Itoa(WH.EmployeeID) + " Start working at : " + WH.StartTime.String() + "Leave at : " + WH.EndTime.String()

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func SetTodayWorkingHours(w http.ResponseWriter, r *http.Request) {
	//Get start&end date from front end
	SE := models.StartAndEnd{}
	if err := json.NewDecoder(r.Body).Decode(&SE); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "")
		return
	}
	WH := gormModels.WorkingHours{}
	sh, _ := strconv.Atoi(SE.StartHour)
	eh, _ := strconv.Atoi(SE.EndHour)
	if sh < eh {
		errorResponses.SendBadRequestResponse(w, "end hour must greater than start hour!")
		return
	}
	WH.TimeWorked = float64(eh - sh)
	t := time.Now()
	WH.StartTime = time.Date(t.Year(), t.Month(), t.Day(), sh, 0, 0, 0, time.Local)
	WH.EndTime = time.Date(t.Year(), t.Month(), t.Day(), eh, 0, 0, 0, time.Local)

	//Get EmployeeID from token
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)
	WH.EmployeeID = employee_id

	//Database operation
	if Dao.InsertWorkingRecord(WH) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Record added!"
	res.Data = "Employee " + strconv.Itoa(WH.EmployeeID) + " Start working at: " + WH.StartTime.String() + " Leave at: " + WH.EndTime.String()

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
