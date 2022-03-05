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

		if err := json.NewDecoder(r.Body).Decode(&workingH); err != nil {
			fmt.Println(err)
			errorResponses.SendBadRequestResponse(w, "bad request!")
			return
		}

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
	//Get EmployeeID from front end
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}

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
	//Get EmployeeID from front end
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}
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
	//Get EmployeeID from front end
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
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
	var user gormModels.User
	//Get EmployeeID from front end
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}
	employee_id := user.EmployeeID
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
	var user gormModels.User
	//Get EmployeeID from front end
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}
	employee_id := user.EmployeeID
	//get today's total hours form database
	result, hourWorked := Dao.GetTodayWorkingHoursByID(employee_id)
	if result == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	res := models.JsonResponseObject{}

	res.Error = ""
	res.Msg = "Today's Total hour of working:"
	res.Data = hourWorked

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
