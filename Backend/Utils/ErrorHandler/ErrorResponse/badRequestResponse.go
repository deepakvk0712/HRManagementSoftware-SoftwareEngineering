package errorResponses

import (
	"net/http"

	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
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
