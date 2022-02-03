package errorResponses

import (
	"net/http"

	models "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
	utils "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
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
