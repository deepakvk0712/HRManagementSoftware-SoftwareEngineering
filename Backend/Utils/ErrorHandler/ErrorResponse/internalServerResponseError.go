package errorResponses

import (
	"net/http"

	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func SendInternalServerErrorResponse(w http.ResponseWriter) {
	internalServerError := errors.InternalServerError

	message := models.JsonResponse{}

	message.Error = internalServerError.Err
	message.Data = ""
	message.Msg = ""

	res := message.ToJSON()

	utils.MessageHandler(w, res, internalServerError.Code)
}
