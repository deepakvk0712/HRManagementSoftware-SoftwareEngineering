package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		userLogin := models.UserLogin{}

		if err := json.NewDecoder(req.Body).Decode(&userLogin); err != nil {
			fmt.Println(err)

			errorResponses.SendBadRequestResponse(w, "")

			return
		}

		ctx := context.WithValue(req.Context(), "user", userLogin)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
