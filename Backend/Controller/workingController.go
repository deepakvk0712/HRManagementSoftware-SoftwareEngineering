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
	//User can only click starWorking once in 12hs
	if !starTime.IsZero() {
		now := time.Now().Sub(starTime).Hours()
		if now < 12 {
			errorResponses.SendBadRequestResponse(w, "You can only click once every 12 hours!")
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
		//User can only click stopWorking once in 12hs
		if !endTime.IsZero() {
			now := time.Now().Sub(endTime).Hours()
			if now < 12 {
				errorResponses.SendBadRequestResponse(w, "You can only click once every 12 hours!")
				return
			}
		}
		//calculate today's working hour
		isWorking = false
		endTime = time.Now()
		workingHours := endTime.Sub(starTime).Truncate(time.Hour).Hours()

		//Database Operations
		var workingH gormModels.WorkingHours
		workingH.StartTime = starTime
		workingH.EndTime = endTime
		workingH.TimeWorked = workingHours

		if err := json.NewDecoder(r.Body).Decode(&workingH); err != nil {
			fmt.Println(err)
			errorResponses.SendBadRequestResponse(w, "bad request!")
			return
		}

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
