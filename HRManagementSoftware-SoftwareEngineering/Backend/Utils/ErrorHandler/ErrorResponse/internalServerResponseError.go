package errorResponses

import (
	"net/http"

	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func SendInternalServerErrorResponse(w http.ResponseWriter) {
	internalServerError := errors.InternalServerError

	//define reply message
	message := models.JsonResponse{}
	message.Error = internalServerError.Err
	message.Data = ""
	message.Msg = ""

	//convert to json
	res := message.ToJSON()
	//send InternalServerError back to front end
	utils.MessageHandler(w, res, internalServerError.Code)
}
