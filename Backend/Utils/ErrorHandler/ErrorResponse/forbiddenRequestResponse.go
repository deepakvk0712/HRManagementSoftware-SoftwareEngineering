package errorResponses

import (
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
)

func SendForbiddenRequestResponse(w http.ResponseWriter) {

	forbiddenRequestError := errors.ForbiddenRequestError

	message := models.JsonResponse{}

	message.Error = forbiddenRequestError.Err
	message.Data = ""
	message.Msg = ""

	res := message.ToJSON()

	utils.MessageHandler(w, res, forbiddenRequestError.Code)
}
