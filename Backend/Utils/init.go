package utils

import (
	"github.com/joho/godotenv"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	"os"
	"strconv"
)

var LoginJWT models.JWTParams

func Init() {
	godotenv.Load("../secrets.env")
	ConnectToDb()

	loginTimeout, _ := strconv.Atoi(os.Getenv("LOGIN_TIMEOUT"))

	LoginJWT = models.JWTParams{
		SecretKey:    os.Getenv("JWT_SECRET_KEY"),
		Issuer:       "Login",
		LoginTimeout: loginTimeout,
	}
}
