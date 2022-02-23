package middleware

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		userLogin := req.Context().Value("user").(models.UserLogin)

		dbEmail, _ := Dao.GetEmailDAO(userLogin.Email)
		if dbEmail != userLogin.Email {
			errorResponses.SendBadRequestResponse(w, "Invalid credentials")

			return
		}

		dbPassword, _ := Dao.GetPasswordDAO(userLogin.Email)
		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userLogin.Password)); err != nil {
			fmt.Println("Passwords do not match")

			errorResponses.SendBadRequestResponse(w, "Invalid credentials")

			return
		}

		next.ServeHTTP(w, req)
	})
}
