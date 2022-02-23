package middleware

import (
	"context"
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

		dbUser, err := Dao.GetUserDetailsDAO(userLogin.Email)
		if err == 0 {
			errorResponses.SendInternalServerErrorResponse(w)

			return
		}

		//dbEmail, _ := Dao.GetEmailDAO(userLogin.Email)
		if dbUser.Email != userLogin.Email {
			errorResponses.SendBadRequestResponse(w, "Invalid credentials")

			return
		}

		//dbPassword, _ := Dao.GetPasswordDAO(userLogin.Email)
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userLogin.Password)); err != nil {
			fmt.Println("Passwords do not match")

			errorResponses.SendBadRequestResponse(w, "Invalid credentials")

			return
		}

		ctx := context.WithValue(req.Context(), "role", dbUser.Role)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
