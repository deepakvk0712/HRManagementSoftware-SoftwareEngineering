package utils

func Init() {
	// errors.InternalServerErrorJSON = errors.InternalServerError.ToJSON()
	SetEnvironmentVariables()
	ConnectToDb()

}
