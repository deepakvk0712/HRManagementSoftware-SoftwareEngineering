package errorResponses

import (
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
)

func SendUnauthorizedResponse(w http.ResponseWriter, msg string) {
	unauthorizedError := errors.UnauthorizedError

	message := models.JsonResponse{}

	message.Error = unauthorizedError.Err
	message.Data = ""
	message.Msg = msg

	res := message.ToJSON()

	utils.MessageHandler(w, res, unauthorizedError.Code)
}
