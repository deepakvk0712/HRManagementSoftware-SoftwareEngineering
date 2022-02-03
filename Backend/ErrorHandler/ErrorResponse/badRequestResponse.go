package errorResponses

import (
	"net/http"

	models "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func SendBadRequestResponse(w http.ResponseWriter) {

	badRequestError := errors.BadServerError

	message := models.JsonResponse{}

	message.Error = badRequestError.Err
	message.Data = ""
	message.Msg = ""

	res := message.ToJSON()

	utils.MessageHandler(w, res, badRequestError.Code)
}
