package Controller

import (
	"encoding/json"
	"fmt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func GetTrainingStatus(w http.ResponseWriter, req *http.Request) {
	//role := req.Context().Value("role").(byte)
	email := req.Context().Value("email").(string)

	trainingCompleted := false

	score, err := Dao.GetTrainingScore(email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	if score >= 7 && score <= 10 {
		trainingCompleted = true
	} else {
		score = -1
	}

	Msg := struct {
		TrainingCompleted bool `json:"trainingCompleted"`
		Score             int  `json:"score"`
	}{
		TrainingCompleted: trainingCompleted,
		Score:             score,
	}

	data, jsonError := json.Marshal(Msg)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = ""
	res.Data = string(data)

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}

func PostTrainingQuestions(w http.ResponseWriter, req *http.Request) {
	email := req.Context().Value("email").(string)

	q := models.TrainingQuestions{}

	if err := json.NewDecoder(req.Body).Decode(&q); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	score, err := Dao.CalculateScore(q, email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	Msg := struct {
		Score int `json:"score"`
	}{
		Score: score,
	}

	data, jsonError := json.Marshal(Msg)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = ""
	res.Data = string(data)

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
