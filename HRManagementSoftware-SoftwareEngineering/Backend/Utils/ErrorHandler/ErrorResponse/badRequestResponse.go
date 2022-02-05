package errorResponses

import (
	"net/http"

	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
)

func SendBadRequestResponse(w http.ResponseWriter) {

	badRequestError := errors.BadServerError
	//define reply message
	message := models.JsonResponse{}
	message.Error = badRequestError.Err
	message.Data = ""
	message.Msg = ""

	//convert to json
	res := message.ToJSON()

	//send response to front end
	utils.MessageHandler(w, res, badRequestError.Code)
}
