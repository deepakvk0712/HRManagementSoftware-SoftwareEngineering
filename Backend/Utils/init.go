package utils

import (
	errors "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
)

func Init() {
	errors.InternalServerErrorJSON = errors.InternalServerError.ToJSON()
}
